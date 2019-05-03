package transparentstore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type TransparentComplianceStore struct {
	baseStore store.ComplianceStore
}

func (s TransparentComplianceStore) Save(compliance *model.Compliance) (*model.Compliance, *model.AppError) {
	return s.baseStore.Save(compliance)
}

func (s TransparentComplianceStore) Update(compliance *model.Compliance) (*model.Compliance, *model.AppError) {
	return s.baseStore.Update(compliance)
}

func (s TransparentComplianceStore) Get(id string) (*model.Compliance, *model.AppError) {
	return s.baseStore.Get(id)
}

func (s TransparentComplianceStore) GetAll(offset int, limit int) (model.Compliances, *model.AppError) {
	return s.baseStore.GetAll(offset, limit)
}

func (s TransparentComplianceStore) ComplianceExport(compliance *model.Compliance) ([]*model.CompliancePost, *model.AppError) {
	return s.baseStore.ComplianceExport(compliance)
}

func (s TransparentComplianceStore) MessageExport(after int64, limit int) ([]*model.MessageExport, *model.AppError) {
	return s.baseStore.MessageExport(after, limit)
}
