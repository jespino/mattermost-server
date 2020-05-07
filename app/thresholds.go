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
		&ActiveUsersThreshold{s: s},
		&RegisteredUsersThreshold{s: s},
		&PostsThreshold{s: s},
		&DbConnectionsThreshold{s: s},
		&PushNotificationsThreshold{s: s},
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

// ActiveUsersThreshold is a threshold check for a maximon number of active users
type ActiveUsersThreshold struct {
	s *Server
}

// Interval returns the time between ActiveUsers checks execution
func (au *ActiveUsersThreshold) Interval() time.Duration {
	return time.Hour * 24
}

// TaskName returns the name used for the periodic task
func (au *ActiveUsersThreshold) TaskName() string {
	return "Check Number Of Active Users Threshold "
}

// IsExeeded returns true if the ActiveUsers threshold is exeeded
func (au *ActiveUsersThreshold) IsExeeded() bool {
	//change MONTH_MILLISECONDS to a different value if we want to capture active users for a different period
	noActiveUsers, err := au.s.Store.User().AnalyticsActiveCount(MONTH_MILLISECONDS, model.UserCountOptions{IncludeBotAccounts: false, IncludeDeleted: false})
	if err != nil {
		mlog.Error("Error to get active registered users.", mlog.Err(err))
	}
	return noActiveUsers > 100
}

// Key returns the ActiveUsers threshold unique key
func (au *ActiveUsersThreshold) Key() string {
	return ThresholdPrefix + "active_users"
}

//// Registered Users Threshold

// RegisteredUsersThreshold is a threshold check for a maximon number of registered users
type RegisteredUsersThreshold struct {
	s *Server
}

// Interval returns the time between RegisteredUsers checks execution
func (ru *RegisteredUsersThreshold) Interval() time.Duration {
	return time.Hour * 24
}

// TaskName returns the name used for the periodic task
func (ru *RegisteredUsersThreshold) TaskName() string {
	return "Check Number Of Registered Users Threshold "
}

// IsExeeded returns true if the RegisteredUsers threshold is exeeded
func (ru *RegisteredUsersThreshold) IsExeeded() bool {
	noUsers, err := ru.s.Store.User().Count(model.UserCountOptions{IncludeBotAccounts: false, IncludeDeleted: false})
	if err != nil {
		mlog.Error("Error to get registered users.", mlog.Err(err))
	}
	return noUsers > 250
}

// Key returns the RegisteredUsers threshold unique key
func (ru *RegisteredUsersThreshold) Key() string {
	return ThresholdPrefix + "registered_users"
}

//// Posts Threshold

// PostsThreshold is a threshold check for a maximon number of registered users
type PostsThreshold struct {
	s *Server
}

// Interval returns the time between Posts checks execution
func (p *PostsThreshold) Interval() time.Duration {
	return time.Hour * 24
}

// TaskName returns the name used for the periodic task
func (p *PostsThreshold) TaskName() string {
	return "Check Number Of Posts Threshold "
}

// IsExeeded returns true if the Posts threshold is exeeded
func (p *PostsThreshold) IsExeeded() bool {
	noPosts, err := p.s.Store.Post().AnalyticsPostCount("", false, false)
	if err != nil {
		mlog.Error("Error to get registered users.", mlog.Err(err))
	}
	return noPosts > 200000
}

// Key returns the Posts threshold unique key
func (p *PostsThreshold) Key() string {
	return ThresholdPrefix + "posts"
}

//// Push Proxy notifications Threshold

// PushNotification is a threshold check for a maximon number of registered users
type PushNotificationsThreshold struct {
	s       *Server
	lastDay int64
}

// Interval returns the time between DB Connections checks execution
func (pn *PushNotificationsThreshold) Interval() time.Duration {
	return time.Hour * 24
}

// TaskName returns the name used for the periodic task
func (pn *PushNotificationsThreshold) TaskName() string {
	return "Check Number Of Push Notifications Threshold "
}

// IsExeeded returns true if the DB Connections threshold is exeeded
func (pn *PushNotificationsThreshold) IsExeeded() bool {
	if (pn.s.pushNotificationCounter - pn.lastDay) > 50 {
		pn.lastDay = pn.s.pushNotificationCounter
		return true
	}
	pn.lastDay = pn.s.pushNotificationCounter
	return false
}

// Key returns the DB Connections threshold unique key
func (pn *PushNotificationsThreshold) Key() string {
	return ThresholdPrefix + "push_notifications"
}

//// DB Connections Threshold

// DbConnectionsThreshold is a threshold check for a maximon number of registered users
type DbConnectionsThreshold struct {
	s *Server
}

// Interval returns the time between DB Connections checks execution
func (dc *DbConnectionsThreshold) Interval() time.Duration {
	return time.Minute * 10
}

// TaskName returns the name used for the periodic task
func (dc *DbConnectionsThreshold) TaskName() string {
	return "Check Number Of DB Connections Threshold "
}

// IsExeeded returns true if the DB Connections threshold is exeeded
func (dc *DbConnectionsThreshold) IsExeeded() bool {
	noDbConnections := dc.s.Store.TotalMasterDbConnections()
	return noDbConnections > 100
}

// Key returns the DB Connections threshold unique key
func (dc *DbConnectionsThreshold) Key() string {
	return ThresholdPrefix + "db_connections"
}
