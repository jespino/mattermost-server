package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaJobStore struct {
	baseStore store.JobStore
	root      *KafkaStore
}

func (s KafkaJobStore) Save(job *model.Job) store.StoreChannel {
	return s.baseStore.Save(job)
}

func (s KafkaJobStore) UpdateOptimistically(job *model.Job, currentStatus string) store.StoreChannel {
	return s.baseStore.UpdateOptimistically(job, currentStatus)
}

func (s KafkaJobStore) UpdateStatus(id string, status string) store.StoreChannel {
	return s.baseStore.UpdateStatus(id, status)
}

func (s KafkaJobStore) UpdateStatusOptimistically(id string, currentStatus string, newStatus string) store.StoreChannel {
	return s.baseStore.UpdateStatusOptimistically(id, currentStatus, newStatus)
}

func (s KafkaJobStore) Get(id string) store.StoreChannel {
	return s.baseStore.Get(id)
}

func (s KafkaJobStore) GetAllPage(offset int, limit int) store.StoreChannel {
	return s.baseStore.GetAllPage(offset, limit)
}

func (s KafkaJobStore) GetAllByType(jobType string) store.StoreChannel {
	return s.baseStore.GetAllByType(jobType)
}

func (s KafkaJobStore) GetAllByTypePage(jobType string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetAllByTypePage(jobType, offset, limit)
}

func (s KafkaJobStore) GetAllByStatus(status string) store.StoreChannel {
	return s.baseStore.GetAllByStatus(status)
}

func (s KafkaJobStore) GetNewestJobByStatusAndType(status string, jobType string) store.StoreChannel {
	return s.baseStore.GetNewestJobByStatusAndType(status, jobType)
}

func (s KafkaJobStore) GetCountByStatusAndType(status string, jobType string) store.StoreChannel {
	return s.baseStore.GetCountByStatusAndType(status, jobType)
}

func (s KafkaJobStore) Delete(id string) store.StoreChannel {
	return s.baseStore.Delete(id)
}
