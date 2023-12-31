package model

import (
	"lms-backend/pkg/error/externalerrors"
	"time"

	"gorm.io/gorm"
)

type AuditLog struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time

	UserID uint      `gorm:"not null"`
	User   *User     `gorm:"->"`
	Action string    `gorm:"not null"`
	Date   time.Time `gorm:"not null"`
}

const (
	AuditLogModelName = "audit_log"
	AuditLogTableName = "audit_logs"
)

func (a *AuditLog) Create(db *gorm.DB) error {
	return db.Create(a).Error
}

func (a *AuditLog) ensureUserExists(db *gorm.DB) error {
	var exists int64

	result := db.Model(&User{}).Where("id = ?", a.UserID).Count(&exists)
	if result.Error != nil {
		return result.Error
	}

	if exists == 0 {
		return externalerrors.BadRequest("user not found")
	}

	return nil
}

func (a *AuditLog) Validate(db *gorm.DB) error {
	if a.Action == "" {
		return externalerrors.BadRequest("action is required")
	}

	return a.ensureUserExists(db)
}

func (a *AuditLog) BeforeCreate(db *gorm.DB) error {
	if (time.Time{}).Equal(a.Date) {
		a.Date = time.Now()
	}

	return a.Validate(db)
}
