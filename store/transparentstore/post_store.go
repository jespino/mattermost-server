package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentPostStore struct {
	baseStore store.PostStore
}

func (s TransparentPostStore) Save(post *model.Post) store.StoreChannel {
	return s.baseStore.Save(post)
}

func (s TransparentPostStore) Update(newPost *model.Post, oldPost *model.Post) store.StoreChannel {
	return s.baseStore.Update(newPost, oldPost)
}

func (s TransparentPostStore) Get(id string) store.StoreChannel {
	return s.baseStore.Get(id)
}

func (s TransparentPostStore) GetSingle(id string) store.StoreChannel {
	return s.baseStore.GetSingle(id)
}

func (s TransparentPostStore) Delete(postId string, time int64, deleteByID string) store.StoreChannel {
	return s.baseStore.Delete(postId, time, deleteByID)
}

func (s TransparentPostStore) PermanentDeleteByUser(userId string) store.StoreChannel {
	return s.baseStore.PermanentDeleteByUser(userId)
}

func (s TransparentPostStore) PermanentDeleteByChannel(channelId string) store.StoreChannel {
	return s.baseStore.PermanentDeleteByChannel(channelId)
}

func (s TransparentPostStore) GetPosts(channelId string, offset int, limit int, allowFromCache bool) store.StoreChannel {
	return s.baseStore.GetPosts(channelId, offset, limit, allowFromCache)
}

func (s TransparentPostStore) GetFlaggedPosts(userId string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetFlaggedPosts(userId, offset, limit)
}

func (s TransparentPostStore) GetFlaggedPostsForTeam(userId string, teamId string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetFlaggedPostsForTeam(userId, teamId, offset, limit)
}

func (s TransparentPostStore) GetFlaggedPostsForChannel(userId string, channelId string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetFlaggedPostsForChannel(userId, channelId, offset, limit)
}

func (s TransparentPostStore) GetPostsBefore(channelId string, postId string, numPosts int, offset int) store.StoreChannel {
	return s.baseStore.GetPostsBefore(channelId, postId, numPosts, offset)
}

func (s TransparentPostStore) GetPostsAfter(channelId string, postId string, numPosts int, offset int) store.StoreChannel {
	return s.baseStore.GetPostsAfter(channelId, postId, numPosts, offset)
}

func (s TransparentPostStore) GetPostsSince(channelId string, time int64, allowFromCache bool) store.StoreChannel {
	return s.baseStore.GetPostsSince(channelId, time, allowFromCache)
}

func (s TransparentPostStore) GetEtag(channelId string, allowFromCache bool) store.StoreChannel {
	return s.baseStore.GetEtag(channelId, allowFromCache)
}

func (s TransparentPostStore) Search(teamId string, userId string, params *model.SearchParams) store.StoreChannel {
	return s.baseStore.Search(teamId, userId, params)
}

func (s TransparentPostStore) AnalyticsUserCountsWithPostsByDay(teamId string) store.StoreChannel {
	return s.baseStore.AnalyticsUserCountsWithPostsByDay(teamId)
}

func (s TransparentPostStore) AnalyticsPostCountsByDay(teamId string) store.StoreChannel {
	return s.baseStore.AnalyticsPostCountsByDay(teamId)
}

func (s TransparentPostStore) AnalyticsPostCount(teamId string, mustHaveFile bool, mustHaveHashtag bool) store.StoreChannel {
	return s.baseStore.AnalyticsPostCount(teamId, mustHaveFile, mustHaveHashtag)
}

func (s TransparentPostStore) ClearCaches() {
	s.baseStore.ClearCaches()
}

func (s TransparentPostStore) InvalidateLastPostTimeCache(channelId string) {
	s.baseStore.InvalidateLastPostTimeCache(channelId)
}

func (s TransparentPostStore) GetPostsCreatedAt(channelId string, time int64) store.StoreChannel {
	return s.baseStore.GetPostsCreatedAt(channelId, time)
}

func (s TransparentPostStore) Overwrite(post *model.Post) store.StoreChannel {
	return s.baseStore.Overwrite(post)
}

func (s TransparentPostStore) GetPostsByIds(postIds []string) store.StoreChannel {
	return s.baseStore.GetPostsByIds(postIds)
}

func (s TransparentPostStore) GetPostsBatchForIndexing(startTime int64, endTime int64, limit int) store.StoreChannel {
	return s.baseStore.GetPostsBatchForIndexing(startTime, endTime, limit)
}

func (s TransparentPostStore) PermanentDeleteBatch(endTime int64, limit int64) store.StoreChannel {
	return s.baseStore.PermanentDeleteBatch(endTime, limit)
}

func (s TransparentPostStore) GetOldest() store.StoreChannel {
	return s.baseStore.GetOldest()
}

func (s TransparentPostStore) GetMaxPostSize() store.StoreChannel {
	return s.baseStore.GetMaxPostSize()
}

func (s TransparentPostStore) GetParentsForExportAfter(limit int, afterId string) store.StoreChannel {
	return s.baseStore.GetParentsForExportAfter(limit, afterId)
}

func (s TransparentPostStore) GetRepliesForExport(parentId string) store.StoreChannel {
	return s.baseStore.GetRepliesForExport(parentId)
}

func (s TransparentPostStore) GetDirectPostParentsForExportAfter(limit int, afterId string) store.StoreChannel {
	return s.baseStore.GetDirectPostParentsForExportAfter(limit, afterId)
}
