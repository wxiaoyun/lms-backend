package model

import (
	"time"
)

type RoleAbilities struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time

	RoleID    uint `gorm:"not null"`
	AbilityID uint `gorm:"not null"`
}
