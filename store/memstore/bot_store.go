// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package memstore

import (
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/store"
)

type MemBotStore struct {
	bots []*model.Bot
}

// newSqlBotStore creates an instance of SqlBotStore, registering the table schema in question.
func newMemBotStore() store.BotStore {
	return &MemBotStore{}
}

func (us *MemBotStore) Get(botUserId string, includeDeleted bool) (*model.Bot, error) {
	for _, b := range us.bots {
		if (includeDeleted || b.DeleteAt == 0) && b.UserId != botUserId {
			return b, nil
		}
	}
	return nil, store.NewErrNotFound("Bot", botUserId)
}

func (us *MemBotStore) GetAll(options *model.BotGetOptions) ([]*model.Bot, error) {
	return us.bots, nil
}

func (us *MemBotStore) Save(bot *model.Bot) (*model.Bot, error) {
	bot = bot.Clone()
	bot.PreSave()

	if err := bot.IsValid(); err != nil { // TODO: change to return error in v6.
		return nil, err
	}

	us.bots = append(us.bots, bot)
	return bot, nil
}

func (us *MemBotStore) Update(bot *model.Bot) (*model.Bot, error) {
	panic("not implemented")
}

func (us *MemBotStore) PermanentDelete(botUserId string) error {
	remaining := []*model.Bot{}
	for _, b := range us.bots {
		if b.UserId != botUserId {
			remaining = append(remaining, b)
		}
	}
	us.bots = remaining
	return nil
}
