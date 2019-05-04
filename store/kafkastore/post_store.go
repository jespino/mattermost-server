package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaPostStore struct {
	baseStore store.PostStore
	root      *KafkaStore
}

func (s KafkaPostStore) Save(post *model.Post) store.StoreChannel {
	s.root.sendMessage("Post.Save", map[string]interface{}{"post": post})
	return s.baseStore.Save(post)
}

func (s KafkaPostStore) Update(newPost *model.Post, oldPost *model.Post) store.StoreChannel {
	return s.baseStore.Update(newPost, oldPost)
}

func (s KafkaPostStore) Get(id string) store.StoreChannel {
	return s.baseStore.Get(id)
}

func (s KafkaPostStore) GetSingle(id string) store.StoreChannel {
	return s.baseStore.GetSingle(id)
}

func (s KafkaPostStore) Delete(postId string, time int64, deleteByID string) store.StoreChannel {
	return s.baseStore.Delete(postId, time, deleteByID)
}

func (s KafkaPostStore) PermanentDeleteByUser(userId string) store.StoreChannel {
	return s.baseStore.PermanentDeleteByUser(userId)
}

func (s KafkaPostStore) PermanentDeleteByChannel(channelId string) store.StoreChannel {
	return s.baseStore.PermanentDeleteByChannel(channelId)
}

func (s KafkaPostStore) GetPosts(channelId string, offset int, limit int, allowFromCache bool) store.StoreChannel {
	return s.baseStore.GetPosts(channelId, offset, limit, allowFromCache)
}

func (s KafkaPostStore) GetFlaggedPosts(userId string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetFlaggedPosts(userId, offset, limit)
}

func (s KafkaPostStore) GetFlaggedPostsForTeam(userId string, teamId string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetFlaggedPostsForTeam(userId, teamId, offset, limit)
}

func (s KafkaPostStore) GetFlaggedPostsForChannel(userId string, channelId string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetFlaggedPostsForChannel(userId, channelId, offset, limit)
}

func (s KafkaPostStore) GetPostsBefore(channelId string, postId string, numPosts int, offset int) store.StoreChannel {
	return s.baseStore.GetPostsBefore(channelId, postId, numPosts, offset)
}

func (s KafkaPostStore) GetPostsAfter(channelId string, postId string, numPosts int, offset int) store.StoreChannel {
	return s.baseStore.GetPostsAfter(channelId, postId, numPosts, offset)
}

func (s KafkaPostStore) GetPostsSince(channelId string, time int64, allowFromCache bool) store.StoreChannel {
	return s.baseStore.GetPostsSince(channelId, time, allowFromCache)
}

func (s KafkaPostStore) GetEtag(channelId string, allowFromCache bool) store.StoreChannel {
	return s.baseStore.GetEtag(channelId, allowFromCache)
}

func (s KafkaPostStore) Search(teamId string, userId string, params *model.SearchParams) store.StoreChannel {
	return s.baseStore.Search(teamId, userId, params)
}

func (s KafkaPostStore) AnalyticsUserCountsWithPostsByDay(teamId string) store.StoreChannel {
	return s.baseStore.AnalyticsUserCountsWithPostsByDay(teamId)
}

func (s KafkaPostStore) AnalyticsPostCountsByDay(teamId string) store.StoreChannel {
	return s.baseStore.AnalyticsPostCountsByDay(teamId)
}

func (s KafkaPostStore) AnalyticsPostCount(teamId string, mustHaveFile bool, mustHaveHashtag bool) store.StoreChannel {
	return s.baseStore.AnalyticsPostCount(teamId, mustHaveFile, mustHaveHashtag)
}

func (s KafkaPostStore) ClearCaches() {
	s.baseStore.ClearCaches()
}

func (s KafkaPostStore) InvalidateLastPostTimeCache(channelId string) {
	s.baseStore.InvalidateLastPostTimeCache(channelId)
}

func (s KafkaPostStore) GetPostsCreatedAt(channelId string, time int64) store.StoreChannel {
	return s.baseStore.GetPostsCreatedAt(channelId, time)
}

func (s KafkaPostStore) Overwrite(post *model.Post) store.StoreChannel {
	return s.baseStore.Overwrite(post)
}

func (s KafkaPostStore) GetPostsByIds(postIds []string) store.StoreChannel {
	return s.baseStore.GetPostsByIds(postIds)
}

func (s KafkaPostStore) GetPostsBatchForIndexing(startTime int64, endTime int64, limit int) store.StoreChannel {
	return s.baseStore.GetPostsBatchForIndexing(startTime, endTime, limit)
}

func (s KafkaPostStore) PermanentDeleteBatch(endTime int64, limit int64) store.StoreChannel {
	return s.baseStore.PermanentDeleteBatch(endTime, limit)
}

func (s KafkaPostStore) GetOldest() store.StoreChannel {
	return s.baseStore.GetOldest()
}

func (s KafkaPostStore) GetMaxPostSize() store.StoreChannel {
	return s.baseStore.GetMaxPostSize()
}

func (s KafkaPostStore) GetParentsForExportAfter(limit int, afterId string) store.StoreChannel {
	return s.baseStore.GetParentsForExportAfter(limit, afterId)
}

func (s KafkaPostStore) GetRepliesForExport(parentId string) store.StoreChannel {
	return s.baseStore.GetRepliesForExport(parentId)
}

func (s KafkaPostStore) GetDirectPostParentsForExportAfter(limit int, afterId string) store.StoreChannel {
	return s.baseStore.GetDirectPostParentsForExportAfter(limit, afterId)
}
