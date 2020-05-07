package app

import (
	"time"

	"github.com/mattermost/mattermost-server/v5/mlog"
	"github.com/mattermost/mattermost-server/v5/model"
)

// ThresholdPrefix represent the prefix used for any threshold information key
const ThresholdPrefix = "threshold_"

// StartThresholds start all the threshold periodic checks
func StartThresholds(s *Server) {
	thresholds := []Threshold{
		&ActiveUsers{s: s},
		&RegisteredUsers{s: s},
	}
	for _, threshold := range thresholds {
		RunThreshold(threshold, s)
		ScheduleThreshold(threshold, s)
	}
}

func RunThreshold(threshold Threshold, server *Server) {
	data, err := server.Store.System().GetByName(threshold.Key())
	if err == nil && data.Value == "ack" || data.Value == "true" {
		return
	}

	if threshold.IsExeeded() {
		if err := server.Store.System().SaveOrUpdate(&model.System{Name: threshold.Key(), Value: "true"}); err != nil {
			mlog.Error("Unable to write to database.", mlog.Err(err))
			return
		}
		message := model.NewWebSocketEvent(model.WEBSOCKET_THRESHOLD_EXEEDED, "", "", "", nil)
		message.Add(threshold.Key(), "true")
		server.FakeApp().Publish(message)
		return
	}
}

func ScheduleThreshold(threshold Threshold, server *Server) {
	model.CreateRecurringTask(threshold.TaskName(), func() {
		RunThreshold(threshold, server)
	}, threshold.Interval())
}

// Threshold represent any runnable check for threshold surpassing
type Threshold interface {
	IsExeeded() bool
	Key() string
	Interval() time.Duration
	TaskName() string
}

//// Active Users Threshold

// ActiveUsers is a threshold check for a maximon number of active users
type ActiveUsers struct {
	s *Server
}

// Interval returns the time between ActiveUsers checks execution
func (au *ActiveUsers) Interval() time.Duration {
	return time.Hour * 24
}

// TaskName returns the name used for the periodic task
func (au *ActiveUsers) TaskName() string {
	return "Check Number Of Active Users Threshold "
}

// IsExeeded returns true if the ActiveUsers threshold is exeeded
func (au *ActiveUsers) IsExeeded() bool {
	//change MONTH_MILLISECONDS to a different value if we want to capture active users for a different period
	noActiveUsers, err := au.s.Store.User().AnalyticsActiveCount(MONTH_MILLISECONDS, model.UserCountOptions{IncludeBotAccounts: false, IncludeDeleted: false})
	if err != nil {
		mlog.Error("Error to get active registered users.", mlog.Err(err))
	}
	return noActiveUsers > 100
}

// Key returns the ActiveUsers threshold unique key
func (au *ActiveUsers) Key() string {
	return ThresholdPrefix + "active_users"
}

//// Registered Users Threshold

// RegisteredUsers is a threshold check for a maximon number of registered users
type RegisteredUsers struct {
	s *Server
}

// Interval returns the time between RegisteredUsers checks execution
func (ru *RegisteredUsers) Interval() time.Duration {
	return time.Hour * 24
}

// TaskName returns the name used for the periodic task
func (ru *RegisteredUsers) TaskName() string {
	return "Check Number Of Registered Users Threshold "
}

// IsExeeded returns true if the RegisteredUsers threshold is exeeded
func (ru *RegisteredUsers) IsExeeded() bool {
	noUsers, err := ru.s.Store.User().Count(model.UserCountOptions{IncludeBotAccounts: false, IncludeDeleted: false})
	if err != nil {
		mlog.Error("Error to get registered users.", mlog.Err(err))
	}
	return noUsers > 250
}

// Key returns the RegisteredUsers threshold unique key
func (ru *RegisteredUsers) Key() string {
	return ThresholdPrefix + "registered_users"
}

//// Posts Threshold

// Posts is a threshold check for a maximon number of registered users
type Posts struct {
	s *Server
}

// Interval returns the time between Posts checks execution
func (p *Posts) Interval() time.Duration {
	return time.Hour * 24
}

// TaskName returns the name used for the periodic task
func (p *Posts) TaskName() string {
	return "Check Number Of Posts Threshold "
}

// IsExeeded returns true if the Posts threshold is exeeded
func (p *Posts) IsExeeded() bool {
	noPosts, err := p.s.Store.Post().AnalyticsPostCount("", false, false)
	if err != nil {
		mlog.Error("Error to get registered users.", mlog.Err(err))
	}
	return noPosts > 200000
}

// Key returns the Posts threshold unique key
func (p *Posts) Key() string {
	return ThresholdPrefix + "posts"
}
