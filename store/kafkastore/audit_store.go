package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaAuditStore struct {
	baseStore store.AuditStore
	root      *KafkaStore
}

func (s KafkaAuditStore) Save(audit *model.Audit) *model.AppError {
	s.root.sendMessage("Audit.Save", map[string]interface{}{"audit": audit})
	return s.baseStore.Save(audit)
}

func (s KafkaAuditStore) Get(user_id string, offset int, limit int) (model.Audits, *model.AppError) {
	return s.baseStore.Get(user_id, offset, limit)
}

func (s KafkaAuditStore) PermanentDeleteByUser(userId string) *model.AppError {
	return s.baseStore.PermanentDeleteByUser(userId)
}

func (s KafkaAuditStore) PermanentDeleteBatch(endTime int64, limit int64) (int64, *model.AppError) {
	return s.baseStore.PermanentDeleteBatch(endTime, limit)
}
