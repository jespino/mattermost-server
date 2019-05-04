package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaUserAccessTokenStore struct {
	baseStore store.UserAccessTokenStore
	root      *KafkaStore
}

func (s KafkaUserAccessTokenStore) Save(token *model.UserAccessToken) store.StoreChannel {
	return s.baseStore.Save(token)
}

func (s KafkaUserAccessTokenStore) Delete(tokenId string) store.StoreChannel {
	return s.baseStore.Delete(tokenId)
}

func (s KafkaUserAccessTokenStore) DeleteAllForUser(userId string) store.StoreChannel {
	return s.baseStore.DeleteAllForUser(userId)
}

func (s KafkaUserAccessTokenStore) Get(tokenId string) store.StoreChannel {
	return s.baseStore.Get(tokenId)
}

func (s KafkaUserAccessTokenStore) GetAll(offset int, limit int) store.StoreChannel {
	return s.baseStore.GetAll(offset, limit)
}

func (s KafkaUserAccessTokenStore) GetByToken(tokenString string) store.StoreChannel {
	return s.baseStore.GetByToken(tokenString)
}

func (s KafkaUserAccessTokenStore) GetByUser(userId string, page int, perPage int) store.StoreChannel {
	return s.baseStore.GetByUser(userId, page, perPage)
}

func (s KafkaUserAccessTokenStore) Search(term string) store.StoreChannel {
	return s.baseStore.Search(term)
}

func (s KafkaUserAccessTokenStore) UpdateTokenEnable(tokenId string) store.StoreChannel {
	return s.baseStore.UpdateTokenEnable(tokenId)
}

func (s KafkaUserAccessTokenStore) UpdateTokenDisable(tokenId string) store.StoreChannel {
	return s.baseStore.UpdateTokenDisable(tokenId)
}
