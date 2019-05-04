package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaTokenStore struct {
	baseStore store.TokenStore
	root      *KafkaStore
}

func (s KafkaTokenStore) Save(recovery *model.Token) store.StoreChannel {
	return s.baseStore.Save(recovery)
}

func (s KafkaTokenStore) Delete(token string) store.StoreChannel {
	return s.baseStore.Delete(token)
}

func (s KafkaTokenStore) GetByToken(token string) store.StoreChannel {
	return s.baseStore.GetByToken(token)
}

func (s KafkaTokenStore) Cleanup() {
	s.baseStore.Cleanup()
}

func (s KafkaTokenStore) RemoveAllTokensByType(tokenType string) store.StoreChannel {
	return s.baseStore.RemoveAllTokensByType(tokenType)
}
