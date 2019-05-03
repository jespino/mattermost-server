package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentUserTermsOfServiceStore struct {
	baseStore store.UserTermsOfServiceStore
}

func (s TransparentUserTermsOfServiceStore) GetByUser(userId string) store.StoreChannel {
	return s.baseStore.GetByUser(userId)
}

func (s TransparentUserTermsOfServiceStore) Save(userTermsOfService *model.UserTermsOfService) store.StoreChannel {
	return s.baseStore.Save(userTermsOfService)
}

func (s TransparentUserTermsOfServiceStore) Delete(userId string, termsOfServiceId string) store.StoreChannel {
	return s.baseStore.Delete(userId, termsOfServiceId)
}
