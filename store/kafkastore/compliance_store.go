package kafkastore

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
)

type KafkaComplianceStore struct {
	baseStore store.ComplianceStore
	root      *KafkaStore
}

func (s KafkaComplianceStore) Save(compliance *model.Compliance) (*model.Compliance, *model.AppError) {
	return s.baseStore.Save(compliance)
}

func (s KafkaComplianceStore) Update(compliance *model.Compliance) (*model.Compliance, *model.AppError) {
	return s.baseStore.Update(compliance)
}

func (s KafkaComplianceStore) Get(id string) (*model.Compliance, *model.AppError) {
	return s.baseStore.Get(id)
}

func (s KafkaComplianceStore) GetAll(offset int, limit int) (model.Compliances, *model.AppError) {
	return s.baseStore.GetAll(offset, limit)
}

func (s KafkaComplianceStore) ComplianceExport(compliance *model.Compliance) ([]*model.CompliancePost, *model.AppError) {
	return s.baseStore.ComplianceExport(compliance)
}

func (s KafkaComplianceStore) MessageExport(after int64, limit int) ([]*model.MessageExport, *model.AppError) {
	return s.baseStore.MessageExport(after, limit)
}
