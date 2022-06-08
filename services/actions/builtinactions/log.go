package builtinactions

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-server/v6/services/actions"
	"github.com/mattermost/mattermost-server/v6/shared/mlog"
)

const LogID = "print"

func NewLog(logger *mlog.Logger) *actions.ActionDefinition {
	handler := func(config map[string]string, data map[string]string) (map[string]string, error) {
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
			logFunc(fmt.Sprintf("Data: %s", data))
			return nil, nil
		}

		message := config["template"]
		logFunc(message)
		return nil, nil
	}

	return &actions.ActionDefinition{
		ID:               LogID,
		Name:             "Log",
		Description:      "Logs the data passed as parameter.",
		ConfigDefinition: map[string]string{"template": "string", "level": "string"},
		Handler:          handler,
	}
}
