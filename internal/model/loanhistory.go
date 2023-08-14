package model

import (
	"slices"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type LoanHistoryAction = string

type LoanHistory struct {
	gorm.Model

	LoanID uint              `gorm:"not null"`
	Action LoanHistoryAction `gorm:"not null"`
}

const (
	LoanHistoryModelName = "loan_history"
	LoanHistoryTableName = "loan_histories"
)

const (
	LoanHistoryActionBorrow       LoanHistoryAction = "borrow"
	LoanHistoryActionReturn       LoanHistoryAction = "return"
	LoanHistoryActionExtend       LoanHistoryAction = "extend"
	LoanHistoryActionIntervention LoanHistoryAction = "intervention"
)

func (l *LoanHistory) Create(db *gorm.DB) error {
	return db.Create(l).Error
}

func (l *LoanHistory) Delete(db *gorm.DB) error {
	return db.Delete(l).Error
}

func (l *LoanHistory) ensureLoanExistsAndPresent(db *gorm.DB) error {
	if l.LoanID == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "loan id is required")
	}

	var exists int64

	result := db.Model(&Loan{}).Where("id = ?", l.LoanID).Count(&exists)
	if result.Error != nil {
		return result.Error
	}

	if exists == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "loan does not exist")
	}

	return nil
}

func (l *LoanHistory) ValidateAction() error {
	if l.Action == "" {
		return fiber.NewError(fiber.StatusBadRequest, "action is required")
	}

	if !slices.Contains([]LoanHistoryAction{
		LoanHistoryActionBorrow,
		LoanHistoryActionExtend,
		LoanHistoryActionReturn,
		LoanHistoryActionIntervention,
	}, l.Action) {
		return fiber.NewError(fiber.StatusBadRequest, "invalid action")
	}

	return nil
}

func (l *LoanHistory) Validate(db *gorm.DB) error {
	if err := l.ensureLoanExistsAndPresent(db); err != nil {
		return err
	}

	return l.ValidateAction()
}

func (l *LoanHistory) BeforeCreate(db *gorm.DB) error {
	return l.Validate(db)
}

func (l *LoanHistory) BeforeUpdate(db *gorm.DB) error {
	return l.Validate(db)
}
