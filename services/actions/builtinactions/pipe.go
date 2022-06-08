package builtinactions

import (
	"encoding/json"

	"github.com/mattermost/mattermost-server/v6/services/actions"
)

const PipeID = "pipe"

func NewPipe(actionsService *actions.Actions) *actions.ActionDefinition {
	handler := func(config map[string]string, data map[string]string) (map[string]string, error) {
		var actionsData []map[string]map[string]string
		err := json.Unmarshal([]byte(config["actions"]), &actionsData)
		if err != nil {
			return nil, err
		}

		actionDefinitions, err := actionsService.ListActions()
		if err != nil {
			return nil, err
		}

		actionsMap := map[string]*actions.ActionDefinition{}
		for _, actionDefinition := range actionDefinitions {
			actionsMap[actionDefinition.ID] = actionDefinition
		}

		currentData := data
		for _, actionPair := range actionsData {
			for actionID, actionConfig := range actionPair {
				var err error
				actionDefinition := actionsMap[actionID]
				currentData, err = actionDefinition.Handler(actionConfig, currentData)
				if err != nil {
					return nil, err
				}
				if currentData == nil {
					return nil, nil
				}
			}
		}

		return nil, nil
	}

	return &actions.ActionDefinition{
		ID:               PipeID,
		Name:             "Pipe",
		Description:      "Generate a pipe of actions to execute, sending the output of the action as the input of the next action",
		ConfigDefinition: map[string]string{"actions": "json"},
		Handler:          handler,
	}
}
