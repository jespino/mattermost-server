package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentWebhookStore struct {
	baseStore store.WebhookStore
}

func (s TransparentWebhookStore) SaveIncoming(webhook *model.IncomingWebhook) (*model.IncomingWebhook, *model.AppError) {
	return s.baseStore.SaveIncoming(webhook)
}

func (s TransparentWebhookStore) GetIncoming(id string, allowFromCache bool) (*model.IncomingWebhook, *model.AppError) {
	return s.baseStore.GetIncoming(id, allowFromCache)
}

func (s TransparentWebhookStore) GetIncomingList(offset int, limit int) ([]*model.IncomingWebhook, *model.AppError) {
	return s.baseStore.GetIncomingList(offset, limit)
}

func (s TransparentWebhookStore) GetIncomingByTeam(teamId string, offset int, limit int) ([]*model.IncomingWebhook, *model.AppError) {
	return s.baseStore.GetIncomingByTeam(teamId, offset, limit)
}

func (s TransparentWebhookStore) UpdateIncoming(webhook *model.IncomingWebhook) (*model.IncomingWebhook, *model.AppError) {
	return s.baseStore.UpdateIncoming(webhook)
}

func (s TransparentWebhookStore) GetIncomingByChannel(channelId string) ([]*model.IncomingWebhook, *model.AppError) {
	return s.baseStore.GetIncomingByChannel(channelId)
}

func (s TransparentWebhookStore) DeleteIncoming(webhookId string, time int64) *model.AppError {
	return s.baseStore.DeleteIncoming(webhookId, time)
}

func (s TransparentWebhookStore) PermanentDeleteIncomingByChannel(channelId string) *model.AppError {
	return s.baseStore.PermanentDeleteIncomingByChannel(channelId)
}

func (s TransparentWebhookStore) PermanentDeleteIncomingByUser(userId string) *model.AppError {
	return s.baseStore.PermanentDeleteIncomingByUser(userId)
}

func (s TransparentWebhookStore) SaveOutgoing(webhook *model.OutgoingWebhook) (*model.OutgoingWebhook, *model.AppError) {
	return s.baseStore.SaveOutgoing(webhook)
}

func (s TransparentWebhookStore) GetOutgoing(id string) (*model.OutgoingWebhook, *model.AppError) {
	return s.baseStore.GetOutgoing(id)
}

func (s TransparentWebhookStore) GetOutgoingList(offset int, limit int) ([]*model.OutgoingWebhook, *model.AppError) {
	return s.baseStore.GetOutgoingList(offset, limit)
}

func (s TransparentWebhookStore) GetOutgoingByChannel(channelId string, offset int, limit int) store.StoreChannel {
	return s.baseStore.GetOutgoingByChannel(channelId, offset, limit)
}

func (s TransparentWebhookStore) GetOutgoingByTeam(teamId string, offset int, limit int) ([]*model.OutgoingWebhook, *model.AppError) {
	return s.baseStore.GetOutgoingByTeam(teamId, offset, limit)
}

func (s TransparentWebhookStore) DeleteOutgoing(webhookId string, time int64) *model.AppError {
	return s.baseStore.DeleteOutgoing(webhookId, time)
}

func (s TransparentWebhookStore) PermanentDeleteOutgoingByChannel(channelId string) *model.AppError {
	return s.baseStore.PermanentDeleteOutgoingByChannel(channelId)
}

func (s TransparentWebhookStore) PermanentDeleteOutgoingByUser(userId string) *model.AppError {
	return s.baseStore.PermanentDeleteOutgoingByUser(userId)
}

func (s TransparentWebhookStore) UpdateOutgoing(hook *model.OutgoingWebhook) (*model.OutgoingWebhook, *model.AppError) {
	return s.baseStore.UpdateOutgoing(hook)
}

func (s TransparentWebhookStore) AnalyticsIncomingCount(teamId string) (int64, *model.AppError) {
	return s.baseStore.AnalyticsIncomingCount(teamId)
}

func (s TransparentWebhookStore) AnalyticsOutgoingCount(teamId string) (int64, *model.AppError) {
	return s.baseStore.AnalyticsOutgoingCount(teamId)
}

func (s TransparentWebhookStore) InvalidateWebhookCache(webhook string) {
	s.baseStore.InvalidateWebhookCache(webhook)
}

func (s TransparentWebhookStore) ClearCaches() {
	s.baseStore.ClearCaches()
}
