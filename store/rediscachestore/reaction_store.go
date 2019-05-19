package rediscachestore

import (
	"github.com/go-redis/redis"
	"github.com/mattermost/mattermost-server/mlog"
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type RedisCacheReactionStore struct {
	store.ReactionStore
	client    *redis.Client
	rootStore *RedisCacheStore
}

func (s RedisCacheReactionStore) Save(reaction *model.Reaction) (*model.Reaction, *model.AppError) {
	if err := s.client.Del("reactions:" + reaction.PostId).Err(); err != nil {
		mlog.Error("Redis failed to remove key reactions:" + reaction.PostId + " Error: " + err.Error())
	}
	return s.ReactionStore.Save(reaction)
}

func (s RedisCacheReactionStore) Delete(reaction *model.Reaction) (*model.Reaction, *model.AppError) {
	if err := s.client.Del("reactions:" + reaction.PostId).Err(); err != nil {
		mlog.Error("Redis failed to remove key reactions:" + reaction.PostId + " Error: " + err.Error())
	}
	return s.ReactionStore.Delete(reaction)
}

func (s RedisCacheReactionStore) GetForPost(postId string, allowFromCache bool) ([]*model.Reaction, *model.AppError) {
	var reaction []*model.Reaction
	found, err := s.rootStore.load("reactions:"+postId, &reaction)
	if found {
		return reaction, nil
	}

	if err != nil {
		mlog.Error("Redis encountered an error on read: " + err.Error())
	}

	reaction, appErr := s.ReactionStore.GetForPost(postId, false)
	if appErr != nil {
		return nil, appErr
	}

	if err := s.rootStore.save("reactions:"+postId, reaction, REDIS_EXPIRY_TIME); err != nil {
		mlog.Error("Redis encountered and error on write: " + err.Error())
	}

	return reaction, nil
}
