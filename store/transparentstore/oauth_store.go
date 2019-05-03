package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentOAuthStore struct {
	baseStore store.OAuthStore
}

func (s TransparentOAuthStore) SaveApp(app *model.OAuthApp) store.StoreChannel {
	return s.baseStore.SaveApp(app)
}

func (s TransparentOAuthStore) UpdateApp(app *model.OAuthApp) store.StoreChannel {
	return s.baseStore.UpdateApp(app)
}

func (s TransparentOAuthStore) GetApp(id string) store.StoreChannel {
	return s.baseStore.GetApp(id)
}

func (s TransparentOAuthStore) GetAppByUser(userId string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetAppByUser(userId, offset, limit)
}

func (s TransparentOAuthStore) GetApps(offset int, limit int) store.StoreChannel {
	return s.baseStore.GetApps(offset, limit)
}

func (s TransparentOAuthStore) GetAuthorizedApps(userId string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetAuthorizedApps(userId, offset, limit)
}

func (s TransparentOAuthStore) DeleteApp(id string) store.StoreChannel {
	return s.baseStore.DeleteApp(id)
}

func (s TransparentOAuthStore) SaveAuthData(authData *model.AuthData) store.StoreChannel {
	return s.baseStore.SaveAuthData(authData)
}

func (s TransparentOAuthStore) GetAuthData(code string) store.StoreChannel {
	return s.baseStore.GetAuthData(code)
}

func (s TransparentOAuthStore) RemoveAuthData(code string) store.StoreChannel {
	return s.baseStore.RemoveAuthData(code)
}

func (s TransparentOAuthStore) PermanentDeleteAuthDataByUser(userId string) store.StoreChannel {
	return s.baseStore.PermanentDeleteAuthDataByUser(userId)
}

func (s TransparentOAuthStore) SaveAccessData(accessData *model.AccessData) store.StoreChannel {
	return s.baseStore.SaveAccessData(accessData)
}

func (s TransparentOAuthStore) UpdateAccessData(accessData *model.AccessData) store.StoreChannel {
	return s.baseStore.UpdateAccessData(accessData)
}

func (s TransparentOAuthStore) GetAccessData(token string) store.StoreChannel {
	return s.baseStore.GetAccessData(token)
}

func (s TransparentOAuthStore) GetAccessDataByUserForApp(userId string, clientId string) store.StoreChannel {
	return s.baseStore.GetAccessDataByUserForApp(userId, clientId)
}

func (s TransparentOAuthStore) GetAccessDataByRefreshToken(token string) store.StoreChannel {
	return s.baseStore.GetAccessDataByRefreshToken(token)
}

func (s TransparentOAuthStore) GetPreviousAccessData(userId string, clientId string) store.StoreChannel {
	return s.baseStore.GetPreviousAccessData(userId, clientId)
}

func (s TransparentOAuthStore) RemoveAccessData(token string) store.StoreChannel {
	return s.baseStore.RemoveAccessData(token)
}
