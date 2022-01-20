// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package memstore

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/store"
	"github.com/mattermost/mattermost-server/v6/utils"
)

type MemSystemStore struct {
	values map[string]string
}

func newMemSystemStore() store.SystemStore {
	return &MemSystemStore{
		values: make(map[string]string),
	}
}

func (s *MemSystemStore) Save(system *model.System) error {
	if _, ok := s.values[system.Name]; ok {
		return errors.New("Already existing key")
	}

	s.values[system.Name] = system.Value
	return nil
}

func (s *MemSystemStore) SaveOrUpdate(system *model.System) error {
	s.values[system.Name] = system.Value
	return nil
}

func (s *MemSystemStore) SaveOrUpdateWithWarnMetricHandling(system *model.System) error {
	if err := s.SaveOrUpdate(system); err != nil {
		return err
	}

	if strings.HasPrefix(system.Name, model.WarnMetricStatusStorePrefix) &&
		(system.Value == model.WarnMetricStatusRunonce || system.Value == model.WarnMetricStatusLimitReached) {
		if err := s.SaveOrUpdate(&model.System{
			Name:  model.SystemWarnMetricLastRunTimestampKey,
			Value: strconv.FormatInt(utils.MillisFromTime(time.Now()), 10),
		}); err != nil {
			return errors.Wrapf(err, "failed to save system property with name=%s", model.SystemWarnMetricLastRunTimestampKey)
		}
	}

	return nil
}

func (s *MemSystemStore) Update(system *model.System) error {
	s.values[system.Name] = system.Value
	return nil
}

func (s *MemSystemStore) Get() (model.StringMap, error) {
	return s.values, nil
}

func (s *MemSystemStore) GetByName(name string) (*model.System, error) {
	value, ok := s.values[name]
	if !ok {
		return nil, store.NewErrNotFound("System", fmt.Sprintf("name=%s", name))
	}
	return &model.System{
		Name:  name,
		Value: value,
	}, nil
}

func (s *MemSystemStore) PermanentDeleteByName(name string) (*model.System, error) {
	delete(s.values, name)
	return nil, nil
}

// InsertIfExists inserts a given system value if it does not already exist. If a value
// already exists, it returns the old one, else returns the new one.
func (s *MemSystemStore) InsertIfExists(system *model.System) (*model.System, error) {
	value, ok := s.values[system.Name]
	if ok {
		return &model.System{
			Name:  system.Name,
			Value: value,
		}, nil
	}

	s.values[system.Name] = system.Value

	return system, nil
}
