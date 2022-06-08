package actions

import (
	"bytes"
	"fmt"
	"sync"
	"text/template"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/services/systembus"
	"github.com/mattermost/mattermost-server/v6/shared/mlog"
)

type ActionDefinition struct {
	ID               string                                                                            `json:"id"`
	Name             string                                                                            `json:"name"`
	Description      string                                                                            `json:"description"`
	ConfigDefinition map[string]string                                                                 `json:"config_definition"`
	Handler          func(config map[string]string, data map[string]string) (map[string]string, error) `json:"-"`
}

type LinkEventAction struct {
	ID       string            `json:"id"`
	EventID  string            `json:"event_id"`
	ActionID string            `json:"action_id"`
	Config   map[string]string `json:"config"`
}

func (lea *LinkEventAction) ApplyConfigTemplates(event *systembus.Event) (map[string]string, error) {
	result := map[string]string{}
	for key, value := range lea.Config {
		tmpl, err := template.New("template").Parse(value)
		if err != nil {
			return nil, err
		}
		var buf bytes.Buffer
		tmpl.Execute(&buf, event.Data)
		result[key] = buf.String()
	}
	return result, nil
}

type Actions struct {
	systemBus    systembus.SystemBus
	actions      map[string]*ActionDefinition
	links        map[string]*LinkEventAction
	linksByEvent map[string]map[string]*LinkEventAction
	mutex        sync.RWMutex
	logger       *mlog.Logger
}

func EventToActionSubscription(a *Actions) func(event *systembus.Event) (*systembus.Event, error) {
	return func(event *systembus.Event) (*systembus.Event, error) {
		a.mutex.RLock()
		defer a.mutex.RUnlock()

		links, ok := a.linksByEvent[event.ID]
		if !ok {
			return event, nil
		}

		for _, link := range links {
			config, err := link.ApplyConfigTemplates(event)
			if err != nil {
				a.logger.Debug("Invalid link config", mlog.Err(err))
				continue
			}

			_, err = a.Run(link.ActionID, config, event.Data)
			if err != nil {
				a.logger.Debug("Action failed", mlog.Err(err))
				continue
			}
		}
		return nil, nil
	}
}

func New(logger *mlog.Logger, systemBus systembus.SystemBus) *Actions {
	actions := &Actions{
		systemBus:    systemBus,
		actions:      map[string]*ActionDefinition{},
		links:        map[string]*LinkEventAction{},
		linksByEvent: map[string]map[string]*LinkEventAction{},
		logger:       logger,
	}
	systemBus.Subscribe(EventToActionSubscription(actions))
	return actions
}

func (a *Actions) RegisterAction(action *ActionDefinition) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.actions[action.ID] = action
	return nil
}

func (a *Actions) ListActions() ([]*ActionDefinition, error) {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	actions := []*ActionDefinition{}
	for _, action := range a.actions {
		actions = append(actions, action)
	}
	return actions, nil
}

func (a *Actions) ListLinks() ([]*LinkEventAction, error) {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	links := []*LinkEventAction{}
	for _, link := range a.links {
		links = append(links, link)
	}
	return links, nil
}

func (a *Actions) Run(actionID string, config map[string]string, data map[string]string) (map[string]string, error) {
	action, ok := a.actions[actionID]
	if !ok {
		return nil, fmt.Errorf("Action %s not found", actionID)
	}
	return action.Handler(config, data)
}

func (a *Actions) LinkEventAction(eventID string, actionID string, config map[string]string) (*LinkEventAction, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	newLink := LinkEventAction{
		ID:       model.NewId(),
		EventID:  eventID,
		ActionID: actionID,
		Config:   config,
	}

	a.links[newLink.ID] = &newLink
	if _, ok := a.linksByEvent[newLink.EventID]; ok {
		a.linksByEvent[newLink.EventID][newLink.ID] = &newLink
	} else {
		a.linksByEvent[newLink.EventID] = map[string]*LinkEventAction{newLink.ID: &newLink}
	}
	return &newLink, nil
}

func (a *Actions) UnlinkEventAction(linkID string) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	link := a.links[linkID]
	delete(a.linksByEvent[link.EventID], link.ID)
	delete(a.links, link.ID)
	return nil
}

func (a *Actions) Start() error {
	return nil
}

func (a *Actions) Shutdown() error {
	return nil
}
