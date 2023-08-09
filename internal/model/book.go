package model

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model

	Title           string    `gorm:"not null"`
	Author          string    `gorm:"not null"`
	ISBN            string    `gorm:"not null"`
	Publisher       string    `gorm:"not null"`
	PublicationDate time.Time `gorm:"not null"`
	Genre           string    `gorm:"not null"`
	Language        string    `gorm:"not null"`
}

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
		return fiber.NewError(fiber.StatusBadRequest, "title is required")
	}

	if b.Author == "" {
		return fiber.NewError(fiber.StatusBadRequest, "author is required")
	}

	if b.ISBN == "" {
		return fiber.NewError(fiber.StatusBadRequest, "isbn is required")
	}

	if b.Publisher == "" {
		return fiber.NewError(fiber.StatusBadRequest, "publisher is required")
	}

	if (time.Time{}).Equal(b.PublicationDate) {
		return fiber.NewError(fiber.StatusBadRequest, "publication date is required")
	}

	if b.Genre == "" {
		return fiber.NewError(fiber.StatusBadRequest, "genre is required")
	}

	if b.Language == "" {
		return fiber.NewError(fiber.StatusBadRequest, "language is required")
	}

	return nil
}

func (b *Book) BeforeCreate(_ *gorm.DB) error {
	return b.Validate(nil)
}

func (b *Book) BeforeUpdate(_ *gorm.DB) error {
	return b.Validate(nil)
}
