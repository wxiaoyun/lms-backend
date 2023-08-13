package model

import (
	"lms-backend/pkg/error/externalerrors"
	"time"

	"gorm.io/gorm"
)

type BookStatus = string
type UserStatus = string

type Book struct {
	gorm.Model

	Title           string    `gorm:"not null"`
	Author          string    `gorm:"not null"`
	ISBN            string    `gorm:"not null"`
	Publisher       string    `gorm:"not null"`
	PublicationDate time.Time `gorm:"not null"`
	Genre           string    `gorm:"not null"`
	Language        string    `gorm:"not null"`
	Loans           []Loan    `gorm:"->"`
}

const (
	BookModelName = "book"
	BookTableName = "books"
)

const (
	BookStatusAvailable   BookStatus = "available"
	BookStatusUnavailable BookStatus = "unavailable"
	BookStatusOnLoan      BookStatus = "on loan"
	BookStatusOnReserve   BookStatus = "on reserve"
)

func (b *Book) Create(db *gorm.DB) error {
	return db.Create(b).Error
}

func (b *Book) Update(db *gorm.DB) error {
	return db.Updates(b).Error
}

func (b *Book) Delete(db *gorm.DB) error {
	return db.Delete(b).Error
}

func (b *Book) Validate(_ *gorm.DB) error {
	if b.Title == "" {
		return externalerrors.BadRequest("title is required")
	}

	if b.Author == "" {
		return externalerrors.BadRequest("author is required")
	}

	if b.ISBN == "" {
		return externalerrors.BadRequest("isbn is required")
	}

	if b.Publisher == "" {
		return externalerrors.BadRequest("publisher is required")
	}

	if (time.Time{}).Equal(b.PublicationDate) {
		return externalerrors.BadRequest("publication date is required")
	}

	if b.Genre == "" {
		return externalerrors.BadRequest("genre is required")
	}

	if b.Language == "" {
		return externalerrors.BadRequest("language is required")
	}

	return nil
}

func (b *Book) BeforeCreate(_ *gorm.DB) error {
	return b.Validate(nil)
}

func (b *Book) BeforeUpdate(_ *gorm.DB) error {
	return b.Validate(nil)
}
