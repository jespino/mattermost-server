package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaPreferenceStore struct {
	baseStore store.PreferenceStore
	root      *KafkaStore
}

func (s KafkaPreferenceStore) Save(preferences *model.Preferences) store.StoreChannel {
	return s.baseStore.Save(preferences)
}

func (s KafkaPreferenceStore) Get(userId string, category string, name string) store.StoreChannel {
	return s.baseStore.Get(userId, category, name)
}

func (s KafkaPreferenceStore) GetCategory(userId string, category string) store.StoreChannel {
	return s.baseStore.GetCategory(userId, category)
}

func (s KafkaPreferenceStore) GetAll(userId string) store.StoreChannel {
	return s.baseStore.GetAll(userId)
}

func (s KafkaPreferenceStore) Delete(userId string, category string, name string) store.StoreChannel {
	return s.baseStore.Delete(userId, category, name)
}

func (s KafkaPreferenceStore) DeleteCategory(userId string, category string) store.StoreChannel {
	return s.baseStore.DeleteCategory(userId, category)
}

func (s KafkaPreferenceStore) DeleteCategoryAndName(category string, name string) store.StoreChannel {
	return s.baseStore.DeleteCategoryAndName(category, name)
}

func (s KafkaPreferenceStore) PermanentDeleteByUser(userId string) store.StoreChannel {
	return s.baseStore.PermanentDeleteByUser(userId)
}

func (s KafkaPreferenceStore) IsFeatureEnabled(feature string, userId string) store.StoreChannel {
	return s.baseStore.IsFeatureEnabled(feature, userId)
}

func (s KafkaPreferenceStore) CleanupFlagsBatch(limit int64) store.StoreChannel {
	return s.baseStore.CleanupFlagsBatch(limit)
}
