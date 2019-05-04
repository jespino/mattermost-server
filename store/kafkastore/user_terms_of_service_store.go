package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaUserTermsOfServiceStore struct {
	baseStore store.UserTermsOfServiceStore
	root      *KafkaStore
}

func (s KafkaUserTermsOfServiceStore) GetByUser(userId string) store.StoreChannel {
	return s.baseStore.GetByUser(userId)
}

func (s KafkaUserTermsOfServiceStore) Save(userTermsOfService *model.UserTermsOfService) store.StoreChannel {
	return s.baseStore.Save(userTermsOfService)
}

func (s KafkaUserTermsOfServiceStore) Delete(userId string, termsOfServiceId string) store.StoreChannel {
	return s.baseStore.Delete(userId, termsOfServiceId)
}
