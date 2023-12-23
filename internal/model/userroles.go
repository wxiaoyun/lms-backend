package model

import (
	"time"
)

type UserRole struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time

	UserID uint `gorm:"not null"`
	RoleID uint `gorm:"not null"`
}
