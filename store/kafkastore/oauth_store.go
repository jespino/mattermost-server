package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaOAuthStore struct {
	baseStore store.OAuthStore
	root      *KafkaStore
}

func (s KafkaOAuthStore) SaveApp(app *model.OAuthApp) store.StoreChannel {
	return s.baseStore.SaveApp(app)
}

func (s KafkaOAuthStore) UpdateApp(app *model.OAuthApp) store.StoreChannel {
	return s.baseStore.UpdateApp(app)
}

func (s KafkaOAuthStore) GetApp(id string) store.StoreChannel {
	return s.baseStore.GetApp(id)
}

func (s KafkaOAuthStore) GetAppByUser(userId string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetAppByUser(userId, offset, limit)
}

func (s KafkaOAuthStore) GetApps(offset int, limit int) store.StoreChannel {
	return s.baseStore.GetApps(offset, limit)
}

func (s KafkaOAuthStore) GetAuthorizedApps(userId string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetAuthorizedApps(userId, offset, limit)
}

func (s KafkaOAuthStore) DeleteApp(id string) store.StoreChannel {
	return s.baseStore.DeleteApp(id)
}

func (s KafkaOAuthStore) SaveAuthData(authData *model.AuthData) store.StoreChannel {
	return s.baseStore.SaveAuthData(authData)
}

func (s KafkaOAuthStore) GetAuthData(code string) store.StoreChannel {
	return s.baseStore.GetAuthData(code)
}

func (s KafkaOAuthStore) RemoveAuthData(code string) store.StoreChannel {
	return s.baseStore.RemoveAuthData(code)
}

func (s KafkaOAuthStore) PermanentDeleteAuthDataByUser(userId string) store.StoreChannel {
	return s.baseStore.PermanentDeleteAuthDataByUser(userId)
}

func (s KafkaOAuthStore) SaveAccessData(accessData *model.AccessData) store.StoreChannel {
	return s.baseStore.SaveAccessData(accessData)
}

func (s KafkaOAuthStore) UpdateAccessData(accessData *model.AccessData) store.StoreChannel {
	return s.baseStore.UpdateAccessData(accessData)
}

func (s KafkaOAuthStore) GetAccessData(token string) store.StoreChannel {
	return s.baseStore.GetAccessData(token)
}

func (s KafkaOAuthStore) GetAccessDataByUserForApp(userId string, clientId string) store.StoreChannel {
	return s.baseStore.GetAccessDataByUserForApp(userId, clientId)
}

func (s KafkaOAuthStore) GetAccessDataByRefreshToken(token string) store.StoreChannel {
	return s.baseStore.GetAccessDataByRefreshToken(token)
}

func (s KafkaOAuthStore) GetPreviousAccessData(userId string, clientId string) store.StoreChannel {
	return s.baseStore.GetPreviousAccessData(userId, clientId)
}

func (s KafkaOAuthStore) RemoveAccessData(token string) store.StoreChannel {
	return s.baseStore.RemoveAccessData(token)
}
