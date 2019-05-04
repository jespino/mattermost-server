package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaCommandStore struct {
	baseStore store.CommandStore
	root      *KafkaStore
}

func (s KafkaCommandStore) Save(webhook *model.Command) (*model.Command, *model.AppError) {
	return s.baseStore.Save(webhook)
}

func (s KafkaCommandStore) Get(id string) store.StoreChannel {
	return s.baseStore.Get(id)
}

func (s KafkaCommandStore) GetByTeam(teamId string) ([]*model.Command, *model.AppError) {
	return s.baseStore.GetByTeam(teamId)
}

func (s KafkaCommandStore) GetByTrigger(teamId string, trigger string) store.StoreChannel {
	return s.baseStore.GetByTrigger(teamId, trigger)
}

func (s KafkaCommandStore) Delete(commandId string, time int64) *model.AppError {
	return s.baseStore.Delete(commandId, time)
}

func (s KafkaCommandStore) PermanentDeleteByTeam(teamId string) *model.AppError {
	return s.baseStore.PermanentDeleteByTeam(teamId)
}

func (s KafkaCommandStore) PermanentDeleteByUser(userId string) *model.AppError {
	return s.baseStore.PermanentDeleteByUser(userId)
}

func (s KafkaCommandStore) Update(hook *model.Command) (*model.Command, *model.AppError) {
	return s.baseStore.Update(hook)
}

func (s KafkaCommandStore) AnalyticsCommandCount(teamId string) (int64, *model.AppError) {
	return s.baseStore.AnalyticsCommandCount(teamId)
}
