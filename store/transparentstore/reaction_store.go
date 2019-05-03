package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentReactionStore struct {
	baseStore store.ReactionStore
}

func (s TransparentReactionStore) Save(reaction *model.Reaction) (*model.Reaction, *model.AppError) {
	return s.baseStore.Save(reaction)
}

func (s TransparentReactionStore) Delete(reaction *model.Reaction) (*model.Reaction, *model.AppError) {
	return s.baseStore.Delete(reaction)
}

func (s TransparentReactionStore) GetForPost(postId string, allowFromCache bool) ([]*model.Reaction, *model.AppError) {
	return s.baseStore.GetForPost(postId, allowFromCache)
}

func (s TransparentReactionStore) DeleteAllWithEmojiName(emojiName string) *model.AppError {
	return s.baseStore.DeleteAllWithEmojiName(emojiName)
}

func (s TransparentReactionStore) PermanentDeleteBatch(endTime int64, limit int64) (int64, *model.AppError) {
	return s.baseStore.PermanentDeleteBatch(endTime, limit)
}

func (s TransparentReactionStore) BulkGetForPosts(postIds []string) ([]*model.Reaction, *model.AppError) {
	return s.baseStore.BulkGetForPosts(postIds)
}
