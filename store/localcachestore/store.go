package localcachestore

import (
	"github.com/mattermost/mattermost-server/einterfaces"
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
	"github.com/mattermost/mattermost-server/utils"
)

const (
	REACTION_CACHE_SIZE = 20000
	REACTION_CACHE_SEC  = 30 * 60

	ROLE_CACHE_SIZE = 20000
	ROLE_CACHE_SEC  = 30 * 60

	SCHEME_CACHE_SIZE = 20000
	SCHEME_CACHE_SEC  = 30 * 60

	GROUP_CACHE_SIZE = 20000
	GROUP_CACHE_SEC  = 30 * 60

	CLEAR_CACHE_MESSAGE_DATA = ""
)

type LocalCacheStore struct {
	baseStore            store.Store
	metrics              einterfaces.MetricsInterface
	cluster              einterfaces.ClusterInterface
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
	reactionCache        *utils.Cache
	role                 store.RoleStore
	roleCache            *utils.Cache
	scheme               store.SchemeStore
	schemeCache          *utils.Cache
	job                  store.JobStore
	userAccessToken      store.UserAccessTokenStore
	channelMemberHistory store.ChannelMemberHistoryStore
	plugin               store.PluginStore
	termsOfService       store.TermsOfServiceStore
	group                store.GroupStore
	groupCache           *utils.Cache
	userTermsOfService   store.UserTermsOfServiceStore
	linkMetadata         store.LinkMetadataStore
	notificationRegistry store.NotificationRegistryStore
}

func NewLocalCacheSupplier(baseStore store.Store, metrics einterfaces.MetricsInterface, cluster einterfaces.ClusterInterface) LocalCacheStore {
	localCacheStore := LocalCacheStore{
		baseStore:            baseStore,
		cluster:              cluster,
		metrics:              metrics,
		team:                 baseStore.Team(),
		channel:              baseStore.Channel(),
		post:                 baseStore.Post(),
		user:                 baseStore.User(),
		bot:                  baseStore.Bot(),
		audit:                baseStore.Audit(),
		clusterDiscovery:     baseStore.ClusterDiscovery(),
		compliance:           baseStore.Compliance(),
		session:              baseStore.Session(),
		oAuth:                baseStore.OAuth(),
		system:               baseStore.System(),
		webhook:              baseStore.Webhook(),
		command:              baseStore.Command(),
		commandWebhook:       baseStore.CommandWebhook(),
		preference:           baseStore.Preference(),
		license:              baseStore.License(),
		token:                baseStore.Token(),
		emoji:                baseStore.Emoji(),
		status:               baseStore.Status(),
		fileInfo:             baseStore.FileInfo(),
		job:                  baseStore.Job(),
		userAccessToken:      baseStore.UserAccessToken(),
		channelMemberHistory: baseStore.ChannelMemberHistory(),
		plugin:               baseStore.Plugin(),
		termsOfService:       baseStore.TermsOfService(),
		userTermsOfService:   baseStore.UserTermsOfService(),
		linkMetadata:         baseStore.LinkMetadata(),
		notificationRegistry: baseStore.NotificationRegistry(),
	}
	localCacheStore.reactionCache = utils.NewLruWithParams(REACTION_CACHE_SIZE, "Reaction", REACTION_CACHE_SEC, model.CLUSTER_EVENT_INVALIDATE_CACHE_FOR_REACTIONS)
	localCacheStore.roleCache = utils.NewLruWithParams(ROLE_CACHE_SIZE, "Role", ROLE_CACHE_SEC, model.CLUSTER_EVENT_INVALIDATE_CACHE_FOR_ROLES)
	localCacheStore.schemeCache = utils.NewLruWithParams(SCHEME_CACHE_SIZE, "Scheme", SCHEME_CACHE_SEC, model.CLUSTER_EVENT_INVALIDATE_CACHE_FOR_SCHEMES)
	localCacheStore.groupCache = utils.NewLruWithParams(GROUP_CACHE_SIZE, "Group", GROUP_CACHE_SEC, model.CLUSTER_EVENT_INVALIDATE_CACHE_FOR_GROUPS)
	localCacheStore.reaction = LocalCacheReactionStore{baseStore: baseStore.Reaction(), rootStore: &localCacheStore}
	localCacheStore.role = LocalCacheRoleStore{baseStore: baseStore.Role(), rootStore: &localCacheStore}
	localCacheStore.scheme = LocalCacheSchemeStore{baseStore: baseStore.Scheme(), rootStore: &localCacheStore}
	localCacheStore.group = LocalCacheGroupStore{baseStore: baseStore.Group(), rootStore: &localCacheStore}
	return localCacheStore
}

func (s LocalCacheStore) Team() store.TeamStore {
	return s.team
}

func (s LocalCacheStore) Channel() store.ChannelStore {
	return s.channel
}

func (s LocalCacheStore) Post() store.PostStore {
	return s.post
}

func (s LocalCacheStore) User() store.UserStore {
	return s.user
}

func (s LocalCacheStore) Bot() store.BotStore {
	return s.bot
}

func (s LocalCacheStore) Audit() store.AuditStore {
	return s.audit
}

