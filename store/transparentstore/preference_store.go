package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentPreferenceStore struct {
	baseStore store.PreferenceStore
}

func (s TransparentPreferenceStore) Save(preferences *model.Preferences) store.StoreChannel {
	return s.baseStore.Save(preferences)
}

func (s TransparentPreferenceStore) Get(userId string, category string, name string) store.StoreChannel {
	return s.baseStore.Get(userId, category, name)
}

func (s TransparentPreferenceStore) GetCategory(userId string, category string) store.StoreChannel {
	return s.baseStore.GetCategory(userId, category)
}

func (s TransparentPreferenceStore) GetAll(userId string) store.StoreChannel {
	return s.baseStore.GetAll(userId)
}

func (s TransparentPreferenceStore) Delete(userId string, category string, name string) store.StoreChannel {
	return s.baseStore.Delete(userId, category, name)
}

func (s TransparentPreferenceStore) DeleteCategory(userId string, category string) store.StoreChannel {
	return s.baseStore.DeleteCategory(userId, category)
}

func (s TransparentPreferenceStore) DeleteCategoryAndName(category string, name string) store.StoreChannel {
	return s.baseStore.DeleteCategoryAndName(category, name)
}

func (s TransparentPreferenceStore) PermanentDeleteByUser(userId string) store.StoreChannel {
	return s.baseStore.PermanentDeleteByUser(userId)
}

func (s TransparentPreferenceStore) IsFeatureEnabled(feature string, userId string) store.StoreChannel {
	return s.baseStore.IsFeatureEnabled(feature, userId)
}

func (s TransparentPreferenceStore) CleanupFlagsBatch(limit int64) store.StoreChannel {
	return s.baseStore.CleanupFlagsBatch(limit)
}
