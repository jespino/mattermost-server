// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package sqlstore

import (
	"sync"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/store"
)

type MemWebhookStore struct {
	incomingWebhooks []*model.IncomingWebhook
	outgoingWebhooks []*model.OutgoingWebhook
	iMutex           sync.RWMutex
	oMutex           sync.RWMutex
}

func (s *MemWebhookStore) ClearCaches() {
}

func newMemWebhookStore() store.WebhookStore {
	s := &MemWebhookStore{
		incomingWebhooks: []*model.IncomingWebhook{},
		outgoingWebhooks: []*model.OutgoingWebhook{},
	}
	return s
}

func (s *MemWebhookStore) InvalidateWebhookCache(webhookId string) {
}

func (s *MemWebhookStore) SaveIncoming(webhook *model.IncomingWebhook) (*model.IncomingWebhook, error) {
	s.iMutex.Lock()
	defer s.iMutex.Unlock()

	if webhook.Id != "" {
		return nil, store.NewErrInvalidInput("IncomingWebhook", "id", webhook.Id)
	}

	webhook.PreSave()
	if err := webhook.IsValid(); err != nil {
		return nil, err
	}

	s.incomingWebhooks = append(s.incomingWebhooks, webhook)

	return webhook, nil
}

func (s *MemWebhookStore) UpdateIncoming(hook *model.IncomingWebhook) (*model.IncomingWebhook, error) {
	s.iMutex.Lock()
	defer s.iMutex.Unlock()
	hook.UpdateAt = model.GetMillis()

	for i, wh := range s.incomingWebhooks {
		if wh.Id == hook.Id {
			s.incomingWebhooks[i] = hook
		}
	}

	return hook, nil
}

func (s *MemWebhookStore) GetIncoming(id string, allowFromCache bool) (*model.IncomingWebhook, error) {
	s.iMutex.RLock()
	defer s.iMutex.RUnlock()
	for _, wh := range s.incomingWebhooks {
		if wh.DeleteAt == 0 && wh.Id == id {
			return wh, nil
		}
	}

	return nil, store.NewErrNotFound("IncomingWebhook", id)
}

func (s *MemWebhookStore) DeleteIncoming(webhookId string, time int64) error {
	s.iMutex.Lock()
	defer s.iMutex.Unlock()
	for _, wh := range s.incomingWebhooks {
		if wh.Id == webhookId {
			wh.DeleteAt = time
			wh.UpdateAt = time
		}
	}

	return nil
}

func (s *MemWebhookStore) PermanentDeleteIncomingByUser(userId string) error {
	s.iMutex.Lock()
	defer s.iMutex.Unlock()
	result := []*model.IncomingWebhook{}
	for _, wh := range s.incomingWebhooks {
		if wh.UserId != userId {
			result = append(result, wh)
		}
	}
	s.incomingWebhooks = result

	return nil
}

func (s *MemWebhookStore) PermanentDeleteIncomingByChannel(channelId string) error {
	s.iMutex.Lock()
	defer s.iMutex.Unlock()
	result := []*model.IncomingWebhook{}
	for _, wh := range s.incomingWebhooks {
		if wh.ChannelId != channelId {
			result = append(result, wh)
		}
	}
	s.incomingWebhooks = result

	return nil
}

func (s *MemWebhookStore) GetIncomingList(offset, limit int) ([]*model.IncomingWebhook, error) {
	return s.GetIncomingListByUser("", offset, limit)
}

func (s *MemWebhookStore) GetIncomingListByUser(userId string, offset, limit int) ([]*model.IncomingWebhook, error) {
	s.iMutex.RLock()
	defer s.iMutex.RUnlock()
	webhooks := []*model.IncomingWebhook{}
	counter := 0
	for _, wh := range s.incomingWebhooks {
		if wh.DeleteAt == 0 && (userId == "" || wh.UserId == userId) {
			counter++
			if counter > offset {
				webhooks = append(webhooks, wh)
			}

			if counter >= offset+limit {
				return webhooks, nil
			}
		}
	}
	return webhooks, nil
}

func (s *MemWebhookStore) GetIncomingByTeamByUser(teamId string, userId string, offset, limit int) ([]*model.IncomingWebhook, error) {
	s.iMutex.RLock()
	defer s.iMutex.RUnlock()
	webhooks := []*model.IncomingWebhook{}
	counter := 0
	for _, wh := range s.incomingWebhooks {
		if wh.DeleteAt == 0 && wh.TeamId == teamId && (userId == "" || wh.UserId == userId) {
			counter++
			if counter > offset {
				webhooks = append(webhooks, wh)
			}

			if counter >= offset+limit {
				return webhooks, nil
			}
		}
	}
	return webhooks, nil
}

func (s *MemWebhookStore) GetIncomingByTeam(teamId string, offset, limit int) ([]*model.IncomingWebhook, error) {
	return s.GetIncomingByTeamByUser(teamId, "", offset, limit)
}

func (s *MemWebhookStore) GetIncomingByChannel(channelId string) ([]*model.IncomingWebhook, error) {
	s.iMutex.RLock()
	defer s.iMutex.RUnlock()
	webhooks := []*model.IncomingWebhook{}
	for _, wh := range s.incomingWebhooks {
		if wh.DeleteAt == 0 && wh.ChannelId == channelId {
			webhooks = append(webhooks, wh)
			return webhooks, nil
		}
	}
	return webhooks, nil
}

