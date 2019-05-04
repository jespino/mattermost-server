package kafkastore

import (
	"context"
	"encoding/json"

	"github.com/mattermost/mattermost-server/store"
	kafka "github.com/segmentio/kafka-go"
)

type KafkaStore struct {
	baseStore            store.Store
	conn                 *kafka.Conn
	team                 store.TeamStore
	channel              store.ChannelStore
	post                 store.PostStore
	user                 store.UserStore
	bot                  store.BotStore
	audit                store.AuditStore
	clusterDiscovery     store.ClusterDiscoveryStore
	compliance           store.ComplianceStore
	session              store.SessionStore
	oAuth                store.OAuthStore
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
	reaction             store.ReactionStore
	role                 store.RoleStore
	scheme               store.SchemeStore
	job                  store.JobStore
	userAccessToken      store.UserAccessTokenStore
	channelMemberHistory store.ChannelMemberHistoryStore
	plugin               store.PluginStore
	termsOfService       store.TermsOfServiceStore
	group                store.GroupStore
	userTermsOfService   store.UserTermsOfServiceStore
	linkMetadata         store.LinkMetadataStore
	notificationRegistry store.NotificationRegistryStore
}

func NewKafkaSupplier(baseStore store.Store) KafkaStore {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "mattermost-store", 0)
	kafkaStore := KafkaStore{
		baseStore: baseStore,
		conn:      conn,
	}

	kafkaStore.team = KafkaTeamStore{baseStore: baseStore.Team(), root: &kafkaStore}
	kafkaStore.channel = KafkaChannelStore{baseStore: baseStore.Channel(), root: &kafkaStore}
	kafkaStore.post = KafkaPostStore{baseStore: baseStore.Post(), root: &kafkaStore}
	kafkaStore.user = KafkaUserStore{baseStore: baseStore.User(), root: &kafkaStore}
	kafkaStore.bot = KafkaBotStore{baseStore: baseStore.Bot(), root: &kafkaStore}
	kafkaStore.audit = KafkaAuditStore{baseStore: baseStore.Audit(), root: &kafkaStore}
	kafkaStore.clusterDiscovery = KafkaClusterDiscoveryStore{baseStore: baseStore.ClusterDiscovery(), root: &kafkaStore}
	kafkaStore.compliance = KafkaComplianceStore{baseStore: baseStore.Compliance(), root: &kafkaStore}
	kafkaStore.session = KafkaSessionStore{baseStore: baseStore.Session(), root: &kafkaStore}
	kafkaStore.oAuth = KafkaOAuthStore{baseStore: baseStore.OAuth(), root: &kafkaStore}
	kafkaStore.system = KafkaSystemStore{baseStore: baseStore.System(), root: &kafkaStore}
	kafkaStore.webhook = KafkaWebhookStore{baseStore: baseStore.Webhook(), root: &kafkaStore}
	kafkaStore.command = KafkaCommandStore{baseStore: baseStore.Command(), root: &kafkaStore}
	kafkaStore.commandWebhook = KafkaCommandWebhookStore{baseStore: baseStore.CommandWebhook(), root: &kafkaStore}
	kafkaStore.preference = KafkaPreferenceStore{baseStore: baseStore.Preference(), root: &kafkaStore}
	kafkaStore.license = KafkaLicenseStore{baseStore: baseStore.License(), root: &kafkaStore}
	kafkaStore.token = KafkaTokenStore{baseStore: baseStore.Token(), root: &kafkaStore}
	kafkaStore.emoji = KafkaEmojiStore{baseStore: baseStore.Emoji(), root: &kafkaStore}
	kafkaStore.status = KafkaStatusStore{baseStore: baseStore.Status(), root: &kafkaStore}
	kafkaStore.fileInfo = KafkaFileInfoStore{baseStore: baseStore.FileInfo(), root: &kafkaStore}
	kafkaStore.reaction = KafkaReactionStore{baseStore: baseStore.Reaction(), root: &kafkaStore}
	kafkaStore.role = KafkaRoleStore{baseStore: baseStore.Role(), root: &kafkaStore}
	kafkaStore.scheme = KafkaSchemeStore{baseStore: baseStore.Scheme(), root: &kafkaStore}
	kafkaStore.job = KafkaJobStore{baseStore: baseStore.Job(), root: &kafkaStore}
	kafkaStore.userAccessToken = KafkaUserAccessTokenStore{baseStore: baseStore.UserAccessToken(), root: &kafkaStore}
	kafkaStore.channelMemberHistory = KafkaChannelMemberHistoryStore{baseStore: baseStore.ChannelMemberHistory(), root: &kafkaStore}
	kafkaStore.plugin = KafkaPluginStore{baseStore: baseStore.Plugin(), root: &kafkaStore}
	kafkaStore.termsOfService = KafkaTermsOfServiceStore{baseStore: baseStore.TermsOfService(), root: &kafkaStore}
	kafkaStore.group = KafkaGroupStore{baseStore: baseStore.Group(), root: &kafkaStore}
	kafkaStore.userTermsOfService = KafkaUserTermsOfServiceStore{baseStore: baseStore.UserTermsOfService(), root: &kafkaStore}
	kafkaStore.linkMetadata = KafkaLinkMetadataStore{baseStore: baseStore.LinkMetadata(), root: &kafkaStore}
	kafkaStore.notificationRegistry = KafkaNotificationRegistryStore{baseStore: baseStore.NotificationRegistry(), root: &kafkaStore}
	return kafkaStore
}

