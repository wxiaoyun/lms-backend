package auditlog

import (
	"lms-backend/internal/model"
	"lms-backend/internal/orm"

	"gorm.io/gorm"
)

func preloadAssociations(db *gorm.DB) *gorm.DB {
	return db.
		Preload("User").
		Preload("User.Person")
}

func ReadDetailed(db *gorm.DB, id int64) (*model.AuditLog, error) {
	var log model.AuditLog
	result := db.Model(&model.AuditLog{}).
		Scopes(preloadAssociations).
		Where("id = ?", id).
		First(&log)
	if err := result.Error; err != nil {
		if orm.IsRecordNotFound(err) {
			return nil, orm.ErrRecordNotFound(model.AuditLogTableName)
		}
		return nil, err
	}

	return &log, nil
}

func Create(db *gorm.DB, auditLog *model.AuditLog) (*model.AuditLog, error) {
	if err := auditLog.Create(db); err != nil {
		return nil, err
	}

	return ReadDetailed(db, int64(auditLog.ID))
}

func List(db *gorm.DB) ([]model.AuditLog, error) {
	var logs []model.AuditLog

	result := db.Model(&model.AuditLog{}).
		Find(&logs)
	if result.Error != nil {
		return nil, result.Error
	}

	return logs, nil
}

func ListDetailed(db *gorm.DB) ([]model.AuditLog, error) {
	var logs []model.AuditLog

	result := db.Model(&model.AuditLog{}).
		Scopes(preloadAssociations).
		Find(&logs)
	if result.Error != nil {
		return nil, result.Error
	}

	return logs, nil
}

func Count(db *gorm.DB) (int64, error) {
	var count int64

	result := orm.CloneSession(db).
		Model(&model.AuditLog{}).
		Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}
