package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentEmojiStore struct {
	baseStore store.EmojiStore
}

func (s TransparentEmojiStore) Save(emoji *model.Emoji) store.StoreChannel {
	return s.baseStore.Save(emoji)
}

func (s TransparentEmojiStore) Get(id string, allowFromCache bool) store.StoreChannel {
	return s.baseStore.Get(id, allowFromCache)
}

func (s TransparentEmojiStore) GetByName(name string) store.StoreChannel {
	return s.baseStore.GetByName(name)
}

func (s TransparentEmojiStore) GetMultipleByName(names []string) store.StoreChannel {
	return s.baseStore.GetMultipleByName(names)
}

func (s TransparentEmojiStore) GetList(offset int, limit int, sort string) store.StoreChannel {
	return s.baseStore.GetList(offset, limit, sort)
}

func (s TransparentEmojiStore) Delete(id string, time int64) store.StoreChannel {
	return s.baseStore.Delete(id, time)
}

func (s TransparentEmojiStore) Search(name string, prefixOnly bool, limit int) store.StoreChannel {
	return s.baseStore.Search(name, prefixOnly, limit)
}
