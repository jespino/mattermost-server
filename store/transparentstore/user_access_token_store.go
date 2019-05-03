package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentUserAccessTokenStore struct {
	baseStore store.UserAccessTokenStore
}

func (s TransparentUserAccessTokenStore) Save(token *model.UserAccessToken) store.StoreChannel {
	return s.baseStore.Save(token)
}

func (s TransparentUserAccessTokenStore) Delete(tokenId string) store.StoreChannel {
	return s.baseStore.Delete(tokenId)
}

func (s TransparentUserAccessTokenStore) DeleteAllForUser(userId string) store.StoreChannel {
	return s.baseStore.DeleteAllForUser(userId)
}

func (s TransparentUserAccessTokenStore) Get(tokenId string) store.StoreChannel {
	return s.baseStore.Get(tokenId)
}

func (s TransparentUserAccessTokenStore) GetAll(offset int, limit int) store.StoreChannel {
	return s.baseStore.GetAll(offset, limit)
}

func (s TransparentUserAccessTokenStore) GetByToken(tokenString string) store.StoreChannel {
	return s.baseStore.GetByToken(tokenString)
}

func (s TransparentUserAccessTokenStore) GetByUser(userId string, page int, perPage int) store.StoreChannel {
	return s.baseStore.GetByUser(userId, page, perPage)
}

func (s TransparentUserAccessTokenStore) Search(term string) store.StoreChannel {
	return s.baseStore.Search(term)
}

func (s TransparentUserAccessTokenStore) UpdateTokenEnable(tokenId string) store.StoreChannel {
	return s.baseStore.UpdateTokenEnable(tokenId)
}

func (s TransparentUserAccessTokenStore) UpdateTokenDisable(tokenId string) store.StoreChannel {
	return s.baseStore.UpdateTokenDisable(tokenId)
}
