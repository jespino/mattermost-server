package localcachestore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type LocalCacheReactionStore struct {
	store.ReactionStore
	rootStore *LocalCacheStore
}

func (s LocalCacheReactionStore) Save(reaction *model.Reaction) (*model.Reaction, *model.AppError) {
	defer s.rootStore.doInvalidateCacheCluster(s.rootStore.reactionCache, reaction.PostId)
	return s.ReactionStore.Save(reaction)
}

func (s LocalCacheReactionStore) Delete(reaction *model.Reaction) (*model.Reaction, *model.AppError) {
	defer s.rootStore.doInvalidateCacheCluster(s.rootStore.reactionCache, reaction.PostId)
	return s.ReactionStore.Delete(reaction)
}

func (s LocalCacheReactionStore) GetForPost(postId string, allowFromCache bool) ([]*model.Reaction, *model.AppError) {
	if reaction := s.rootStore.doStandardReadCache(s.rootStore.reactionCache, postId); reaction != nil {
		return reaction.([]*model.Reaction), nil
	}

	reaction, err := s.ReactionStore.GetForPost(postId, false)
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
	return s.ReactionStore.DeleteAllWithEmojiName(emojiName)
}
