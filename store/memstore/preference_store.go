// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package memstore

import (
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/store"
)

type MemPreferenceStore struct {
	preferences []*model.Preference
}

func newMemPreferenceStore() store.PreferenceStore {
	return &MemPreferenceStore{}
}

func (s *MemPreferenceStore) Save(preferences model.Preferences) error {
	for _, preference := range preferences {
		preference.PreUpdate()
		if err := preference.IsValid(); err != nil {
			return err
		}
		updated := false
		for _, p := range s.preferences {
			if p.UserId == preference.UserId && p.Category == preference.Category && p.Name == preference.Name {
				p.Value = preference.Value
				updated = true
				break
			}
		}
		if !updated {
			s.preferences = append(s.preferences, &preference)
		}
	}
	return nil
}

func (s *MemPreferenceStore) Get(userId string, category string, name string) (*model.Preference, error) {
	panic("not implemented")
}

func (s *MemPreferenceStore) GetCategory(userId string, category string) (model.Preferences, error) {
	panic("not implemented")
}

func (s *MemPreferenceStore) GetAll(userId string) (model.Preferences, error) {
	panic("not implemented")
}

func (s *MemPreferenceStore) PermanentDeleteByUser(userId string) error {
	panic("not implemented")
}

func (s *MemPreferenceStore) Delete(userId, category, name string) error {
	panic("not implemented")
}

func (s *MemPreferenceStore) DeleteCategory(userId string, category string) error {
	panic("not implemented")
}

func (s *MemPreferenceStore) DeleteCategoryAndName(category string, name string) error {
	panic("not implemented")
}

func (s *MemPreferenceStore) DeleteOrphanedRows(limit int) (deleted int64, err error) {
	panic("not implemented")
}

func (s *MemPreferenceStore) CleanupFlagsBatch(limit int64) (int64, error) {
	panic("not implemented")
}
