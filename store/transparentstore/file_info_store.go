package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentFileInfoStore struct {
	baseStore store.FileInfoStore
}

func (s TransparentFileInfoStore) Save(info *model.FileInfo) store.StoreChannel {
	return s.baseStore.Save(info)
}

func (s TransparentFileInfoStore) Get(id string) store.StoreChannel {
	return s.baseStore.Get(id)
}

func (s TransparentFileInfoStore) GetByPath(path string) store.StoreChannel {
	return s.baseStore.GetByPath(path)
}

func (s TransparentFileInfoStore) GetForPost(postId string, readFromMaster bool, allowFromCache bool) store.StoreChannel {
	return s.baseStore.GetForPost(postId, readFromMaster, allowFromCache)
}

func (s TransparentFileInfoStore) GetForUser(userId string) store.StoreChannel {
	return s.baseStore.GetForUser(userId)
}

func (s TransparentFileInfoStore) InvalidateFileInfosForPostCache(postId string) {
	s.baseStore.InvalidateFileInfosForPostCache(postId)
}

func (s TransparentFileInfoStore) AttachToPost(fileId string, postId string, creatorId string) store.StoreChannel {
	return s.baseStore.AttachToPost(fileId, postId, creatorId)
}

func (s TransparentFileInfoStore) DeleteForPost(postId string) store.StoreChannel {
	return s.baseStore.DeleteForPost(postId)
}

func (s TransparentFileInfoStore) PermanentDelete(fileId string) store.StoreChannel {
	return s.baseStore.PermanentDelete(fileId)
}

func (s TransparentFileInfoStore) PermanentDeleteBatch(endTime int64, limit int64) store.StoreChannel {
	return s.baseStore.PermanentDeleteBatch(endTime, limit)
}

func (s TransparentFileInfoStore) PermanentDeleteByUser(userId string) store.StoreChannel {
	return s.baseStore.PermanentDeleteByUser(userId)
}

func (s TransparentFileInfoStore) ClearCaches() {
	s.baseStore.ClearCaches()
}
