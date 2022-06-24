package builtinactions

import (
	"github.com/mattermost/mattermost-server/v6/services/actions"
	"github.com/robertkrimen/otto"
)

const JavascriptID = "javascript"

func NewJavascript() *actions.ActionDefinition {
	handler := func(data map[string]string) (map[string]string, error) {
		vm := otto.New()
		err := vm.Set("data", data)
		if err != nil {
			return nil, err
		}

		if _, err := vm.Run(data["code"]); err != nil {
			return nil, err
		}

		return nil, nil
	}

	return &actions.ActionDefinition{
		ID:               JavascriptID,
		Name:             "Javascript",
		Description:      "Javascript custom action.",
		ConfigDefinition: map[string]string{"code": "longstring"},
		Handler:          handler,
	}
}
