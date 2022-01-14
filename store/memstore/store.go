// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package sqlstore

import (
	"context"
	"time"

	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/store"
)

type migrationDirection string

const (
	IndexTypeFullText                 = "full_text"
	IndexTypeFullTextFunc             = "full_text_func"
	IndexTypeDefault                  = "default"
	PGDupTableErrorCode               = "42P07"      // see https://github.com/lib/pq/blob/master/error.go#L268
	MySQLDupTableErrorCode            = uint16(1050) // see https://dev.mysql.com/doc/mysql-errors/5.7/en/server-error-reference.html#error_er_table_exists_error
	PGForeignKeyViolationErrorCode    = "23503"
	MySQLForeignKeyViolationErrorCode = 1452
	PGDuplicateObjectErrorCode        = "42710"
	MySQLDuplicateObjectErrorCode     = 1022
	DBPingAttempts                    = 18
	DBPingTimeoutSecs                 = 10
	// This is a numerical version string by postgres. The format is
	// 2 characters for major, minor, and patch version prior to 10.
	// After 10, it's major and minor only.
	// 10.1 would be 100001.
	// 9.6.3 would be 90603.
	minimumRequiredPostgresVersion = 100000
	// major*1000 + minor*100 + patch
	minimumRequiredMySQLVersion = 5712

	migrationsDirectionUp   migrationDirection = "up"
	migrationsDirectionDown migrationDirection = "down"

	replicaLagPrefix = "replica-lag"
)

type MemStoreStores struct {
	team                 store.TeamStore
	channel              store.ChannelStore
	post                 store.PostStore
	retentionPolicy      store.RetentionPolicyStore
	thread               store.ThreadStore
	user                 store.UserStore
	bot                  store.BotStore
	audit                store.AuditStore
	cluster              store.ClusterDiscoveryStore
	remoteCluster        store.RemoteClusterStore
	compliance           store.ComplianceStore
	session              store.SessionStore
	oauth                store.OAuthStore
	system               store.SystemStore
	webhook              store.WebhookStore
	command              store.CommandStore
	commandWebhook       store.CommandWebhookStore
	preference           store.PreferenceStore
	license              store.LicenseStore
	token                store.TokenStore
	emoji                store.EmojiStore
	status               store.StatusStore
	fileInfo             store.FileInfoStore
	uploadSession        store.UploadSessionStore
	reaction             store.ReactionStore
	job                  store.JobStore
	userAccessToken      store.UserAccessTokenStore
	plugin               store.PluginStore
	channelMemberHistory store.ChannelMemberHistoryStore
	role                 store.RoleStore
	scheme               store.SchemeStore
	TermsOfService       store.TermsOfServiceStore
	productNotices       store.ProductNoticesStore
	group                store.GroupStore
	UserTermsOfService   store.UserTermsOfServiceStore
	linkMetadata         store.LinkMetadataStore
	sharedchannel        store.SharedChannelStore
}

type MemStore struct {
	stores MemStoreStores

	context context.Context
}

// ColumnInfo holds information about a column.
type ColumnInfo struct {
	DataType          string
	DefaultValue      string
	CharMaximumLength int
}

func New() *MemStore {
	store := &MemStore{}
	store.stores.team = nil
	store.stores.channel = nil
	store.stores.post = nil
	store.stores.retentionPolicy = nil
	store.stores.user = nil
	store.stores.bot = nil
	store.stores.audit = nil
	store.stores.cluster = nil
	store.stores.remoteCluster = nil
	store.stores.compliance = nil
	store.stores.session = nil
	store.stores.oauth = nil
	store.stores.system = newMemSystemStore()
	store.stores.webhook = newMemWebhookStore()
	store.stores.command = newMemCommandStore()
	store.stores.commandWebhook = nil
	store.stores.preference = nil
	store.stores.license = nil
	store.stores.token = nil
	store.stores.emoji = nil
	store.stores.status = nil
	store.stores.fileInfo = nil
	store.stores.uploadSession = nil
	store.stores.thread = nil
	store.stores.job = nil
	store.stores.userAccessToken = nil
	store.stores.channelMemberHistory = nil
	store.stores.plugin = nil
	store.stores.TermsOfService = nil
	store.stores.UserTermsOfService = nil
	store.stores.linkMetadata = nil
	store.stores.sharedchannel = nil
	store.stores.reaction = nil
	store.stores.role = nil
	store.stores.scheme = nil
	store.stores.group = nil
	store.stores.productNotices = nil

	return store
}

func (ss *MemStore) SetContext(context context.Context) {
	ss.context = context
}

func (ss *MemStore) Context() context.Context {
	return ss.context
}

func (ss *MemStore) GetCurrentSchemaVersion() string {
	return "fake-mem-db"
}

// GetDbVersion returns the version of the database being used.
// If numerical is set to true, it attempts to return a numerical version string
// that can be parsed by callers.
func (ss *MemStore) GetDbVersion(numerical bool) (string, error) {
	return "fake-mem-db", nil

}

