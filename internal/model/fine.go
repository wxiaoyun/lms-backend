package model

import (
	"lms-backend/pkg/error/externalerrors"
	"slices"

	"gorm.io/gorm"
)

type FineStatus = string

type Fine struct {
	gorm.Model

	UserID uint       `gorm:"not null"`
	LoanID uint       `gorm:"not null"`
	Status FineStatus `gorm:"not null"`
	Amount float64    `gorm:"not null"`
}

const (
	FineModelName = "fine"
	FineTableName = "fines"
)

const (
	FineStatusOutstanding FineStatus = "outstanding"
	FineStatusPaid        FineStatus = "paid"
)

const (
	OverdueFine = 1000
)

func (f *Fine) Create(db *gorm.DB) error {
	return db.Create(f).Error
}

func (f *Fine) Update(db *gorm.DB) error {
	return db.Updates(f).Error
}

func (f *Fine) Delete(db *gorm.DB) error {
	return db.Delete(f).Error
}

func (f *Fine) ensureUserExists(db *gorm.DB) error {
	var exists int64

	result := db.Model(&User{}).
		Where("id = ?", f.UserID).
		Count(&exists)
	if result.Error != nil {
		return result.Error
	}

	if exists == 0 {
		return externalerrors.BadRequest("User does not exist")
	}

	return nil
}

func (f *Fine) ensureLoanExists(db *gorm.DB) error {
	var exists int64

	result := db.Model(&Loan{}).
		Where("id = ?", f.LoanID).
		Count(&exists)
	if result.Error != nil {
		return result.Error
	}

	if exists == 0 {
		return externalerrors.BadRequest("Loan does not exist")
	}

	return nil
}

func (f *Fine) Validate(db *gorm.DB) error {
	if f.Amount <= 0 {
		return externalerrors.BadRequest("Amount must be greater than 0")
	}

	if !slices.Contains([]FineStatus{
		FineStatusOutstanding,
		FineStatusPaid,
	}, f.Status) {
		return externalerrors.BadRequest("Invalid status")
	}

	if err := f.ensureUserExists(db); err != nil {
		return err
	}

	return f.ensureLoanExists(db)
}

func (f *Fine) BeforeCreate(db *gorm.DB) error {
	return f.Validate(db)
}

func (f *Fine) BeforeUpdate(db *gorm.DB) error {
	return f.Validate(db)
}
