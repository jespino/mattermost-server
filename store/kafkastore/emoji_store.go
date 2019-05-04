package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaEmojiStore struct {
	baseStore store.EmojiStore
	root      *KafkaStore
}

func (s KafkaEmojiStore) Save(emoji *model.Emoji) store.StoreChannel {
	return s.baseStore.Save(emoji)
}

func (s KafkaEmojiStore) Get(id string, allowFromCache bool) store.StoreChannel {
	return s.baseStore.Get(id, allowFromCache)
}

func (s KafkaEmojiStore) GetByName(name string) store.StoreChannel {
	return s.baseStore.GetByName(name)
}

func (s KafkaEmojiStore) GetMultipleByName(names []string) store.StoreChannel {
	return s.baseStore.GetMultipleByName(names)
}

func (s KafkaEmojiStore) GetList(offset int, limit int, sort string) store.StoreChannel {
	return s.baseStore.GetList(offset, limit, sort)
}

func (s KafkaEmojiStore) Delete(id string, time int64) store.StoreChannel {
	return s.baseStore.Delete(id, time)
}

func (s KafkaEmojiStore) Search(name string, prefixOnly bool, limit int) store.StoreChannel {
	return s.baseStore.Search(name, prefixOnly, limit)
}
