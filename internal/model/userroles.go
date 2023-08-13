package model

import (
	"time"
)

type UserRoles struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time

	UserID uint `gorm:"not null"`
	RoleID uint `gorm:"not null"`
}
