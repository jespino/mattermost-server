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
	store.Store
	metrics       einterfaces.MetricsInterface
	cluster       einterfaces.ClusterInterface
	reaction      LocalCacheReactionStore
	role          LocalCacheRoleStore
	scheme        LocalCacheSchemeStore
	reactionCache *utils.Cache
	roleCache     *utils.Cache
	schemeCache   *utils.Cache
	groupCache    *utils.Cache
}

func NewLocalCacheSupplier(baseStore store.Store, metrics einterfaces.MetricsInterface, cluster einterfaces.ClusterInterface) LocalCacheStore {
	localCacheStore := LocalCacheStore{
		Store:   baseStore,
		cluster: cluster,
		metrics: metrics,
	}
	localCacheStore.reactionCache = utils.NewLruWithParams(REACTION_CACHE_SIZE, "Reaction", REACTION_CACHE_SEC, model.CLUSTER_EVENT_INVALIDATE_CACHE_FOR_REACTIONS)
	localCacheStore.roleCache = utils.NewLruWithParams(ROLE_CACHE_SIZE, "Role", ROLE_CACHE_SEC, model.CLUSTER_EVENT_INVALIDATE_CACHE_FOR_ROLES)
	localCacheStore.schemeCache = utils.NewLruWithParams(SCHEME_CACHE_SIZE, "Scheme", SCHEME_CACHE_SEC, model.CLUSTER_EVENT_INVALIDATE_CACHE_FOR_SCHEMES)
	localCacheStore.groupCache = utils.NewLruWithParams(GROUP_CACHE_SIZE, "Group", GROUP_CACHE_SEC, model.CLUSTER_EVENT_INVALIDATE_CACHE_FOR_GROUPS)
	localCacheStore.reaction = LocalCacheReactionStore{ReactionStore: baseStore.Reaction(), rootStore: &localCacheStore}
	localCacheStore.role = LocalCacheRoleStore{RoleStore: baseStore.Role(), rootStore: &localCacheStore}
	localCacheStore.scheme = LocalCacheSchemeStore{SchemeStore: baseStore.Scheme(), rootStore: &localCacheStore}
	return localCacheStore
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
