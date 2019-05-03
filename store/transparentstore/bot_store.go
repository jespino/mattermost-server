package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentBotStore struct {
	baseStore store.BotStore
}

func (s TransparentBotStore) Get(userId string, includeDeleted bool) store.StoreChannel {
	return s.baseStore.Get(userId, includeDeleted)
}

func (s TransparentBotStore) GetAll(options *model.BotGetOptions) store.StoreChannel {
	return s.baseStore.GetAll(options)
}

func (s TransparentBotStore) Save(bot *model.Bot) store.StoreChannel {
	return s.baseStore.Save(bot)
}

func (s TransparentBotStore) Update(bot *model.Bot) store.StoreChannel {
	return s.baseStore.Update(bot)
}

func (s TransparentBotStore) PermanentDelete(userId string) store.StoreChannel {
	return s.baseStore.PermanentDelete(userId)
}
