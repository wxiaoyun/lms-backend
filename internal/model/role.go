package model

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	Name      string    `gorm:"not null"`
	Abilities []Ability `gorm:"many2many:role_abilities;->;<-:create"`
}

func (r *Role) Create(db *gorm.DB) error {
	return db.Create(r).Error
}
