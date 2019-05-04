package localcachestore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type LocalCacheReactionStore struct {
	baseStore store.ReactionStore
	rootStore *LocalCacheStore
}

func (s LocalCacheReactionStore) Save(reaction *model.Reaction) (*model.Reaction, *model.AppError) {
	defer s.rootStore.doInvalidateCacheCluster(s.rootStore.reactionCache, reaction.PostId)
	return s.baseStore.Save(reaction)
}

func (s LocalCacheReactionStore) Delete(reaction *model.Reaction) (*model.Reaction, *model.AppError) {
	defer s.rootStore.doInvalidateCacheCluster(s.rootStore.reactionCache, reaction.PostId)
	return s.baseStore.Delete(reaction)
}

func (s LocalCacheReactionStore) GetForPost(postId string, allowFromCache bool) ([]*model.Reaction, *model.AppError) {
	if reaction := s.rootStore.doStandardReadCache(s.rootStore.reactionCache, postId); reaction != nil {
		return reaction.([]*model.Reaction), nil
	}

	reaction, err := s.baseStore.GetForPost(postId, false)
	if err != nil {
		return nil, err
	}

	s.rootStore.doStandardAddToCache(s.rootStore.reactionCache, postId, reaction)

	return reaction, nil
}

func (s LocalCacheReactionStore) DeleteAllWithEmojiName(emojiName string) *model.AppError {
	// This could be improved. Right now we just clear the whole
	// cache because we don't have a way find what post Ids have this emoji name.
	defer s.rootStore.doClearCacheCluster(s.rootStore.reactionCache)
	return s.baseStore.DeleteAllWithEmojiName(emojiName)
}

func (s LocalCacheReactionStore) PermanentDeleteBatch(endTime int64, limit int64) (int64, *model.AppError) {
	// Don't bother to clear the cache as the posts will be gone anyway and the reactions being deleted will
	// expire from the cache in due course.
	return s.baseStore.PermanentDeleteBatch(endTime, limit)
}

func (s LocalCacheReactionStore) BulkGetForPosts(postIds []string) ([]*model.Reaction, *model.AppError) {
	// Ignoring this.
	return s.baseStore.BulkGetForPosts(postIds)
}
