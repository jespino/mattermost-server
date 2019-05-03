package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentLinkMetadataStore struct {
	baseStore store.LinkMetadataStore
}

func (s TransparentLinkMetadataStore) Save(linkMetadata *model.LinkMetadata) store.StoreChannel {
	return s.baseStore.Save(linkMetadata)
}

func (s TransparentLinkMetadataStore) Get(url string, timestamp int64) store.StoreChannel {
	return s.baseStore.Get(url, timestamp)
}
