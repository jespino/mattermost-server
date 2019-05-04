package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaTermsOfServiceStore struct {
	baseStore store.TermsOfServiceStore
	root      *KafkaStore
}

func (s KafkaTermsOfServiceStore) Save(termsOfService *model.TermsOfService) store.StoreChannel {
	return s.baseStore.Save(termsOfService)
}

func (s KafkaTermsOfServiceStore) GetLatest(allowFromCache bool) store.StoreChannel {
	return s.baseStore.GetLatest(allowFromCache)
}

func (s KafkaTermsOfServiceStore) Get(id string, allowFromCache bool) store.StoreChannel {
	return s.baseStore.Get(id, allowFromCache)
}
