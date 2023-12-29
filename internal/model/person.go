package model

import (
	"lms-backend/pkg/error/externalerrors"
	"unicode/utf8"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model

	FullName      string `gorm:"not null"`
	PreferredName string
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
		return externalerrors.BadRequest("last name is required")
	}

	if utf8.RuneCountInString(p.FullName) > MaximumNameLength {
		return externalerrors.BadRequest("fullname is too long")
	}

	if utf8.RuneCountInString(p.FullName) < MinimumNameLength {
		return externalerrors.BadRequest("fullname is too short")
	}

	return nil
}

func (p *Person) Validate() error {
	return p.ValidateName()
}

func (p *Person) BeforeCreate(_ *gorm.DB) error {
	return p.Validate()
}

func (p *Person) BeforeUpdate(_ *gorm.DB) error {
	return p.Validate()
}
