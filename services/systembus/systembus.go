package systembus

type EventDefinition struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Fields      []string `json:"fields"`
}

type Event struct {
	ID   string            `json:"id"`
	Data map[string]string `json:"data"`
}

type ActionDefinition struct {
	ID               string                                                       `json:"id"`
	Name             string                                                       `json:"name"`
	Description      string                                                       `json:"description"`
	ConfigDefinition map[string]string                                            `json:"config_definition"`
	Handler          func(event *Event, config map[string]string) (*Event, error) `json:"-"`
}

type LinkEventAction struct {
	ID       string            `json:"id"`
	EventID  string            `json:"event_id"`
	ActionID string            `json:"action_id"`
	Config   map[string]string `json:"config"`
}

type SystemBus interface {
	RegisterEvent(*EventDefinition) error
	RegisterAction(*ActionDefinition) error

	SendEvent(*Event) error

	ListEvents() ([]*EventDefinition, error)
	ListActions() ([]*ActionDefinition, error)
	ListLinks() ([]*LinkEventAction, error)

	LinkEventAction(eventID string, actionID string, config map[string]string) (*LinkEventAction, error)
	UnlinkEventAction(linkID string) error

	Start() error
	Shutdown() error
}
