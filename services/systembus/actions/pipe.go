package actions

import (
	"encoding/json"

	"github.com/mattermost/mattermost-server/v6/services/systembus"
)

const PipeID = "pipe"

func NewPipe(systemBus systembus.SystemBus) *systembus.ActionDefinition {
	handler := func(event *systembus.Event, config map[string]string) (*systembus.Event, error) {
		var actions []map[string]map[string]string
		err := json.Unmarshal([]byte(config["actions"]), &actions)
		if err != nil {
			return nil, err
		}

		actionDefinitions, err := systemBus.ListActions()
		if err != nil {
			return nil, err
		}

		actionsMap := map[string]*systembus.ActionDefinition{}
		for _, actionDefinition := range actionDefinitions {
			actionsMap[actionDefinition.ID] = actionDefinition
		}

		currentEvent := event
		for _, actionPair := range actions {
			for actionID, actionConfig := range actionPair {
				var err error
				actionDefinition := actionsMap[actionID]
				currentEvent, err = actionDefinition.Handler(currentEvent, actionConfig)
				if err != nil {
					return nil, err
				}
				if currentEvent == nil {
					return nil, nil
				}
			}
		}

		return nil, nil
	}

	return &systembus.ActionDefinition{
		ID:               PipeID,
		Name:             "Pipe",
		Description:      "Generate a pipe of actions to execute, sending the output of the action as the input of the next action",
		ConfigDefinition: map[string]string{"actions": "json"},
		Handler:          handler,
	}
}
