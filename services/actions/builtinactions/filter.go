package builtinactions

import (
	"strings"

	"github.com/mattermost/mattermost-server/v6/services/actions"
)

const FilterID = "filter"

func NewFilter() *actions.ActionDefinition {
	handler := func(data map[string]string) (map[string]string, error) {
		template1 := data["template1"]
		template2 := data["template2"]

		switch data["comparison"] {
		case "":
			if template1 == template2 {
				return data, nil
			}
		case "eq":
			if template1 == template2 {
				return data, nil
			}
		case "gt":
			if template1 > template2 {
				return data, nil
			}
		case "lt":
			if template1 < template2 {
				return data, nil
			}
		case "gte":
			if template1 >= template2 {
				return data, nil
			}
		case "lte":
			if template1 <= template2 {
				return data, nil
			}
		case "contains":
			if strings.Contains(template1, template2) {
				return data, nil
			}
		case "prefix":
			if strings.HasPrefix(template1, template2) {
				return data, nil
			}
		case "suffix":
			if strings.HasSuffix(template1, template2) {
				return data, nil
			}
		}

		return nil, nil
	}

	return &actions.ActionDefinition{
		ID:               FilterID,
		Name:             "Filter",
		Description:      "Filter events based on a criteria, used in combination with pipes or other multi action actions",
		ConfigDefinition: map[string]string{"template1": "string", "template2": "string", "comparison": "eq|gt|lt|gte|lte|contains|prefix|suffix"},
		Handler:          handler,
	}
}
