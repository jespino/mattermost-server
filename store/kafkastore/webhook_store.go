package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaWebhookStore struct {
	baseStore store.WebhookStore
	root      *KafkaStore
}

func (s KafkaWebhookStore) SaveIncoming(webhook *model.IncomingWebhook) (*model.IncomingWebhook, *model.AppError) {
	return s.baseStore.SaveIncoming(webhook)
}

func (s KafkaWebhookStore) GetIncoming(id string, allowFromCache bool) (*model.IncomingWebhook, *model.AppError) {
	return s.baseStore.GetIncoming(id, allowFromCache)
}

func (s KafkaWebhookStore) GetIncomingList(offset int, limit int) ([]*model.IncomingWebhook, *model.AppError) {
	return s.baseStore.GetIncomingList(offset, limit)
}

func (s KafkaWebhookStore) GetIncomingByTeam(teamId string, offset int, limit int) ([]*model.IncomingWebhook, *model.AppError) {
	return s.baseStore.GetIncomingByTeam(teamId, offset, limit)
}

func (s KafkaWebhookStore) UpdateIncoming(webhook *model.IncomingWebhook) (*model.IncomingWebhook, *model.AppError) {
	return s.baseStore.UpdateIncoming(webhook)
}

func (s KafkaWebhookStore) GetIncomingByChannel(channelId string) ([]*model.IncomingWebhook, *model.AppError) {
	return s.baseStore.GetIncomingByChannel(channelId)
}

func (s KafkaWebhookStore) DeleteIncoming(webhookId string, time int64) *model.AppError {
	return s.baseStore.DeleteIncoming(webhookId, time)
}

func (s KafkaWebhookStore) PermanentDeleteIncomingByChannel(channelId string) *model.AppError {
	return s.baseStore.PermanentDeleteIncomingByChannel(channelId)
}

func (s KafkaWebhookStore) PermanentDeleteIncomingByUser(userId string) *model.AppError {
	return s.baseStore.PermanentDeleteIncomingByUser(userId)
}

func (s KafkaWebhookStore) SaveOutgoing(webhook *model.OutgoingWebhook) (*model.OutgoingWebhook, *model.AppError) {
	return s.baseStore.SaveOutgoing(webhook)
}

func (s KafkaWebhookStore) GetOutgoing(id string) (*model.OutgoingWebhook, *model.AppError) {
	return s.baseStore.GetOutgoing(id)
}

func (s KafkaWebhookStore) GetOutgoingList(offset int, limit int) ([]*model.OutgoingWebhook, *model.AppError) {
	return s.baseStore.GetOutgoingList(offset, limit)
}

func (s KafkaWebhookStore) GetOutgoingByChannel(channelId string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetOutgoingByChannel(channelId, offset, limit)
}

func (s KafkaWebhookStore) GetOutgoingByTeam(teamId string, offset int, limit int) ([]*model.OutgoingWebhook, *model.AppError) {
	return s.baseStore.GetOutgoingByTeam(teamId, offset, limit)
}

func (s KafkaWebhookStore) DeleteOutgoing(webhookId string, time int64) *model.AppError {
	return s.baseStore.DeleteOutgoing(webhookId, time)
}

func (s KafkaWebhookStore) PermanentDeleteOutgoingByChannel(channelId string) *model.AppError {
	return s.baseStore.PermanentDeleteOutgoingByChannel(channelId)
}

func (s KafkaWebhookStore) PermanentDeleteOutgoingByUser(userId string) *model.AppError {
	return s.baseStore.PermanentDeleteOutgoingByUser(userId)
}

func (s KafkaWebhookStore) UpdateOutgoing(hook *model.OutgoingWebhook) (*model.OutgoingWebhook, *model.AppError) {
	return s.baseStore.UpdateOutgoing(hook)
}

func (s KafkaWebhookStore) AnalyticsIncomingCount(teamId string) (int64, *model.AppError) {
	return s.baseStore.AnalyticsIncomingCount(teamId)
}

func (s KafkaWebhookStore) AnalyticsOutgoingCount(teamId string) (int64, *model.AppError) {
	return s.baseStore.AnalyticsOutgoingCount(teamId)
}

func (s KafkaWebhookStore) InvalidateWebhookCache(webhook string) {
	s.baseStore.InvalidateWebhookCache(webhook)
}

func (s KafkaWebhookStore) ClearCaches() {
	s.baseStore.ClearCaches()
}
