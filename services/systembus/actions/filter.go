package actions

import (
	"strings"

	"github.com/mattermost/mattermost-server/v6/services/systembus"
)

const FilterID = "filter"

func NewFilter() *systembus.ActionDefinition {
	handler := func(event *systembus.Event, config map[string]string) (*systembus.Event, error) {
		template1, err := applyTemplate(config["template1"], event.Data)
		if err != nil {
			return nil, err
		}
		template2, err := applyTemplate(config["template2"], event.Data)
		if err != nil {
			return nil, err
		}

		switch config["comparison"] {
		case "":
			if template1 == template2 {
				return event, nil
			}
		case "eq":
			if template1 == template2 {
				return event, nil
			}
		case "gt":
			if template1 > template2 {
				return event, nil
			}
		case "lt":
			if template1 < template2 {
				return event, nil
			}
		case "gte":
			if template1 >= template2 {
				return event, nil
			}
		case "lte":
			if template1 <= template2 {
				return event, nil
			}
		case "contains":
			if strings.Contains(template1, template2) {
				return event, nil
			}
		case "prefix":
			if strings.HasPrefix(template1, template2) {
				return event, nil
			}
		case "suffix":
			if strings.HasSuffix(template1, template2) {
				return event, nil
			}
		}

		return nil, nil
	}

	return &systembus.ActionDefinition{
		ID:               FilterID,
		Name:             "Filter",
		Description:      "Filter events based on a criteria, used in combination with pipes or other multi action actions",
		ConfigDefinition: map[string]string{"template1": "string", "template2": "string", "comparison": "eq|gt|lt|gte|lte|contains|prefix|suffix"},
		Handler:          handler,
	}
}
