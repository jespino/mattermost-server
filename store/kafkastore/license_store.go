package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaLicenseStore struct {
	baseStore store.LicenseStore
	root      *KafkaStore
}

func (s KafkaLicenseStore) Save(license *model.LicenseRecord) (*model.LicenseRecord, *model.AppError) {
	return s.baseStore.Save(license)
}

func (s KafkaLicenseStore) Get(id string) (*model.LicenseRecord, *model.AppError) {
	return s.baseStore.Get(id)
}
