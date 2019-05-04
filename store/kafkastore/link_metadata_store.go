package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaLinkMetadataStore struct {
	baseStore store.LinkMetadataStore
	root      *KafkaStore
}

func (s KafkaLinkMetadataStore) Save(linkMetadata *model.LinkMetadata) store.StoreChannel {
	return s.baseStore.Save(linkMetadata)
}

func (s KafkaLinkMetadataStore) Get(url string, timestamp int64) store.StoreChannel {
	return s.baseStore.Get(url, timestamp)
}
