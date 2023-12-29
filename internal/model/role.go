package model

import (
	"time"

	"gorm.io/gorm"
)

const (
	AdminRole    = 1
	LibAdminRole = 2
	StaffRole    = 3
	MemberRole   = 4
)

type Role struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time

	Name      string    `gorm:"not null"`
	Abilities []Ability `gorm:"many2many:role_abilities;->"`
}

func (r *Role) Create(db *gorm.DB) error {
	return db.Create(r).Error
}
