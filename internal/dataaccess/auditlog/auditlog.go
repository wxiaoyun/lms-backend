package auditlog

import (
	"lms-backend/internal/model"
	collection "lms-backend/pkg/collectionquery"

	"gorm.io/gorm"
)

func Create(db *gorm.DB, auditLog *model.AuditLog) (*model.AuditLog, error) {
	if err := auditLog.Create(db); err != nil {
		return nil, err
	}

	return auditLog, nil
}

func List(db *gorm.DB, cq *collection.Query) ([]model.AuditLog, error) {
	var logs []model.AuditLog

	result := db.Model(&model.AuditLog{}).
		Where("action ILIKE ?", "%"+cq.Search+"%").
		Offset(cq.Offset).
		Limit(cq.Limit).
		Find(&logs)
	if result.Error != nil {
		return nil, result.Error
	}

	return logs, nil
}

func Count(db *gorm.DB) (int64, error) {
	var count int64

	result := db.Model(&model.AuditLog{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func CountFiltered(db *gorm.DB, cq *collection.Query) (int64, error) {
	var count int64

	result := db.Model(&model.AuditLog{}).
		Where("action ILIKE ?", "%"+cq.Search+"%").
		Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}
