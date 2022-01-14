// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package memstore

import (
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/store"
)

type MemOAuthStore struct {
}

func newMemOAuthStore() store.OAuthStore {
	return &MemOAuthStore{}
}

func (as *MemOAuthStore) SaveApp(app *model.OAuthApp) (*model.OAuthApp, error) {
	panic("not implemented")
}

func (as *MemOAuthStore) UpdateApp(app *model.OAuthApp) (*model.OAuthApp, error) {
	panic("not implemented")
}

func (as *MemOAuthStore) GetApp(id string) (*model.OAuthApp, error) {
	panic("not implemented")
}

func (as *MemOAuthStore) GetAppByUser(userId string, offset, limit int) ([]*model.OAuthApp, error) {
	panic("not implemented")
}

func (as *MemOAuthStore) GetApps(offset, limit int) ([]*model.OAuthApp, error) {
	panic("not implemented")
}

func (as *MemOAuthStore) GetAuthorizedApps(userId string, offset, limit int) ([]*model.OAuthApp, error) {
	panic("not implemented")
}

func (as *MemOAuthStore) DeleteApp(id string) error {
	panic("not implemented")
}

func (as *MemOAuthStore) SaveAccessData(accessData *model.AccessData) (*model.AccessData, error) {
	panic("not implemented")
}

func (as *MemOAuthStore) GetAccessData(token string) (*model.AccessData, error) {
	panic("not implemented")
}

func (as *MemOAuthStore) GetAccessDataByUserForApp(userID, clientID string) ([]*model.AccessData, error) {
	panic("not implemented")
}

func (as *MemOAuthStore) GetAccessDataByRefreshToken(token string) (*model.AccessData, error) {
	panic("not implemented")
}

func (as *MemOAuthStore) GetPreviousAccessData(userID, clientID string) (*model.AccessData, error) {
	panic("not implemented")
}

func (as *MemOAuthStore) UpdateAccessData(accessData *model.AccessData) (*model.AccessData, error) {
	panic("not implemented")
}

func (as *MemOAuthStore) RemoveAccessData(token string) error {
	panic("not implemented")
}

func (as *MemOAuthStore) RemoveAllAccessData() error {
	panic("not implemented")
}

func (as *MemOAuthStore) SaveAuthData(authData *model.AuthData) (*model.AuthData, error) {
	panic("not implemented")
}

func (as *MemOAuthStore) GetAuthData(code string) (*model.AuthData, error) {
	panic("not implemented")
}

func (as *MemOAuthStore) RemoveAuthData(code string) error {
	panic("not implemented")
}

func (as *MemOAuthStore) PermanentDeleteAuthDataByUser(userId string) error {
	panic("not implemented")
}
