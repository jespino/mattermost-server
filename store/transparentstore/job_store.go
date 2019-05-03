package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentJobStore struct {
	baseStore store.JobStore
}

func (s TransparentJobStore) Save(job *model.Job) store.StoreChannel {
	return s.baseStore.Save(job)
}

func (s TransparentJobStore) UpdateOptimistically(job *model.Job, currentStatus string) store.StoreChannel {
	return s.baseStore.UpdateOptimistically(job, currentStatus)
}

func (s TransparentJobStore) UpdateStatus(id string, status string) store.StoreChannel {
	return s.baseStore.UpdateStatus(id, status)
}

func (s TransparentJobStore) UpdateStatusOptimistically(id string, currentStatus string, newStatus string) store.StoreChannel {
	return s.baseStore.UpdateStatusOptimistically(id, currentStatus, newStatus)
}

func (s TransparentJobStore) Get(id string) store.StoreChannel {
	return s.baseStore.Get(id)
}

func (s TransparentJobStore) GetAllPage(offset int, limit int) store.StoreChannel {
	return s.baseStore.GetAllPage(offset, limit)
}

func (s TransparentJobStore) GetAllByType(jobType string) store.StoreChannel {
	return s.baseStore.GetAllByType(jobType)
}

func (s TransparentJobStore) GetAllByTypePage(jobType string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetAllByTypePage(jobType, offset, limit)
}

func (s TransparentJobStore) GetAllByStatus(status string) store.StoreChannel {
	return s.baseStore.GetAllByStatus(status)
}

func (s TransparentJobStore) GetNewestJobByStatusAndType(status string, jobType string) store.StoreChannel {
	return s.baseStore.GetNewestJobByStatusAndType(status, jobType)
}

func (s TransparentJobStore) GetCountByStatusAndType(status string, jobType string) store.StoreChannel {
	return s.baseStore.GetCountByStatusAndType(status, jobType)
}

func (s TransparentJobStore) Delete(id string) store.StoreChannel {
	return s.baseStore.Delete(id)
}
