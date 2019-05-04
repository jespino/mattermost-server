package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaFileInfoStore struct {
	baseStore store.FileInfoStore
	root      *KafkaStore
}

func (s KafkaFileInfoStore) Save(info *model.FileInfo) store.StoreChannel {
	return s.baseStore.Save(info)
}

func (s KafkaFileInfoStore) Get(id string) store.StoreChannel {
	return s.baseStore.Get(id)
}

func (s KafkaFileInfoStore) GetByPath(path string) store.StoreChannel {
	return s.baseStore.GetByPath(path)
}

func (s KafkaFileInfoStore) GetForPost(postId string, readFromMaster bool, allowFromCache bool) store.StoreChannel {
	return s.baseStore.GetForPost(postId, readFromMaster, allowFromCache)
}

func (s KafkaFileInfoStore) GetForUser(userId string) store.StoreChannel {
	return s.baseStore.GetForUser(userId)
}

func (s KafkaFileInfoStore) InvalidateFileInfosForPostCache(postId string) {
	s.baseStore.InvalidateFileInfosForPostCache(postId)
}

func (s KafkaFileInfoStore) AttachToPost(fileId string, postId string, creatorId string) store.StoreChannel {
	return s.baseStore.AttachToPost(fileId, postId, creatorId)
}

func (s KafkaFileInfoStore) DeleteForPost(postId string) store.StoreChannel {
	return s.baseStore.DeleteForPost(postId)
}

func (s KafkaFileInfoStore) PermanentDelete(fileId string) store.StoreChannel {
	return s.baseStore.PermanentDelete(fileId)
}

func (s KafkaFileInfoStore) PermanentDeleteBatch(endTime int64, limit int64) store.StoreChannel {
	return s.baseStore.PermanentDeleteBatch(endTime, limit)
}

func (s KafkaFileInfoStore) PermanentDeleteByUser(userId string) store.StoreChannel {
	return s.baseStore.PermanentDeleteByUser(userId)
}

func (s KafkaFileInfoStore) ClearCaches() {
	s.baseStore.ClearCaches()
}
