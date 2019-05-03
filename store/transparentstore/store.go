package transparentstore

import "github.com/mattermost/mattermost-server/store"

type TransparentStore struct {
	baseStore            store.Store
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

func NewTransparentSupplier(baseStore store.Store) TransparentStore {
	return TransparentStore{
		baseStore:            baseStore,
		team:                 TransparentTeamStore{baseStore: baseStore.Team()},
		channel:              TransparentChannelStore{baseStore: baseStore.Channel()},
		post:                 TransparentPostStore{baseStore: baseStore.Post()},
		user:                 TransparentUserStore{baseStore: baseStore.User()},
		bot:                  TransparentBotStore{baseStore: baseStore.Bot()},
		audit:                TransparentAuditStore{baseStore: baseStore.Audit()},
		clusterDiscovery:     TransparentClusterDiscoveryStore{baseStore: baseStore.ClusterDiscovery()},
		compliance:           TransparentComplianceStore{baseStore: baseStore.Compliance()},
		session:              TransparentSessionStore{baseStore: baseStore.Session()},
		oAuth:                TransparentOAuthStore{baseStore: baseStore.OAuth()},
		system:               TransparentSystemStore{baseStore: baseStore.System()},
		webhook:              TransparentWebhookStore{baseStore: baseStore.Webhook()},
		command:              TransparentCommandStore{baseStore: baseStore.Command()},
		commandWebhook:       TransparentCommandWebhookStore{baseStore: baseStore.CommandWebhook()},
		preference:           TransparentPreferenceStore{baseStore: baseStore.Preference()},
		license:              TransparentLicenseStore{baseStore: baseStore.License()},
		token:                TransparentTokenStore{baseStore: baseStore.Token()},
		emoji:                TransparentEmojiStore{baseStore: baseStore.Emoji()},
		status:               TransparentStatusStore{baseStore: baseStore.Status()},
		fileInfo:             TransparentFileInfoStore{baseStore: baseStore.FileInfo()},
		reaction:             TransparentReactionStore{baseStore: baseStore.Reaction()},
		role:                 TransparentRoleStore{baseStore: baseStore.Role()},
		scheme:               TransparentSchemeStore{baseStore: baseStore.Scheme()},
		job:                  TransparentJobStore{baseStore: baseStore.Job()},
		userAccessToken:      TransparentUserAccessTokenStore{baseStore: baseStore.UserAccessToken()},
		channelMemberHistory: TransparentChannelMemberHistoryStore{baseStore: baseStore.ChannelMemberHistory()},
		plugin:               TransparentPluginStore{baseStore: baseStore.Plugin()},
		termsOfService:       TransparentTermsOfServiceStore{baseStore: baseStore.TermsOfService()},
		group:                TransparentGroupStore{baseStore: baseStore.Group()},
		userTermsOfService:   TransparentUserTermsOfServiceStore{baseStore: baseStore.UserTermsOfService()},
		linkMetadata:         TransparentLinkMetadataStore{baseStore: baseStore.LinkMetadata()},
		notificationRegistry: TransparentNotificationRegistryStore{baseStore: baseStore.NotificationRegistry()},
	}
}

func (s TransparentStore) Team() store.TeamStore {
	return s.team
}

func (s TransparentStore) Channel() store.ChannelStore {
	return s.channel
}

func (s TransparentStore) Post() store.PostStore {
	return s.post
}

func (s TransparentStore) User() store.UserStore {
	return s.user
}

func (s TransparentStore) Bot() store.BotStore {
	return s.bot
}

func (s TransparentStore) Audit() store.AuditStore {
	return s.audit
}

func (s TransparentStore) ClusterDiscovery() store.ClusterDiscoveryStore {
	return s.clusterDiscovery
}

func (s TransparentStore) Compliance() store.ComplianceStore {
	return s.compliance
}

func (s TransparentStore) Session() store.SessionStore {
	return s.session
}

func (s TransparentStore) OAuth() store.OAuthStore {
	return s.oAuth
}

func (s TransparentStore) System() store.SystemStore {
	return s.system
}

func (s TransparentStore) Webhook() store.WebhookStore {
	return s.webhook
}

func (s TransparentStore) Command() store.CommandStore {
	return s.command
}

func (s TransparentStore) CommandWebhook() store.CommandWebhookStore {
	return s.commandWebhook
}

func (s TransparentStore) Preference() store.PreferenceStore {
	return s.preference
}

func (s TransparentStore) License() store.LicenseStore {
	return s.license
}

func (s TransparentStore) Token() store.TokenStore {
	return s.token
}

func (s TransparentStore) Emoji() store.EmojiStore {
	return s.emoji
}

func (s TransparentStore) Status() store.StatusStore {
	return s.status
}

func (s TransparentStore) FileInfo() store.FileInfoStore {
	return s.fileInfo
}

func (s TransparentStore) Reaction() store.ReactionStore {
	return s.reaction
}

func (s TransparentStore) Role() store.RoleStore {
	return s.role
}

func (s TransparentStore) Scheme() store.SchemeStore {
	return s.scheme
}

func (s TransparentStore) Job() store.JobStore {
	return s.job
}

func (s TransparentStore) UserAccessToken() store.UserAccessTokenStore {
	return s.userAccessToken
}

func (s TransparentStore) ChannelMemberHistory() store.ChannelMemberHistoryStore {
	return s.channelMemberHistory
}

func (s TransparentStore) Plugin() store.PluginStore {
	return s.plugin
}

func (s TransparentStore) TermsOfService() store.TermsOfServiceStore {
	return s.termsOfService
}

func (s TransparentStore) Group() store.GroupStore {
	return s.group
}

func (s TransparentStore) UserTermsOfService() store.UserTermsOfServiceStore {
	return s.userTermsOfService
}

func (s TransparentStore) LinkMetadata() store.LinkMetadataStore {
	return s.linkMetadata
}

func (s TransparentStore) MarkSystemRanUnitTests() {
	s.baseStore.MarkSystemRanUnitTests()
}

func (s TransparentStore) Close() {
	s.baseStore.Close()
}

func (s TransparentStore) LockToMaster() {
	s.baseStore.LockToMaster()
}

func (s TransparentStore) UnlockFromMaster() {
	s.baseStore.UnlockFromMaster()
}

func (s TransparentStore) DropAllTables() {
	s.baseStore.DropAllTables()
}

func (s TransparentStore) TotalMasterDbConnections() int {
	return s.baseStore.TotalMasterDbConnections()
}

func (s TransparentStore) TotalReadDbConnections() int {
	return s.baseStore.TotalReadDbConnections()
}

func (s TransparentStore) TotalSearchDbConnections() int {
	return s.baseStore.TotalSearchDbConnections()
}

func (s TransparentStore) NotificationRegistry() store.NotificationRegistryStore {
	return s.notificationRegistry
}
