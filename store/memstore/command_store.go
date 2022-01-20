// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package memstore

import (
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/store"
)

type MemCommandStore struct {
	commands []*model.Command
}

func newMemCommandStore() store.CommandStore {
	return &MemCommandStore{}
}

func (s *MemCommandStore) Save(command *model.Command) (*model.Command, error) {
	if command.Id != "" {
		return nil, store.NewErrInvalidInput("Command", "CommandId", command.Id)
	}

	command.PreSave()
	if err := command.IsValid(); err != nil {
		return nil, err
	}

	s.commands = append(s.commands, command)

	return command, nil
}

func (s *MemCommandStore) Get(id string) (*model.Command, error) {
	for _, command := range s.commands {
		if command.Id == id && command.DeleteAt == 0 {
			return command, nil
		}
	}

	return nil, store.NewErrNotFound("Command", id)
}

func (s *MemCommandStore) GetByTeam(teamId string) ([]*model.Command, error) {
	result := []*model.Command{}
	for _, command := range s.commands {
		if command.TeamId == teamId && command.DeleteAt == 0 {
			result = append(result, command)
		}
	}

	return result, nil
}

func (s *MemCommandStore) GetByTrigger(teamId string, trigger string) (*model.Command, error) {
	for _, command := range s.commands {
		if command.TeamId == teamId && command.Trigger == trigger && command.DeleteAt == 0 {
			return command, nil
		}
	}

	return nil, store.NewErrNotFound("Command", trigger)
}

func (s *MemCommandStore) Delete(commandId string, time int64) error {
	for _, command := range s.commands {
		if command.Id == commandId {
			command.DeleteAt = time
			command.UpdateAt = time
		}
	}
	return nil
}

func (s *MemCommandStore) PermanentDeleteByTeam(teamId string) error {
	result := []*model.Command{}
	for _, command := range s.commands {
		if command.TeamId != teamId {
			result = append(result, command)
		}
	}
	s.commands = result
	return nil
}

func (s *MemCommandStore) PermanentDeleteByUser(userId string) error {
	result := []*model.Command{}
	for _, command := range s.commands {
		if command.CreatorId != userId {
			result = append(result, command)
		}
	}
	s.commands = result
	return nil
}

func (s *MemCommandStore) Update(cmd *model.Command) (*model.Command, error) {
	cmd.UpdateAt = model.GetMillis()

	if err := cmd.IsValid(); err != nil {
		return nil, err
	}

	for i, command := range s.commands {
		if command.Id == cmd.Id {
			s.commands[i] = cmd
		}
	}
	return cmd, nil
}

func (s *MemCommandStore) AnalyticsCommandCount(teamId string) (int64, error) {
	counter := int64(0)
	for _, command := range s.commands {
		if command.DeleteAt == 0 {
			if teamId == "" || command.TeamId == teamId {
				counter += 1
			}
		}
	}
	return counter, nil
}
