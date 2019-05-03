package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentSessionStore struct {
	baseStore store.SessionStore
}

func (s TransparentSessionStore) Save(session *model.Session) store.StoreChannel {
	return s.baseStore.Save(session)
}

func (s TransparentSessionStore) Get(sessionIdOrToken string) store.StoreChannel {
	return s.baseStore.Get(sessionIdOrToken)
}

func (s TransparentSessionStore) GetSessions(userId string) store.StoreChannel {
	return s.baseStore.GetSessions(userId)
}

func (s TransparentSessionStore) GetSessionsWithActiveDeviceIds(userId string) store.StoreChannel {
	return s.baseStore.GetSessionsWithActiveDeviceIds(userId)
}

func (s TransparentSessionStore) Remove(sessionIdOrToken string) store.StoreChannel {
	return s.baseStore.Remove(sessionIdOrToken)
}

func (s TransparentSessionStore) RemoveAllSessions() store.StoreChannel {
	return s.baseStore.RemoveAllSessions()
}

func (s TransparentSessionStore) PermanentDeleteSessionsByUser(teamId string) store.StoreChannel {
	return s.baseStore.PermanentDeleteSessionsByUser(teamId)
}

func (s TransparentSessionStore) UpdateLastActivityAt(sessionId string, time int64) store.StoreChannel {
	return s.baseStore.UpdateLastActivityAt(sessionId, time)
}

func (s TransparentSessionStore) UpdateRoles(userId string, roles string) store.StoreChannel {
	return s.baseStore.UpdateRoles(userId, roles)
}

func (s TransparentSessionStore) UpdateDeviceId(id string, deviceId string, expiresAt int64) store.StoreChannel {
	return s.baseStore.UpdateDeviceId(id, deviceId, expiresAt)
}

func (s TransparentSessionStore) AnalyticsSessionCount() store.StoreChannel {
	return s.baseStore.AnalyticsSessionCount()
}

func (s TransparentSessionStore) Cleanup(expiryTime int64, batchSize int64) {
	s.baseStore.Cleanup(expiryTime, batchSize)
}
