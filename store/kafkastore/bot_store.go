package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaBotStore struct {
	baseStore store.BotStore
	root      *KafkaStore
}

func (s KafkaBotStore) Get(userId string, includeDeleted bool) store.StoreChannel {
	return s.baseStore.Get(userId, includeDeleted)
}

func (s KafkaBotStore) GetAll(options *model.BotGetOptions) store.StoreChannel {
	return s.baseStore.GetAll(options)
}

func (s KafkaBotStore) Save(bot *model.Bot) store.StoreChannel {
	return s.baseStore.Save(bot)
}

func (s KafkaBotStore) Update(bot *model.Bot) store.StoreChannel {
	return s.baseStore.Update(bot)
}

func (s KafkaBotStore) PermanentDelete(userId string) store.StoreChannel {
	return s.baseStore.PermanentDelete(userId)
}
