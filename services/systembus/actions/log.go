package actions

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-server/v6/services/systembus"
	"github.com/mattermost/mattermost-server/v6/shared/mlog"
)

const LogID = "print"

func NewLog(logger *mlog.Logger) *systembus.ActionDefinition {
	handler := func(event *systembus.Event, config map[string]string) (*systembus.Event, error) {
		logLevel := config["level"]
		logFunc := logger.Debug
		switch strings.ToLower(logLevel) {
		case "trace":
			logFunc = logger.Trace
		case "info":
			logFunc = logger.Info
		case "warn":
			logFunc = logger.Warn
		case "error":
			logFunc = logger.Error
		case "critical":
			logFunc = logger.Critical
		}

		if config["template"] == "" {
			logFunc(fmt.Sprintf("Event: %s", event))
			return nil, nil
		}

		message, err := applyTemplate(config["template"], event.Data)
		if err != nil {
			return nil, err
		}

		logFunc(message)
		return nil, nil
	}

	return &systembus.ActionDefinition{
		ID:               LogID,
		Name:             "Log",
		Description:      "Logs the data passed as parameter.",
		ConfigDefinition: map[string]string{"template": "string", "level": "string"},
		Handler:          handler,
	}
}
