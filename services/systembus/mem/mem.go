package mem

import (
	"sync"

	"github.com/mattermost/mattermost-server/v6/services/systembus"
	"github.com/mattermost/mattermost-server/v6/shared/mlog"
)

type SystemBus struct {
	events      map[string]*systembus.EventDefinition
	logger      *mlog.Logger
	mutex       sync.RWMutex
	subscribers []func(event *systembus.Event) (*systembus.Event, error)
}

func New(logger *mlog.Logger) *SystemBus {
	return &SystemBus{
		events: map[string]*systembus.EventDefinition{},
		logger: logger,
	}
}

func (sb *SystemBus) RegisterEvent(event *systembus.EventDefinition) error {
	sb.mutex.Lock()
	defer sb.mutex.Unlock()
	sb.events[event.ID] = event
	return nil
}

func (sb *SystemBus) Subscribe(subscriber func(*systembus.Event) (*systembus.Event, error)) error {
	sb.mutex.Lock()
	defer sb.mutex.Unlock()
	sb.subscribers = append(sb.subscribers, subscriber)
	return nil
}

func (sb *SystemBus) SendEvent(event *systembus.Event) error {
	sb.mutex.RLock()
	defer sb.mutex.RUnlock()

	for _, subscriber := range sb.subscribers {
		_, err := subscriber(event)
		if err != nil {
			sb.logger.Debug("Subscriber failed", mlog.Err(err))
		}
	}
	return nil
}

func (sb *SystemBus) ListEvents() ([]*systembus.EventDefinition, error) {
	sb.mutex.RLock()
	defer sb.mutex.RUnlock()
	events := []*systembus.EventDefinition{}
	for _, event := range sb.events {
		events = append(events, event)
	}
	return events, nil
}

func (sb *SystemBus) Start() error {
	return nil
}

func (sb *SystemBus) Shutdown() error {
	return nil
}
