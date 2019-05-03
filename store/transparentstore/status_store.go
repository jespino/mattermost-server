package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentStatusStore struct {
	baseStore store.StatusStore
}

func (s TransparentStatusStore) SaveOrUpdate(status *model.Status) store.StoreChannel {
	return s.baseStore.SaveOrUpdate(status)
}

func (s TransparentStatusStore) Get(userId string) store.StoreChannel {
	return s.baseStore.Get(userId)
}

func (s TransparentStatusStore) GetByIds(userIds []string) store.StoreChannel {
	return s.baseStore.GetByIds(userIds)
}

func (s TransparentStatusStore) GetOnlineAway() store.StoreChannel {
	return s.baseStore.GetOnlineAway()
}

func (s TransparentStatusStore) GetOnline() store.StoreChannel {
	return s.baseStore.GetOnline()
}

func (s TransparentStatusStore) GetAllFromTeam(teamId string) store.StoreChannel {
	return s.baseStore.GetAllFromTeam(teamId)
}

func (s TransparentStatusStore) ResetAll() store.StoreChannel {
	return s.baseStore.ResetAll()
}

func (s TransparentStatusStore) GetTotalActiveUsersCount() store.StoreChannel {
	return s.baseStore.GetTotalActiveUsersCount()
}

func (s TransparentStatusStore) UpdateLastActivityAt(userId string, lastActivityAt int64) store.StoreChannel {
	return s.baseStore.UpdateLastActivityAt(userId, lastActivityAt)
}
