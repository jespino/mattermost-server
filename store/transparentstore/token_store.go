package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentTokenStore struct {
	baseStore store.TokenStore
}

func (s TransparentTokenStore) Save(recovery *model.Token) store.StoreChannel {
	return s.baseStore.Save(recovery)
}

func (s TransparentTokenStore) Delete(token string) store.StoreChannel {
	return s.baseStore.Delete(token)
}

func (s TransparentTokenStore) GetByToken(token string) store.StoreChannel {
	return s.baseStore.GetByToken(token)
}

func (s TransparentTokenStore) Cleanup() {
	s.baseStore.Cleanup()
}

func (s TransparentTokenStore) RemoveAllTokensByType(tokenType string) store.StoreChannel {
	return s.baseStore.RemoveAllTokensByType(tokenType)
}
