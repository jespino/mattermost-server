// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package file_content_extract

import (
	"context"
	"net/http"
	"time"

	"github.com/mattermost/mattermost-server/v5/app"
	"github.com/mattermost/mattermost-server/v5/jobs"
	"github.com/mattermost/mattermost-server/v5/mlog"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/services/docextractor"
)

const (
	JobName             = "FileContentExtract"
	BatchSize           = 1000
	EstimatedFilesCount = 100000
	TimeBetweenBatches  = 100
)

type Worker struct {
	name      string
	stop      chan bool
	stopped   chan bool
	jobs      chan model.Job
	jobServer *jobs.JobServer
	app       *app.App
}

type ExtractContentProgress struct {
	Now          time.Time
	StartAtTime  int64
	EndAtTime    int64
	LastFileTime int64
	TotalCount   int64
	DoneCount    int64
	Done         bool
}

func (ecp *ExtractContentProgress) CurrentProgress() int64 {
	return (ecp.DoneCount * 100) / ecp.TotalCount
}

func (m *FileContentExtractJobInterfaceImpl) MakeWorker() model.Worker {
	worker := Worker{
		name:      JobName,
		stop:      make(chan bool, 1),
		stopped:   make(chan bool, 1),
		jobs:      make(chan model.Job),
		jobServer: m.App.Srv().Jobs,
		app:       m.App,
	}
	return &worker
}

func (worker *Worker) Run() {
	mlog.Debug("Worker started", mlog.String("worker", worker.name))

	defer func() {
		mlog.Debug("Worker finished", mlog.String("worker", worker.name))
		worker.stopped <- true
	}()

	for {
		select {
		case <-worker.stop:
			mlog.Debug("Worker received stop signal", mlog.String("worker", worker.name))
			return
		case job := <-worker.jobs:
			mlog.Debug("Worker received a new candidate job.", mlog.String("worker", worker.name))
			worker.DoJob(&job)
		}
	}
}

func (worker *Worker) Stop() {
	mlog.Debug("Worker stopping", mlog.String("worker", worker.name))
	worker.stop <- true
	<-worker.stopped
}

func (worker *Worker) JobChannel() chan<- model.Job {
	return worker.jobs
}

func (worker *Worker) DoJob(job *model.Job) {
	if claimed, err := worker.jobServer.ClaimJob(job); err != nil {
		mlog.Warn("Worker experienced an error while trying to claim job",
			mlog.String("worker", worker.name),
			mlog.String("job_id", job.Id),
			mlog.String("error", err.Error()))
		return
	} else if !claimed {
		return
	}

	// TODO: Run the job code here
	progress := ExtractContentProgress{
		Done:       false,
		TotalCount: 0,
		DoneCount:  0,
	}
	worker.jobServer.SetJobProgress(job, progress.CurrentProgress())

	if count, err := worker.jobServer.Store.FileInfo().CountAll(); err != nil {
		mlog.Warn("Worker: Failed to fetch total file info count for job. An estimated value will be used for progress reporting.", mlog.String("workername", worker.name), mlog.String("job_id", job.Id), mlog.Err(err))
		progress.TotalCount = EstimatedFilesCount
	} else {
		progress.TotalCount = count
	}

	cancelCtx, cancelCancelWatcher := context.WithCancel(context.Background())
	cancelWatcherChan := make(chan interface{}, 1)
	go worker.jobServer.CancellationWatcher(cancelCtx, job.Id, cancelWatcherChan)

	defer cancelCancelWatcher()

	for {
		select {
		case <-cancelWatcherChan:
			mlog.Info("Worker: Indexing job has been canceled via CancellationWatcher", mlog.String("workername", worker.name), mlog.String("job_id", job.Id))
			if err := worker.jobServer.SetJobCanceled(job); err != nil {
				mlog.Error("Worker: Failed to mark job as cancelled", mlog.String("workername", worker.name), mlog.String("job_id", job.Id), mlog.Err(err))
			}
			return

		case <-worker.stop:
			mlog.Info("Worker: Indexing has been canceled via Worker Stop", mlog.String("workername", worker.name), mlog.String("job_id", job.Id))
			if err := worker.jobServer.SetJobCanceled(job); err != nil {
				mlog.Error("Worker: Failed to mark job as canceled", mlog.String("workername", worker.name), mlog.String("job_id", job.Id), mlog.Err(err))
			}
			return

		case <-time.After(TimeBetweenBatches * time.Millisecond):
			var err *model.AppError
			if progress, err = worker.ExtractBatch(progress); err != nil {
				mlog.Error("Worker: Failed to index batch for job", mlog.String("workername", worker.name), mlog.String("job_id", job.Id), mlog.Err(err))
				if err2 := worker.jobServer.SetJobError(job, err); err2 != nil {
					mlog.Error("Worker: Failed to set job error", mlog.String("workername", worker.name), mlog.String("job_id", job.Id), mlog.Err(err2), mlog.NamedErr("set_error", err))
				}
				return
			}

			if err := worker.jobServer.SetJobProgress(job, progress.CurrentProgress()); err != nil {
				mlog.Error("Worker: Failed to set progress for job", mlog.String("workername", worker.name), mlog.String("job_id", job.Id), mlog.Err(err))
				if err2 := worker.jobServer.SetJobError(job, err); err2 != nil {
					mlog.Error("Worker: Failed to set error for job", mlog.String("workername", worker.name), mlog.String("job_id", job.Id), mlog.Err(err2), mlog.NamedErr("set_error", err))
				}
				return
			}

			if progress.Done {
				if err := worker.jobServer.SetJobSuccess(job); err != nil {
					mlog.Error("Worker: Failed to set success for job", mlog.String("workername", worker.name), mlog.String("job_id", job.Id), mlog.Err(err))
					if err2 := worker.jobServer.SetJobError(job, err); err2 != nil {
						mlog.Error("Worker: Failed to set error for job", mlog.String("workername", worker.name), mlog.String("job_id", job.Id), mlog.Err(err2), mlog.NamedErr("set_error", err))
					}
				}
				mlog.Info("Worker: Indexing job finished successfully", mlog.String("workername", worker.name), mlog.String("job_id", job.Id))
				return
			}
		}
	}
}

