package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentLicenseStore struct {
	baseStore store.LicenseStore
}

func (s TransparentLicenseStore) Save(license *model.LicenseRecord) (*model.LicenseRecord, *model.AppError) {
	return s.baseStore.Save(license)
}

func (s TransparentLicenseStore) Get(id string) (*model.LicenseRecord, *model.AppError) {
	return s.baseStore.Get(id)
}
