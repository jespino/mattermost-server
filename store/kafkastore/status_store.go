package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaStatusStore struct {
	baseStore store.StatusStore
	root      *KafkaStore
}

func (s KafkaStatusStore) SaveOrUpdate(status *model.Status) store.StoreChannel {
	return s.baseStore.SaveOrUpdate(status)
}

func (s KafkaStatusStore) Get(userId string) store.StoreChannel {
	return s.baseStore.Get(userId)
}

func (s KafkaStatusStore) GetByIds(userIds []string) store.StoreChannel {
	return s.baseStore.GetByIds(userIds)
}

func (s KafkaStatusStore) GetOnlineAway() store.StoreChannel {
	return s.baseStore.GetOnlineAway()
}

func (s KafkaStatusStore) GetOnline() store.StoreChannel {
	return s.baseStore.GetOnline()
}

func (s KafkaStatusStore) GetAllFromTeam(teamId string) store.StoreChannel {
	return s.baseStore.GetAllFromTeam(teamId)
}

func (s KafkaStatusStore) ResetAll() store.StoreChannel {
	return s.baseStore.ResetAll()
}

func (s KafkaStatusStore) GetTotalActiveUsersCount() store.StoreChannel {
	return s.baseStore.GetTotalActiveUsersCount()
}

func (s KafkaStatusStore) UpdateLastActivityAt(userId string, lastActivityAt int64) store.StoreChannel {
	return s.baseStore.UpdateLastActivityAt(userId, lastActivityAt)
}
