// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package memstore

import (
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/store"
)

// SqlLicenseStore encapsulates the database writes and reads for
// model.LicenseRecord objects.
type MemLicenseStore struct {
	licenses []*model.LicenseRecord
}

func newMemLicenseStore() store.LicenseStore {
	return &MemLicenseStore{}
}

func (ls *MemLicenseStore) Save(license *model.LicenseRecord) (*model.LicenseRecord, error) {
	license.PreSave()
	if err := license.IsValid(); err != nil {
		return nil, err
	}

	for _, l := range ls.licenses {
		if l.Id == license.Id {
			*l = *license
			return license, nil
		}
	}
	ls.licenses = append(ls.licenses, license)
	return license, nil
}

func (ls *MemLicenseStore) Get(id string) (*model.LicenseRecord, error) {
	for _, l := range ls.licenses {
		if l.Id == id {
			return l, nil
		}
	}
	return nil, store.NewErrNotFound("License", id)
}

func (ls *MemLicenseStore) GetAll() ([]*model.LicenseRecord, error) {
	return ls.licenses, nil
}