func (s LocalCacheStore) ClusterDiscovery() store.ClusterDiscoveryStore {
	return s.clusterDiscovery
}

func (s LocalCacheStore) Compliance() store.ComplianceStore {
	return s.compliance
}

func (s LocalCacheStore) Session() store.SessionStore {
	return s.session
}

func (s LocalCacheStore) OAuth() store.OAuthStore {
	return s.oAuth
}

func (s LocalCacheStore) System() store.SystemStore {
	return s.system
}

func (s LocalCacheStore) Webhook() store.WebhookStore {
	return s.webhook
}

func (s LocalCacheStore) Command() store.CommandStore {
	return s.command
}

func (s LocalCacheStore) CommandWebhook() store.CommandWebhookStore {
	return s.commandWebhook
}

func (s LocalCacheStore) Preference() store.PreferenceStore {
	return s.preference
}

func (s LocalCacheStore) License() store.LicenseStore {
	return s.license
}

func (s LocalCacheStore) Token() store.TokenStore {
	return s.token
}

func (s LocalCacheStore) Emoji() store.EmojiStore {
	return s.emoji
}

func (s LocalCacheStore) Status() store.StatusStore {
	return s.status
}

func (s LocalCacheStore) FileInfo() store.FileInfoStore {
	return s.fileInfo
}

func (s LocalCacheStore) Reaction() store.ReactionStore {
	return s.reaction
}

func (s LocalCacheStore) Role() store.RoleStore {
	return s.role
}

func (s LocalCacheStore) Scheme() store.SchemeStore {
	return s.scheme
}

func (s LocalCacheStore) Job() store.JobStore {
	return s.job
}

func (s LocalCacheStore) UserAccessToken() store.UserAccessTokenStore {
	return s.userAccessToken
}

func (s LocalCacheStore) ChannelMemberHistory() store.ChannelMemberHistoryStore {
	return s.channelMemberHistory
}

func (s LocalCacheStore) Plugin() store.PluginStore {
	return s.plugin
}

func (s LocalCacheStore) TermsOfService() store.TermsOfServiceStore {
	return s.termsOfService
}

func (s LocalCacheStore) Group() store.GroupStore {
	return s.group
}

func (s LocalCacheStore) UserTermsOfService() store.UserTermsOfServiceStore {
	return s.userTermsOfService
}

func (s LocalCacheStore) LinkMetadata() store.LinkMetadataStore {
	return s.linkMetadata
}

func (s LocalCacheStore) MarkSystemRanUnitTests() {
	s.baseStore.MarkSystemRanUnitTests()
}

func (s LocalCacheStore) Close() {
	s.baseStore.Close()
}

func (s LocalCacheStore) LockToMaster() {
	s.baseStore.LockToMaster()
}

func (s LocalCacheStore) UnlockFromMaster() {
	s.baseStore.UnlockFromMaster()
}

func (s LocalCacheStore) DropAllTables() {
	s.baseStore.DropAllTables()
}

func (s LocalCacheStore) TotalMasterDbConnections() int {
	return s.baseStore.TotalMasterDbConnections()
}

func (s LocalCacheStore) TotalReadDbConnections() int {
	return s.baseStore.TotalReadDbConnections()
}

func (s LocalCacheStore) TotalSearchDbConnections() int {
	return s.baseStore.TotalSearchDbConnections()
}

func (s LocalCacheStore) NotificationRegistry() store.NotificationRegistryStore {
	return s.notificationRegistry
}

func (s *LocalCacheStore) doInvalidateCacheCluster(cache *utils.Cache, key string) {
	cache.Remove(key)
	if s.cluster != nil {
		msg := &model.ClusterMessage{
			Event:    cache.GetInvalidateClusterEvent(),
			SendType: model.CLUSTER_SEND_BEST_EFFORT,
			Data:     key,
		}
		s.cluster.SendClusterMessage(msg)
	}
}

func (s *LocalCacheStore) doStandardAddToCache(cache *utils.Cache, key string, value interface{}) {
	cache.AddWithDefaultExpires(key, value)
}

func (s *LocalCacheStore) doStandardReadCache(cache *utils.Cache, key string) interface{} {
	if cacheItem, ok := cache.Get(key); ok {
		if s.metrics != nil {
			s.metrics.IncrementMemCacheHitCounter(cache.Name())
		}
		return cacheItem
	}

	if s.metrics != nil {
		s.metrics.IncrementMemCacheMissCounter(cache.Name())
	}

	return nil
}

func (s *LocalCacheStore) doClearCacheCluster(cache *utils.Cache) {
	cache.Purge()
	if s.cluster != nil {
		msg := &model.ClusterMessage{
			Event:    cache.GetInvalidateClusterEvent(),
			SendType: model.CLUSTER_SEND_BEST_EFFORT,
			Data:     CLEAR_CACHE_MESSAGE_DATA,
		}
		s.cluster.SendClusterMessage(msg)
	}
}

func (s *LocalCacheStore) Invalidate() {
	s.doClearCacheCluster(s.reactionCache)
	s.doClearCacheCluster(s.roleCache)
	s.doClearCacheCluster(s.schemeCache)
	s.doClearCacheCluster(s.groupCache)
}
