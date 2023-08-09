package model

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model

	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
}

const (
	PersonTableName = "people"
)

func (Person) TableName() string {
	return PersonTableName
}

func (p *Person) Create(db *gorm.DB) error {
	return db.Create(p).Error
}

func (p *Person) Update(db *gorm.DB) error {
	return db.Updates(p).Error
}

func (p *Person) Delete(db *gorm.DB) error {
	return db.Delete(p).Error
}

func (p *Person) Validate(_ *gorm.DB) error {
	if p.FirstName == "" {
		return fiber.NewError(fiber.StatusBadRequest, "first name is required")
	}

	if p.LastName == "" {
		return fiber.NewError(fiber.StatusBadRequest, "last name is required")
	}

	return nil
}

func (p *Person) BeforeCreate(_ *gorm.DB) error {
	return p.Validate(nil)
}

func (p *Person) BeforeUpdate(_ *gorm.DB) error {
	return p.Validate(nil)
}
