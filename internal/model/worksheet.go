package model

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Worksheet struct {
	gorm.Model

	Title       string `gorm:"not null"`
	UserID      uint   `gorm:"not null"`
	User        *User
	Cost        float64 `gorm:"not null;check:cost > 0"`
	Price       float64 `gorm:"not null;check:price > 0"`
	Description string  `gorm:"not null"`

	Questions []Question
}

func (w *Worksheet) ensureTitleIsUnique(db *gorm.DB) error {
	var exists int64

	result := db.Model(&Worksheet{}).
		Where("title = ?", w.Title).
		Count(&exists)
	if result.Error != nil {
		return result.Error
	}

	if exists > 0 {
		return fiber.NewError(fiber.StatusBadRequest, "title already exists")
	}

	return nil
}

func (w *Worksheet) Validate(db *gorm.DB) error {
	if w.Title == "" {
		return fiber.NewError(fiber.StatusBadRequest, "title is required")
	}

	if err := w.ensureTitleIsUnique(db); err != nil {
		return err
	}

	if w.Cost <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "cost is required and positive")
	}

	if w.Price <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "price is required and positive")
	}

	if w.Description == "" {
		return fiber.NewError(fiber.StatusBadRequest, "description is required")
	}

	return nil
}

func (w *Worksheet) Create(db *gorm.DB) error {
	return db.Create(w).Error
}

func (w *Worksheet) Update(db *gorm.DB) error {
	return db.Updates(w).Error
}

func (w *Worksheet) Delete(db *gorm.DB) error {
	return db.Delete(w).Error
}

func (w *Worksheet) BeforeCreate(db *gorm.DB) error {
	return w.Validate(db)
}

func (w *Worksheet) BeforeUpdate(db *gorm.DB) error {
	return w.Validate(db)
}

// Assumes questions is properly preloaded
func (w *Worksheet) GetTotalCost() float64 {
	var cost float64
	for _, q := range w.Questions {
		cost += q.GetCost()
	}
	return cost + w.Cost
}

func (w *Worksheet) GetTotalPrice() float64 {
	return w.Price
}

func (w *Worksheet) GetTotalProfit() float64 {
	return w.GetTotalPrice() - w.GetTotalCost()
}

func (w *Worksheet) IsPositiveProfit() bool {
	return w.GetTotalProfit() > 0
}

func (w *Worksheet) IsNegativeProfit() bool {
	return w.GetTotalProfit() < 0
}
