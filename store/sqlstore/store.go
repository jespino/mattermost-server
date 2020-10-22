// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package sqlstore

import (
	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/mattermost/gorp"
)

/*type SqlStore struct {
	master         *gorp.DbMap
	replicas       []*gorp.DbMap
	searchReplicas []*gorp.DbMap
	team           TeamStore
	channel        ChannelStore
	post           PostStore
	user           UserStore
	audit          AuditStore
	compliance     ComplianceStore
	session        SessionStore
	oauth          OAuthStore
	system         SystemStore
	webhook        WebhookStore
	command        CommandStore
	preference     PreferenceStore
	license        LicenseStore
	token          TokenStore
	emoji          EmojiStore
	status         StatusStore
	fileInfo       FileInfoStore
	reaction       ReactionStore
	jobStatus      JobStatusStore
	SchemaVersion  string
	rrCounter      int64
	srCounter      int64
}*/

type SqlStore interface {
	DriverName() string
	GetCurrentSchemaVersion() string
	GetMaster() *gorp.DbMap
	GetSearchReplica() *gorp.DbMap
	GetReplica() *gorp.DbMap
	GetDbVersion() (string, error)
	TotalMasterDbConnections() int
	TotalReadDbConnections() int
	TotalSearchDbConnections() int
	MarkSystemRanUnitTests()
	DoesTableExist(tablename string) bool
	DoesColumnExist(tableName string, columName string) bool
	DoesTriggerExist(triggerName string) bool
	CreateColumnIfNotExists(tableName string, columnName string, mySqlColType string, postgresColType string, defaultValue string) bool
	CreateColumnIfNotExistsNoDefault(tableName string, columnName string, mySqlColType string, postgresColType string) bool
	RemoveColumnIfExists(tableName string, columnName string) bool
	RemoveTableIfExists(tableName string) bool
	RenameColumnIfExists(tableName string, oldColumnName string, newColumnName string, colType string) bool
	GetMaxLengthOfColumnIfExists(tableName string, columnName string) string
	AlterColumnTypeIfExists(tableName string, columnName string, mySqlColType string, postgresColType string) bool
	AlterColumnDefaultIfExists(tableName string, columnName string, mySqlColDefault *string, postgresColDefault *string) bool
	AlterPrimaryKey(tableName string, columnNames []string) bool
	CreateUniqueIndexIfNotExists(indexName string, tableName string, columnName string) bool
	CreateIndexIfNotExists(indexName string, tableName string, columnName string) bool
	CreateCompositeIndexIfNotExists(indexName string, tableName string, columnNames []string) bool
	CreateUniqueCompositeIndexIfNotExists(indexName string, tableName string, columnNames []string) bool
	CreateFullTextIndexIfNotExists(indexName string, tableName string, columnName string) bool
	RemoveIndexIfExists(indexName string, tableName string) bool
	GetAllConns() []*gorp.DbMap
	Close()
	LockToMaster()
	UnlockFromMaster()
	Team() *SqlTeamStore
	Channel() *SqlChannelStore
	Post() *SqlPostStore
	Thread() *SqlThreadStore
	User() *SqlUserStore
	Bot() *SqlBotStore
	Audit() *SqlAuditStore
	ClusterDiscovery() *SqlClusterDiscoveryStore
	Compliance() *SqlComplianceStore
	Session() *SqlSessionStore
	OAuth() *SqlOAuthStore
	System() *SqlSystemStore
	Webhook() *SqlWebhookStore
	Command() *SqlCommandStore
	CommandWebhook() *SqlCommandWebhookStore
	Preference() *SqlPreferenceStore
	License() *SqlLicenseStore
	Token() *SqlTokenStore
	Emoji() *SqlEmojiStore
	Status() *SqlStatusStore
	FileInfo() *SqlFileInfoStore
	UploadSession() *SqlUploadSessionStore
	Reaction() *SqlReactionStore
	Job() *SqlJobStore
	Plugin() *SqlPluginStore
	UserAccessToken() *SqlUserAccessTokenStore
	Role() *SqlRoleStore
	Scheme() *SqlSchemeStore
	TermsOfService() *SqlTermsOfServiceStore
	UserTermsOfService() *SqlUserTermsOfServiceStore
	LinkMetadata() *SqlLinkMetadataStore
	getQueryBuilder() sq.StatementBuilderType
}
