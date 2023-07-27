package model

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model

	Description string  `gorm:"not null"`
	Answer      string  `gorm:"not null"`
	Cost        float64 `gorm:"not null;check:cost > 0"`

	WorksheetID uint `gorm:"not null"`
	Worksheet   *Worksheet
}

func (q *Question) Create(db *gorm.DB) error {
	return db.Create(q).Error
}

func (q *Question) Update(db *gorm.DB) error {
	return db.Updates(q).Error
}

func (q *Question) Delete(db *gorm.DB) error {
	return db.Delete(q).Error
}

func (q *Question) ensureWorksheetExists(db *gorm.DB) error {
	var exists int64

	result := db.Model(&Worksheet{}).
		Where("id = ?", q.WorksheetID).
		Count(&exists)
	if result.Error != nil {
		return result.Error
	}

	if exists == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "worksheet does not exist")
	}

	return nil
}

func (q *Question) Validate(db *gorm.DB) error {
	if q.Description == "" {
		return fiber.NewError(fiber.StatusBadRequest, "description is required")
	}

	if q.Answer == "" {
		return fiber.NewError(fiber.StatusBadRequest, "answer is required")
	}

	if q.Cost <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "cost is required and positive")
	}

	if q.WorksheetID <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "worksheet id is required")
	}

	return q.ensureWorksheetExists(db)
}

func (q *Question) BeforeCreate(db *gorm.DB) error {
	return q.Validate(db)
}

func (q *Question) BeforeUpdate(db *gorm.DB) error {
	return q.Validate(db)
}
