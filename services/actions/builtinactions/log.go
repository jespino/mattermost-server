package builtinactions

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-server/v6/services/actions"
	"github.com/mattermost/mattermost-server/v6/shared/mlog"
)

const LogID = "log"

func NewLog(logger *mlog.Logger) *actions.ActionDefinition {
	handler := func(data map[string]string) (map[string]string, error) {
		logLevel := data["level"]
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

		if data["template"] == "" {
			logFunc(fmt.Sprintf("Data: %s", data))
			return data, nil
		}

		message := data["template"]
		logFunc(message)
		return data, nil
	}

	return &actions.ActionDefinition{
		ID:               LogID,
		Name:             "Log",
		Description:      "Logs the data passed as parameter.",
		ConfigDefinition: map[string]string{"template": "string", "level": "string"},
		Handler:          handler,
	}
}