func (worker *Worker) BulkExtractContentFiles(files []*model.FileForIndexing, progress ExtractContentProgress) (int64, *model.AppError) {
	lastCreateAt := int64(0)

	for _, file := range files {
		if file.DeleteAt == 0 {
			fileReader, err := worker.app.FileReader(file.Path)
			if err != nil {
				mlog.Error("Failed to extract file content", mlog.Err(err))
				continue
			}
			text, nErr := docextractor.Extract(file.Name, fileReader, docextractor.ExtractSettings{
				ArchiveRecursion: *worker.app.Config().FileSettings.ArchiveRecursion,
			})
			if nErr != nil {
				mlog.Error("Failed to extract file content", mlog.Err(nErr))
			}
			if storeErr := worker.jobServer.Store.FileInfo().SetContent(file.Id, text); storeErr != nil {
				mlog.Error("Failed to save the extracted file content", mlog.Err(storeErr))
			}
			fileReader.Close()
		}

		lastCreateAt = file.CreateAt
	}

	return lastCreateAt, nil
}

func (worker *Worker) ExtractBatch(progress ExtractContentProgress) (ExtractContentProgress, *model.AppError) {
	endTime := progress.LastFileTime

	var files []*model.FileForIndexing

	tries := 0
	for files == nil {
		var err error
		files, err = worker.jobServer.Store.FileInfo().GetFilesBatchForIndexing(progress.LastFileTime, endTime, BatchSize)
		if err != nil {
			if tries >= 10 {
				return progress, model.NewAppError("ExtractBatch", "app.post.get_files_batch_for_extract_content.get.app_error", nil, err.Error(), http.StatusInternalServerError)
			}
			mlog.Warn("Failed to get files batch for extract content. Retrying.", mlog.Err(err))

			// Wait a bit before trying again.
			time.Sleep(15 * time.Second)
		}

		tries++
	}

	newLastFileTime, err := worker.BulkExtractContentFiles(files, progress)
	if err != nil {
		return progress, err
	}

	// Due to the "endTime" parameter in the store query, we might get an incomplete batch before the end. In this
	// case, set the "newLastFileTime" to the endTime so we don't get stuck running the same query in a loop.
	if len(files) < BatchSize {
		newLastFileTime = endTime
	}

	// When to Stop: we index either until we pass a batch of messages where the last
	// message is created at or after the specified end time when setting up the batch
	// index, or until two consecutive full batches have the same end time of their final
	// messages. This second case is safe as long as the assumption that the database
	// cannot contain more messages with the same CreateAt time than the batch size holds.
	if progress.EndAtTime <= newLastFileTime {
		progress.Done = true
		progress.LastFileTime = progress.StartAtTime
	} else if progress.LastFileTime == newLastFileTime && len(files) == BatchSize {
		mlog.Warn("More files with the same CreateAt time were detected than the permitted batch size. Aborting indexing job.", mlog.Int64("CreateAt", newLastFileTime), mlog.Int("Batch Size", BatchSize))
		progress.Done = true
		progress.LastFileTime = progress.StartAtTime
	} else {
		progress.LastFileTime = newLastFileTime
	}

	progress.DoneCount += int64(len(files))

	return progress, nil
}

func (worker *Worker) setJobSuccess(job *model.Job) {
	if err := worker.app.Srv().Jobs.SetJobSuccess(job); err != nil {
		mlog.Error("Worker: Failed to set success for job", mlog.String("worker", worker.name), mlog.String("job_id", job.Id), mlog.String("error", err.Error()))
		worker.setJobError(job, err)
	}
}

func (worker *Worker) setJobError(job *model.Job, appError *model.AppError) {
	if err := worker.app.Srv().Jobs.SetJobError(job, appError); err != nil {
		mlog.Error("Worker: Failed to set job error", mlog.String("worker", worker.name), mlog.String("job_id", job.Id), mlog.String("error", err.Error()))
	}
}
