package watchman

import (
	"sync"

	"github.com/mattermost/mattermost-server/v6/services/systembus"
	"github.com/mattermost/mattermost-server/v6/shared/mlog"
)

type LoginError = {
	millis: int64
	count: int
}

type Watchman struct {
	systemBus   systembus.SystemBus
	mutex       sync.RWMutex
	logger      *mlog.Logger
	incidents   []string
	loginErrors map[string]LoginError
}

func New(logger *mlog.Logger, systemBus systembus.SystemBus) *Watchman {
	watchman := &Watchman{
		systemBus: systemBus,
		logger:    logger,
	}
	systemBus.Subscribe(watchman.handleEvent)
	return watchman
}

func (w *Watchman) handleEvent(event *systembus.Event) (*systembus.Event, error) {
	switch event.ID {
	case "login-success":
		w.handleLoginSuccess(event)
	case "login-failure":
		w.handleLoginFailure(event)
	}
}

func (w *Watchman) Start() error {
	return nil
}

func (w *Watchman) Shutdown() error {
	return nil
}

func (w *Watchman) handleLoginSuccess(event *systembus.Event) {
	w.loginErrors[event.Data["Username"]] = LoginError{}

	// TODO: Notify the user (And maybe block the account)
}

func (w *Watchman) handleLoginFailure(event *systembus.Event) {
	w.incidents = append(w.incidents, fmt.Sprintf(
		"Login failure for username %s (UserAgent: %s, IP: %s)",
		event.Data["UsernameInput"],
		event.Data["UserAgent"],
		event.Data["IPAddress"]
	)

	lastLoginFailure, ok := w.loginErrors[event.Data["UsernameInput"]]
	if !ok {
		w.loginErrors[event.Data["UsernameInput"]] = LoginError{millis: model.GetMillis(), count: 1}
		return
	}
	now := model.GetMillis()

	// If the last login failure is older than 5 minutes
	if (now - lastLoginFailure.millis) > (5*60*60*1000) {
		lastLoginFailure.count = 0
	}
	lastLoginFailure.count += 1
	lastLoginFailure.millis = model.GetMillis
	w.loginErrors[event.Data["UsernameInput"]] = lastLoginFailure

	if lastLoginFailure.count > 2 {
		// TODO: Notify the user if it exists (And maybe block the account)
	}
}
