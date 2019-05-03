package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentCommandStore struct {
	baseStore store.CommandStore
}

func (s TransparentCommandStore) Save(webhook *model.Command) (*model.Command, *model.AppError) {
	return s.baseStore.Save(webhook)
}

func (s TransparentCommandStore) Get(id string) store.StoreChannel {
	return s.baseStore.Get(id)
}

func (s TransparentCommandStore) GetByTeam(teamId string) ([]*model.Command, *model.AppError) {
	return s.baseStore.GetByTeam(teamId)
}

func (s TransparentCommandStore) GetByTrigger(teamId string, trigger string) store.StoreChannel {
	return s.baseStore.GetByTrigger(teamId, trigger)
}

func (s TransparentCommandStore) Delete(commandId string, time int64) *model.AppError {
	return s.baseStore.Delete(commandId, time)
}

func (s TransparentCommandStore) PermanentDeleteByTeam(teamId string) *model.AppError {
	return s.baseStore.PermanentDeleteByTeam(teamId)
}

func (s TransparentCommandStore) PermanentDeleteByUser(userId string) *model.AppError {
	return s.baseStore.PermanentDeleteByUser(userId)
}

func (s TransparentCommandStore) Update(hook *model.Command) (*model.Command, *model.AppError) {
	return s.baseStore.Update(hook)
}

func (s TransparentCommandStore) AnalyticsCommandCount(teamId string) (int64, *model.AppError) {
	return s.baseStore.AnalyticsCommandCount(teamId)
}
