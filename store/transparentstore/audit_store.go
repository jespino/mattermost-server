package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentAuditStore struct {
	baseStore store.AuditStore
}

func (s TransparentAuditStore) Save(audit *model.Audit) *model.AppError {
	return s.baseStore.Save(audit)
}

func (s TransparentAuditStore) Get(user_id string, offset int, limit int) (model.Audits, *model.AppError) {
	return s.baseStore.Get(user_id, offset, limit)
}

func (s TransparentAuditStore) PermanentDeleteByUser(userId string) *model.AppError {
	return s.baseStore.PermanentDeleteByUser(userId)
}

func (s TransparentAuditStore) PermanentDeleteBatch(endTime int64, limit int64) (int64, *model.AppError) {
	return s.baseStore.PermanentDeleteBatch(endTime, limit)
}
