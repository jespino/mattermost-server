package systembus

type EventDefinition struct {
	ID          string
	Name        string
	Description string
	Fields      []string
}

type Event struct {
	ID   string
	Data map[string]string
}

type ActionDefinition struct {
	ID               string
	Name             string
	Description      string
	ConfigDefinition map[string]string
	Handler          func(event *Event, config map[string]string) (*Event, error)
}

type LinkEventAction struct {
	ID       string
	EventID  string
	ActionID string
	Config   map[string]string
}

type SystemBus interface {
	RegisterEvent(*EventDefinition) error
	RegisterAction(*ActionDefinition) error

	SendEvent(*Event) error

	ListEvents() ([]*EventDefinition, error)
	ListActions() ([]*ActionDefinition, error)
	ListLinks() ([]*LinkEventAction, error)

	LinkEventAction(eventID string, actionID string, config map[string]string) error
	UnlinkEventAction(linkID string) error

	Start() error
	Shutdown() error
}
