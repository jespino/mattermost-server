package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentTermsOfServiceStore struct {
	baseStore store.TermsOfServiceStore
}

func (s TransparentTermsOfServiceStore) Save(termsOfService *model.TermsOfService) store.StoreChannel {
	return s.baseStore.Save(termsOfService)
}

func (s TransparentTermsOfServiceStore) GetLatest(allowFromCache bool) store.StoreChannel {
	return s.baseStore.GetLatest(allowFromCache)
}

func (s TransparentTermsOfServiceStore) Get(id string, allowFromCache bool) store.StoreChannel {
	return s.baseStore.Get(id, allowFromCache)
}
