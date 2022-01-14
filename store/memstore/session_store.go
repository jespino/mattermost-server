// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package memstore

import (
	"context"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/store"
)

type MemSessionStore struct{}

func newMemSessionStore() store.SessionStore {
	return &MemSessionStore{}
}

func (me *MemSessionStore) Save(session *model.Session) (*model.Session, error) {
	panic("not implemented")
}

func (me *MemSessionStore) Get(ctx context.Context, sessionIdOrToken string) (*model.Session, error) {
	panic("not implemented")
}

func (me *MemSessionStore) GetSessions(userId string) ([]*model.Session, error) {
	panic("not implemented")
}

func (me *MemSessionStore) GetSessionsWithActiveDeviceIds(userId string) ([]*model.Session, error) {
	panic("not implemented")
}

func (me *MemSessionStore) GetSessionsExpired(thresholdMillis int64, mobileOnly bool, unnotifiedOnly bool) ([]*model.Session, error) {
	panic("not implemented")
}

func (me *MemSessionStore) UpdateExpiredNotify(sessionId string, notified bool) error {
	panic("not implemented")
}

func (me *MemSessionStore) Remove(sessionIdOrToken string) error {
	panic("not implemented")
}

func (me *MemSessionStore) RemoveAllSessions() error {
	panic("not implemented")
}

func (me *MemSessionStore) PermanentDeleteSessionsByUser(userId string) error {
	panic("not implemented")
}

func (me *MemSessionStore) UpdateExpiresAt(sessionId string, time int64) error {
	panic("not implemented")
}

func (me *MemSessionStore) UpdateLastActivityAt(sessionId string, time int64) error {
	panic("not implemented")
}

func (me *MemSessionStore) UpdateRoles(userId, roles string) (string, error) {
	panic("not implemented")
}

func (me *MemSessionStore) UpdateDeviceId(id string, deviceId string, expiresAt int64) (string, error) {
	panic("not implemented")
}

func (me *MemSessionStore) UpdateProps(session *model.Session) error {
	panic("not implemented")
}

func (me *MemSessionStore) AnalyticsSessionCount() (int64, error) {
	panic("not implemented")
}

func (me *MemSessionStore) Cleanup(expiryTime int64, batchSize int64) error {
	panic("not implemented")
}
