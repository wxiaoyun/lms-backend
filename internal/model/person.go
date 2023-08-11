package model

import (
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model

	FullName           string `gorm:"not null"`
	PreferredName      string
	LanguagePreference string `gorm:"not null"`
}

const (
	PersonTableName = "people"
)

const (
	MaximumNameLength = 255
	MinimumNameLength = 2
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

func (p *Person) ValidateName() error {
	if p.FullName == "" {
		return fiber.NewError(fiber.StatusBadRequest, "last name is required")
	}

	if utf8.RuneCountInString(p.FullName) > MaximumNameLength {
		return fiber.NewError(fiber.StatusBadRequest, "fullname is too long")
	}

	if utf8.RuneCountInString(p.FullName) < MinimumNameLength {
		return fiber.NewError(fiber.StatusBadRequest, "fullname is too short")
	}

	return nil
}

func (p *Person) BeforeCreate(_ *gorm.DB) error {
	return p.ValidateName()
}

func (p *Person) BeforeUpdate(_ *gorm.DB) error {
	return p.ValidateName()
}
