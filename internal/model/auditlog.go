package model

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuditLog struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UserID    uint   `gorm:"not null"`
	Action    string `gorm:"not null"`
}

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
		return fiber.NewError(fiber.StatusNotFound, "user not found")
	}

	return nil
}

func (a *AuditLog) Validate(db *gorm.DB) error {
	if a.Action == "" {
		return fiber.NewError(fiber.StatusBadRequest, "action is required")
	}

	return a.ensureUserExists(db)
}

func (a *AuditLog) BeforeCreate(db *gorm.DB) error {
	return a.Validate(db)
}
