package rediscachestore

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/go-redis/redis"
	"github.com/mattermost/mattermost-server/einterfaces"
	"github.com/mattermost/mattermost-server/mlog"
	"github.com/mattermost/mattermost-server/store"
)

const REDIS_EXPIRY_TIME = 30 * time.Minute

type RedisCacheStore struct {
	store.Store
	metrics  einterfaces.MetricsInterface
	cluster  einterfaces.ClusterInterface
	reaction RedisCacheReactionStore
	role     RedisCacheRoleStore
	client   *redis.Client
}

func NewRedisCacheLayer(baseStore store.Store, metrics einterfaces.MetricsInterface, cluster einterfaces.ClusterInterface) RedisCacheStore {
	redisCacheStore := RedisCacheStore{
		Store:   baseStore,
		cluster: cluster,
		metrics: metrics,
	}
	redisCacheStore.client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if _, err := redisCacheStore.client.Ping().Result(); err != nil {
		mlog.Error("Unable to ping redis server: " + err.Error())
	}
	redisCacheStore.reaction = RedisCacheReactionStore{ReactionStore: baseStore.Reaction(), client: redisCacheStore.client}
	redisCacheStore.role = RedisCacheRoleStore{RoleStore: baseStore.Role(), client: redisCacheStore.client}
	return redisCacheStore
}

func (s RedisCacheStore) Reaction() store.ReactionStore {
	return s.reaction
}

func (s RedisCacheStore) Role() store.RoleStore {
	return s.role
}

func (s *RedisCacheStore) save(key string, value interface{}, expiry time.Duration) error {
	if bytes, err := GetBytes(value); err != nil {
		return err
	} else {
		if err := s.client.Set(key, bytes, expiry).Err(); err != nil {
			return err
		}
	}
	return nil
}

func (s *RedisCacheStore) load(key string, writeTo interface{}) (bool, error) {
	if data, err := s.client.Get(key).Bytes(); err != nil {
		if err == redis.Nil {
			return false, nil
		} else {
			return false, err
		}
	} else {
		if err := DecodeBytes(data, writeTo); err != nil {
			return false, err
		}
	}
	return true, nil
}

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func DecodeBytes(input []byte, thing interface{}) error {
	dec := gob.NewDecoder(bytes.NewReader(input))
	return dec.Decode(thing)
}
