// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package memstore

import (
	"context"
	"sort"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/store"
)

type MemEmojiStore struct {
	emojis []*model.Emoji
}

func newMemEmojiStore() store.EmojiStore {
	return &MemEmojiStore{}
}

func (es *MemEmojiStore) Save(emoji *model.Emoji) (*model.Emoji, error) {
	emoji.PreSave()
	if err := emoji.IsValid(); err != nil {
		return nil, err
	}

	es.emojis = append(es.emojis, emoji)

	return emoji, nil
}

func (es *MemEmojiStore) Get(ctx context.Context, id string, allowFromCache bool) (*model.Emoji, error) {
	for _, e := range es.emojis {
		if e.Id == id {
			return e, nil
		}
	}
	return nil, store.NewErrNotFound("Emoji", id)
}

func (es *MemEmojiStore) GetByName(ctx context.Context, name string, allowFromCache bool) (*model.Emoji, error) {
	for _, e := range es.emojis {
		if e.Name == name {
			return e, nil
		}
	}
	return nil, store.NewErrNotFound("Emoji", name)
}

func (es *MemEmojiStore) GetMultipleByName(names []string) ([]*model.Emoji, error) {
	result := []*model.Emoji{}

	for _, e := range es.emojis {
		if e.DeleteAt == 0 {
			for _, n := range names {
				if e.Name == n {
					result = append(result, e)
				}
			}
		}
	}
	return result, nil
}

func (es *MemEmojiStore) GetList(offset, limit int, sortBy string) ([]*model.Emoji, error) {
	sorted := []*model.Emoji{}
	for _, e := range es.emojis {
		if e.DeleteAt == 0 {
			sorted = append(sorted, e)
		}
	}
	if sortBy == model.EmojiSortByName {
		sort.Slice(sorted, func(i, j int) bool { return sorted[i].Name > sorted[j].Name })
	}
	counter := 0
	result := []*model.Emoji{}
	for _, e := range sorted {
		if counter >= offset && counter < offset+limit {
			counter++
			result = append(result, e)
		}
	}
	return result, nil
}

func (es *MemEmojiStore) Delete(emoji *model.Emoji, time int64) error {
	for _, e := range es.emojis {
		if e.Id == emoji.Id {
			now := model.GetMillis()
			e.DeleteAt = now
			e.UpdateAt = now
			return nil
		}
	}
	return store.NewErrNotFound("Emoji", emoji.Id)
}

func (es *MemEmojiStore) Search(name string, prefixOnly bool, limit int) ([]*model.Emoji, error) {
	panic("not implemented")
}
