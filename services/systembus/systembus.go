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

type SystemBus interface {
	RegisterEvent(*EventDefinition) error
	Subscribe(subscriber func(*Event) (*Event, error)) error

	SendEvent(*Event) error

	ListEvents() ([]*EventDefinition, error)

	Start() error
	Shutdown() error
}
