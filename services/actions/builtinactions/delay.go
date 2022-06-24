package builtinactions

import (
	"strconv"
	"time"

	"github.com/mattermost/mattermost-server/v6/services/actions"
)

const DelayID = "delay"

func NewDelay() *actions.ActionDefinition {
	handler := func(data map[string]string) (map[string]string, error) {
		delay := data["delay"]
		delayInt, err := strconv.Atoi(delay)
		if err != nil {
			return nil, err
		}

		time.Sleep(time.Duration(delayInt) * time.Second)

		return data, nil
	}

	return &actions.ActionDefinition{
		ID:               DelayID,
		Name:             "Delay",
		Description:      "Wait for some seconds and then kep running the task",
		ConfigDefinition: map[string]string{"delay": "number"},
		Handler:          handler,
	}
}