func (ss *MemStore) TotalMasterDbConnections() int {
	return 1
}

// ReplicaLagAbs queries all the replica databases to get the absolute replica lag value
// and updates the Prometheus metric with it.
func (ss *MemStore) ReplicaLagAbs() error {
	return nil
}

// ReplicaLagAbs queries all the replica databases to get the time-based replica lag value
// and updates the Prometheus metric with it.
func (ss *MemStore) ReplicaLagTime() error {
	return nil
}

func (ss *MemStore) TotalReadDbConnections() int {
	return 0
}

func (ss *MemStore) TotalSearchDbConnections() int {
	return 0
}

func (ss *MemStore) MarkSystemRanUnitTests()              {}
func (ss *MemStore) RecycleDBConnections(d time.Duration) {}
func (ss *MemStore) Close()                               {}
func (ss *MemStore) LockToMaster()                        {}
func (ss *MemStore) UnlockFromMaster()                    {}

func (ss *MemStore) Team() store.TeamStore {
	return ss.stores.team
}

func (ss *MemStore) Channel() store.ChannelStore {
	return ss.stores.channel
}

func (ss *MemStore) Post() store.PostStore {
	return ss.stores.post
}

func (ss *MemStore) RetentionPolicy() store.RetentionPolicyStore {
	return ss.stores.retentionPolicy
}

func (ss *MemStore) User() store.UserStore {
	return ss.stores.user
}

func (ss *MemStore) Bot() store.BotStore {
	return ss.stores.bot
}

func (ss *MemStore) Session() store.SessionStore {
	return ss.stores.session
}

func (ss *MemStore) Audit() store.AuditStore {
	return ss.stores.audit
}

func (ss *MemStore) ClusterDiscovery() store.ClusterDiscoveryStore {
	return ss.stores.cluster
}

func (ss *MemStore) RemoteCluster() store.RemoteClusterStore {
	return ss.stores.remoteCluster
}

func (ss *MemStore) Compliance() store.ComplianceStore {
	return ss.stores.compliance
}

func (ss *MemStore) OAuth() store.OAuthStore {
	return ss.stores.oauth
}

func (ss *MemStore) System() store.SystemStore {
	return ss.stores.system
}

func (ss *MemStore) Webhook() store.WebhookStore {
	return ss.stores.webhook
}

func (ss *MemStore) Command() store.CommandStore {
	return ss.stores.command
}

func (ss *MemStore) CommandWebhook() store.CommandWebhookStore {
	return ss.stores.commandWebhook
}

func (ss *MemStore) Preference() store.PreferenceStore {
	return ss.stores.preference
}

func (ss *MemStore) License() store.LicenseStore {
	return ss.stores.license
}

func (ss *MemStore) Token() store.TokenStore {
	return ss.stores.token
}

func (ss *MemStore) Emoji() store.EmojiStore {
	return ss.stores.emoji
}

func (ss *MemStore) Status() store.StatusStore {
	return ss.stores.status
}

func (ss *MemStore) FileInfo() store.FileInfoStore {
	return ss.stores.fileInfo
}

func (ss *MemStore) UploadSession() store.UploadSessionStore {
	return ss.stores.uploadSession
}

func (ss *MemStore) Reaction() store.ReactionStore {
	return ss.stores.reaction
}

func (ss *MemStore) Job() store.JobStore {
	return ss.stores.job
}

func (ss *MemStore) UserAccessToken() store.UserAccessTokenStore {
	return ss.stores.userAccessToken
}

func (ss *MemStore) ChannelMemberHistory() store.ChannelMemberHistoryStore {
	return ss.stores.channelMemberHistory
}

func (ss *MemStore) Plugin() store.PluginStore {
	return ss.stores.plugin
}

func (ss *MemStore) Thread() store.ThreadStore {
	return ss.stores.thread
}

func (ss *MemStore) Role() store.RoleStore {
	return ss.stores.role
}

func (ss *MemStore) TermsOfService() store.TermsOfServiceStore {
	return ss.stores.TermsOfService
}

func (ss *MemStore) ProductNotices() store.ProductNoticesStore {
	return ss.stores.productNotices
}

func (ss *MemStore) UserTermsOfService() store.UserTermsOfServiceStore {
	return ss.stores.UserTermsOfService
}

func (ss *MemStore) Scheme() store.SchemeStore {
	return ss.stores.scheme
}

func (ss *MemStore) Group() store.GroupStore {
	return ss.stores.group
}

func (ss *MemStore) LinkMetadata() store.LinkMetadataStore {
	return ss.stores.linkMetadata
}

func (ss *MemStore) SharedChannel() store.SharedChannelStore {
	return ss.stores.sharedchannel
}

func (ss *MemStore) DropAllTables() {}

func (ss *MemStore) CheckIntegrity() <-chan model.IntegrityCheckResult {
	results := make(chan model.IntegrityCheckResult)
	close(results)
	return results
}
