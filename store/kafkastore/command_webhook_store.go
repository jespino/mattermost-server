package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaCommandWebhookStore struct {
	baseStore store.CommandWebhookStore
	root      *KafkaStore
}

func (s KafkaCommandWebhookStore) Save(webhook *model.CommandWebhook) store.StoreChannel {
	return s.baseStore.Save(webhook)
}

func (s KafkaCommandWebhookStore) Get(id string) store.StoreChannel {
	return s.baseStore.Get(id)
}

func (s KafkaCommandWebhookStore) TryUse(id string, limit int) store.StoreChannel {
	return s.baseStore.TryUse(id, limit)
}

func (s KafkaCommandWebhookStore) Cleanup() {
	s.baseStore.Cleanup()
}
