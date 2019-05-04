package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaReactionStore struct {
	baseStore store.ReactionStore
	root      *KafkaStore
}

func (s KafkaReactionStore) Save(reaction *model.Reaction) (*model.Reaction, *model.AppError) {
	return s.baseStore.Save(reaction)
}

func (s KafkaReactionStore) Delete(reaction *model.Reaction) (*model.Reaction, *model.AppError) {
	return s.baseStore.Delete(reaction)
}

func (s KafkaReactionStore) GetForPost(postId string, allowFromCache bool) ([]*model.Reaction, *model.AppError) {
	return s.baseStore.GetForPost(postId, allowFromCache)
}

func (s KafkaReactionStore) DeleteAllWithEmojiName(emojiName string) *model.AppError {
	return s.baseStore.DeleteAllWithEmojiName(emojiName)
}

func (s KafkaReactionStore) PermanentDeleteBatch(endTime int64, limit int64) (int64, *model.AppError) {
	return s.baseStore.PermanentDeleteBatch(endTime, limit)
}

func (s KafkaReactionStore) BulkGetForPosts(postIds []string) ([]*model.Reaction, *model.AppError) {
	return s.baseStore.BulkGetForPosts(postIds)
}
