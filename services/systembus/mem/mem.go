package mem

import (
	"sync"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/services/systembus"
)

type SystemBus struct {
	events       map[string]*systembus.EventDefinition
	actions      map[string]*systembus.ActionDefinition
	links        map[string]*systembus.LinkEventAction
	linksByEvent map[string]map[string]*systembus.LinkEventAction
	mutex        sync.RWMutex
}

func New() *SystemBus {
	return &SystemBus{
		events:       map[string]*systembus.EventDefinition{},
		actions:      map[string]*systembus.ActionDefinition{},
		links:        map[string]*systembus.LinkEventAction{},
		linksByEvent: map[string]map[string]*systembus.LinkEventAction{},
	}
}

func (es *SystemBus) RegisterEvent(event *systembus.EventDefinition) error {
	es.mutex.Lock()
	defer es.mutex.Unlock()
	es.events[event.ID] = event
	return nil
}

func (es *SystemBus) RegisterAction(action *systembus.ActionDefinition) error {
	es.mutex.Lock()
	defer es.mutex.Unlock()
	es.actions[action.ID] = action
	return nil
}

func (es *SystemBus) SendEvent(event *systembus.Event) error {
	es.mutex.RLock()
	defer es.mutex.RUnlock()
	links, ok := es.linksByEvent[event.ID]
	if !ok {
		return nil
	}

	for _, link := range links {
		action := es.actions[link.ActionID]
		action.Handler(event, link.Config)
	}
	return nil
}

func (es *SystemBus) ListEvents() ([]*systembus.EventDefinition, error) {
	es.mutex.RLock()
	defer es.mutex.RUnlock()
	events := []*systembus.EventDefinition{}
	for _, event := range es.events {
		events = append(events, event)
	}
	return events, nil
}

func (es *SystemBus) ListActions() ([]*systembus.ActionDefinition, error) {
	es.mutex.RLock()
	defer es.mutex.RUnlock()
	actions := []*systembus.ActionDefinition{}
	for _, action := range es.actions {
		actions = append(actions, action)
	}
	return actions, nil
}

func (es *SystemBus) ListLinks() ([]*systembus.LinkEventAction, error) {
	es.mutex.RLock()
	defer es.mutex.RUnlock()
	links := []*systembus.LinkEventAction{}
	for _, link := range es.links {
		links = append(links, link)
	}
	return links, nil
}

func (es *SystemBus) LinkEventAction(eventID string, actionID string, config map[string]string) error {
	es.mutex.Lock()
	defer es.mutex.Unlock()
	newLink := systembus.LinkEventAction{
		ID:       model.NewId(),
		EventID:  eventID,
		ActionID: actionID,
		Config:   config,
	}

	es.links[newLink.ID] = &newLink
	if _, ok := es.linksByEvent[newLink.EventID]; ok {
		es.linksByEvent[newLink.EventID][newLink.ID] = &newLink
	} else {
		es.linksByEvent[newLink.EventID] = map[string]*systembus.LinkEventAction{newLink.ID: &newLink}
	}
	return nil
}

func (es *SystemBus) UnlinkEventAction(linkID string) error {
	es.mutex.Lock()
	defer es.mutex.Unlock()
	link := es.links[linkID]
	delete(es.linksByEvent[link.EventID], link.ID)
	delete(es.links, link.ID)
	return nil
}

func (es *SystemBus) Start() error {
	return nil
}

func (es *SystemBus) Shutdown() error {
	return nil
}
