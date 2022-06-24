package builtinactions

import (
	"fmt"

	"github.com/mattermost/mattermost-server/v6/services/actions"
	lua "github.com/yuin/gopher-lua"
)

const LuaID = "lua"

func NewLua() *actions.ActionDefinition {
	handler := func(data map[string]string) (map[string]string, error) {
		l := lua.NewState()
		defer l.Close()
		for k, v := range data {
			l.SetGlobal(fmt.Sprintf("data_%s", k), lua.LString(v))
		}

		if err := l.DoString(data["code"]); err != nil {
			return nil, err
		}
		return nil, nil
	}

	return &actions.ActionDefinition{
		ID:               LuaID,
		Name:             "Lua",
		Description:      "Lua custom action.",
		ConfigDefinition: map[string]string{"code": "longstring"},
		Handler:          handler,
	}
}