func (s *MemWebhookStore) SaveOutgoing(webhook *model.OutgoingWebhook) (*model.OutgoingWebhook, error) {
	s.oMutex.Lock()
	defer s.oMutex.Unlock()
	if webhook.Id != "" {
		return nil, store.NewErrInvalidInput("IncomingWebhook", "id", webhook.Id)
	}

	webhook.PreSave()
	if err := webhook.IsValid(); err != nil {
		return nil, err
	}

	s.outgoingWebhooks = append(s.outgoingWebhooks, webhook)

	return webhook, nil
}

func (s *MemWebhookStore) GetOutgoing(id string) (*model.OutgoingWebhook, error) {
	s.oMutex.RLock()
	defer s.oMutex.RUnlock()
	for _, wh := range s.outgoingWebhooks {
		if wh.DeleteAt == 0 && wh.Id == id {
			return wh, nil
		}
	}

	return nil, store.NewErrNotFound("IncomingWebhook", id)
}

func (s *MemWebhookStore) GetOutgoingListByUser(userId string, offset, limit int) ([]*model.OutgoingWebhook, error) {
	s.oMutex.RLock()
	defer s.oMutex.RUnlock()
	webhooks := []*model.OutgoingWebhook{}
	counter := 0
	for _, wh := range s.outgoingWebhooks {
		if wh.DeleteAt == 0 && (userId == "" || wh.CreatorId == userId) {
			counter++
			if counter > offset {
				webhooks = append(webhooks, wh)
			}

			if counter >= offset+limit {
				return webhooks, nil
			}
		}
	}
	return webhooks, nil
}

func (s *MemWebhookStore) GetOutgoingList(offset, limit int) ([]*model.OutgoingWebhook, error) {
	s.oMutex.RLock()
	defer s.oMutex.RUnlock()
	return s.GetOutgoingListByUser("", offset, limit)

}

func (s *MemWebhookStore) GetOutgoingByChannelByUser(channelId string, userId string, offset, limit int) ([]*model.OutgoingWebhook, error) {
	s.oMutex.RLock()
	defer s.oMutex.RUnlock()
	webhooks := []*model.OutgoingWebhook{}
	counter := 0
	for _, wh := range s.outgoingWebhooks {
		if wh.DeleteAt == 0 && wh.ChannelId == channelId && (userId == "" || wh.CreatorId == userId) {
			counter++
			if counter > offset {
				webhooks = append(webhooks, wh)
			}

			if counter >= offset+limit {
				return webhooks, nil
			}
		}
	}
	return webhooks, nil
}

func (s *MemWebhookStore) GetOutgoingByChannel(channelId string, offset, limit int) ([]*model.OutgoingWebhook, error) {
	return s.GetOutgoingByChannelByUser(channelId, "", offset, limit)
}

func (s *MemWebhookStore) GetOutgoingByTeamByUser(teamId string, userId string, offset, limit int) ([]*model.OutgoingWebhook, error) {
	s.oMutex.RLock()
	defer s.oMutex.RUnlock()
	webhooks := []*model.OutgoingWebhook{}
	counter := 0
	for _, wh := range s.outgoingWebhooks {
		if wh.DeleteAt == 0 && wh.TeamId == teamId && (userId == "" || wh.CreatorId == userId) {
			counter++
			if counter > offset {
				webhooks = append(webhooks, wh)
			}

			if counter >= offset+limit {
				return webhooks, nil
			}
		}
	}
	return webhooks, nil
}

func (s *MemWebhookStore) GetOutgoingByTeam(teamId string, offset, limit int) ([]*model.OutgoingWebhook, error) {
	return s.GetOutgoingByTeamByUser(teamId, "", offset, limit)
}

func (s *MemWebhookStore) DeleteOutgoing(webhookId string, time int64) error {
	s.oMutex.Lock()
	defer s.oMutex.Unlock()
	for _, wh := range s.outgoingWebhooks {
		if wh.Id == webhookId {
			wh.DeleteAt = time
			wh.UpdateAt = time
		}
	}

	return nil
}

func (s *MemWebhookStore) PermanentDeleteOutgoingByUser(userId string) error {
	s.oMutex.Lock()
	defer s.oMutex.Unlock()
	result := []*model.OutgoingWebhook{}
	for _, wh := range s.outgoingWebhooks {
		if wh.CreatorId != userId {
			result = append(result, wh)
		}
	}
	s.outgoingWebhooks = result

	return nil
}

func (s *MemWebhookStore) PermanentDeleteOutgoingByChannel(channelId string) error {
	s.oMutex.Lock()
	defer s.oMutex.Unlock()
	result := []*model.OutgoingWebhook{}
	for _, wh := range s.outgoingWebhooks {
		if wh.ChannelId != channelId {
			result = append(result, wh)
		}
	}
	s.outgoingWebhooks = result

	return nil
}

func (s *MemWebhookStore) UpdateOutgoing(hook *model.OutgoingWebhook) (*model.OutgoingWebhook, error) {
	s.oMutex.Lock()
	defer s.oMutex.Unlock()
	hook.UpdateAt = model.GetMillis()

	for i, wh := range s.outgoingWebhooks {
		if wh.Id == hook.Id {
			s.outgoingWebhooks[i] = hook
		}
	}

	return hook, nil
}

func (s *MemWebhookStore) AnalyticsIncomingCount(teamId string) (int64, error) {
	s.iMutex.RLock()
	defer s.iMutex.RUnlock()
	count := int64(0)
	for _, wh := range s.incomingWebhooks {
		if wh.DeleteAt == 0 {
			count++
		}
	}
	return count, nil
}

func (s *MemWebhookStore) AnalyticsOutgoingCount(teamId string) (int64, error) {
	s.oMutex.RLock()
	defer s.oMutex.RUnlock()
	count := int64(0)
	for _, wh := range s.outgoingWebhooks {
		if wh.DeleteAt == 0 {
			count++
		}
	}
	return count, nil
}