func (s KafkaStore) Team() store.TeamStore {
	return s.team
}

func (s KafkaStore) Channel() store.ChannelStore {
	return s.channel
}

func (s KafkaStore) Post() store.PostStore {
	return s.post
}

func (s KafkaStore) User() store.UserStore {
	return s.user
}

func (s KafkaStore) Bot() store.BotStore {
	return s.bot
}

func (s KafkaStore) Audit() store.AuditStore {
	return s.audit
}

func (s KafkaStore) ClusterDiscovery() store.ClusterDiscoveryStore {
	return s.clusterDiscovery
}

func (s KafkaStore) Compliance() store.ComplianceStore {
	return s.compliance
}

func (s KafkaStore) Session() store.SessionStore {
	return s.session
}

func (s KafkaStore) OAuth() store.OAuthStore {
	return s.oAuth
}

func (s KafkaStore) System() store.SystemStore {
	return s.system
}

func (s KafkaStore) Webhook() store.WebhookStore {
	return s.webhook
}

func (s KafkaStore) Command() store.CommandStore {
	return s.command
}

func (s KafkaStore) CommandWebhook() store.CommandWebhookStore {
	return s.commandWebhook
}

func (s KafkaStore) Preference() store.PreferenceStore {
	return s.preference
}

func (s KafkaStore) License() store.LicenseStore {
	return s.license
}

func (s KafkaStore) Token() store.TokenStore {
	return s.token
}

func (s KafkaStore) Emoji() store.EmojiStore {
	return s.emoji
}

func (s KafkaStore) Status() store.StatusStore {
	return s.status
}

func (s KafkaStore) FileInfo() store.FileInfoStore {
	return s.fileInfo
}

func (s KafkaStore) Reaction() store.ReactionStore {
	return s.reaction
}

func (s KafkaStore) Role() store.RoleStore {
	return s.role
}

func (s KafkaStore) Scheme() store.SchemeStore {
	return s.scheme
}

func (s KafkaStore) Job() store.JobStore {
	return s.job
}

func (s KafkaStore) UserAccessToken() store.UserAccessTokenStore {
	return s.userAccessToken
}

func (s KafkaStore) ChannelMemberHistory() store.ChannelMemberHistoryStore {
	return s.channelMemberHistory
}

func (s KafkaStore) Plugin() store.PluginStore {
	return s.plugin
}

func (s KafkaStore) TermsOfService() store.TermsOfServiceStore {
	return s.termsOfService
}

func (s KafkaStore) Group() store.GroupStore {
	return s.group
}

func (s KafkaStore) UserTermsOfService() store.UserTermsOfServiceStore {
	return s.userTermsOfService
}

func (s KafkaStore) LinkMetadata() store.LinkMetadataStore {
	return s.linkMetadata
}

func (s KafkaStore) MarkSystemRanUnitTests() {
	s.baseStore.MarkSystemRanUnitTests()
}

func (s KafkaStore) Close() {
	s.conn.Close()
	s.baseStore.Close()
}

func (s KafkaStore) LockToMaster() {
	s.baseStore.LockToMaster()
}

func (s KafkaStore) UnlockFromMaster() {
	s.baseStore.UnlockFromMaster()
}

func (s KafkaStore) DropAllTables() {
	s.baseStore.DropAllTables()
}

func (s KafkaStore) TotalMasterDbConnections() int {
	return s.baseStore.TotalMasterDbConnections()
}

func (s KafkaStore) TotalReadDbConnections() int {
	return s.baseStore.TotalReadDbConnections()
}

func (s KafkaStore) TotalSearchDbConnections() int {
	return s.baseStore.TotalSearchDbConnections()
}

func (s KafkaStore) NotificationRegistry() store.NotificationRegistryStore {
	return s.notificationRegistry
}

func (s KafkaStore) sendMessage(method string, params map[string]interface{}) {
	b, _ := json.Marshal(map[string]interface{}{"method": method, "params": params})
	s.conn.WriteMessages(kafka.Message{Value: b})
}
