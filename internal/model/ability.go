package model

import (
	"time"

	"gorm.io/gorm"
)

type Ability struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time

	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
}

func (a *Ability) Create(db *gorm.DB) error {
	return db.Create(a).Error
}
