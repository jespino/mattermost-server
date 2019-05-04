package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaSessionStore struct {
	baseStore store.SessionStore
	root      *KafkaStore
}

func (s KafkaSessionStore) Save(session *model.Session) store.StoreChannel {
	return s.baseStore.Save(session)
}

func (s KafkaSessionStore) Get(sessionIdOrToken string) store.StoreChannel {
	return s.baseStore.Get(sessionIdOrToken)
}

func (s KafkaSessionStore) GetSessions(userId string) store.StoreChannel {
	return s.baseStore.GetSessions(userId)
}

func (s KafkaSessionStore) GetSessionsWithActiveDeviceIds(userId string) store.StoreChannel {
	return s.baseStore.GetSessionsWithActiveDeviceIds(userId)
}

func (s KafkaSessionStore) Remove(sessionIdOrToken string) store.StoreChannel {
	return s.baseStore.Remove(sessionIdOrToken)
}

func (s KafkaSessionStore) RemoveAllSessions() store.StoreChannel {
	return s.baseStore.RemoveAllSessions()
}

func (s KafkaSessionStore) PermanentDeleteSessionsByUser(teamId string) store.StoreChannel {
	return s.baseStore.PermanentDeleteSessionsByUser(teamId)
}

func (s KafkaSessionStore) UpdateLastActivityAt(sessionId string, time int64) store.StoreChannel {
	return s.baseStore.UpdateLastActivityAt(sessionId, time)
}

func (s KafkaSessionStore) UpdateRoles(userId string, roles string) store.StoreChannel {
	return s.baseStore.UpdateRoles(userId, roles)
}

func (s KafkaSessionStore) UpdateDeviceId(id string, deviceId string, expiresAt int64) store.StoreChannel {
	return s.baseStore.UpdateDeviceId(id, deviceId, expiresAt)
}

func (s KafkaSessionStore) AnalyticsSessionCount() store.StoreChannel {
	return s.baseStore.AnalyticsSessionCount()
}

func (s KafkaSessionStore) Cleanup(expiryTime int64, batchSize int64) {
	s.baseStore.Cleanup(expiryTime, batchSize)
}
