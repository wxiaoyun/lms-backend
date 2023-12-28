package model

import (
	"fmt"
	"lms-backend/pkg/error/externalerrors"
	"lms-backend/util/sliceutil"

	"gorm.io/gorm"
)

type BookStatus = string

type BookCopy struct {
	gorm.Model

	BookID       uint          `gorm:"not null"`
	Book         *Book         `gorm:"->"`
	Status       BookStatus    `gorm:"not null"`
	Loans        []Loan        `gorm:"->"`
	Reservations []Reservation `gorm:"->"`
}

const (
	BookCopyModelName = "bookcopy"
	BookCopyTableName = "bookcopies"
)

const (
	BookStatusAvailable BookStatus = "available"
	BookStatusOnLoan    BookStatus = "loaned"
	BookStatusOnReserve BookStatus = "reserved"
)

func (b *BookCopy) Create(db *gorm.DB) error {
	return db.Create(b).Error
}

// Loans and reservation should not be updated/created here.
func (b *BookCopy) Update(db *gorm.DB) error {
	return db.Updates(b).Error
}

// All loans associated with this book will be deleted.
//
// Need to call preloadAssociations	before calling this method.
func (b *BookCopy) Delete(db *gorm.DB) error {
	for _, loan := range b.Loans {
		if err := loan.Delete(db); err != nil {
			return err
		}
	}

	for _, reserve := range b.Reservations {
		if err := reserve.Delete(db); err != nil {
			return err
		}
	}

	return db.Delete(b).Error
}

func (b *BookCopy) ensureBookExistOrNew(db *gorm.DB) error {
	if b.BookID == 0 {
		return nil
	}

	var exists int64
	result := db.Model(&Book{}).
		Where("id = ?", b.BookID).
		Count(&exists)
	if result.Error != nil {
		return result.Error
	}

	if exists == 0 {
		return externalerrors.BadRequest(fmt.Sprintf("book with id %d does not exist", b.BookID))
	}

	return nil
}

func (b *BookCopy) ensureCopyIsNotOnLoan(db *gorm.DB) error {
	if b.Status != BookStatusOnLoan {
		return nil
	}

	var exists int64
	result := db.Model(&Loan{}).
		Where("book_copy_id = ?", b.ID).
		Where("status = ?", LoanStatusBorrowed).
		Count(&exists)
	if result.Error != nil {
		return result.Error
	}

	if exists > 0 {
		return externalerrors.BadRequest(fmt.Sprintf("book copy with id %d is on loan", b.ID))
	}

	return nil
}

func (b *BookCopy) ensureCopyIsNotOnReserve(db *gorm.DB) error {
	if b.Status != BookStatusOnReserve {
		return nil
	}

	var exists int64
	result := db.Model(&Reservation{}).
		Where("book_copy_id = ?", b.ID).
		Where("status = ?", ReservationStatusPending).
		Count(&exists)
	if result.Error != nil {
		return result.Error
	}

	if exists > 0 {
		return externalerrors.BadRequest(fmt.Sprintf("book copy with id %d is on reserve", b.ID))
	}

	return nil
}

func (b *BookCopy) ValidateStatus() error {
	if b.Status == "" {
		return externalerrors.BadRequest("status is required")
	}

	if !sliceutil.Contains([]BookStatus{
		BookStatusAvailable,
		BookStatusOnLoan,
		BookStatusOnReserve,
	}, b.Status) {
		return externalerrors.BadRequest("invalid status")
	}

	return nil
}

func (b *BookCopy) Validate(db *gorm.DB) error {
	if err := b.ensureBookExistOrNew(db); err != nil {
		return err
	}

	return b.ValidateStatus()
}

func (b *BookCopy) BeforeCreate(db *gorm.DB) error {
	if b.Status == "" {
		b.Status = BookStatusAvailable
	}

	return b.Validate(db)
}

func (b *BookCopy) BeforeUpdate(db *gorm.DB) error {
	return b.Validate(db)
}

func (b *BookCopy) BeforeDelete(db *gorm.DB) error {
	if err := b.ensureCopyIsNotOnLoan(db); err != nil {
		return err
	}

	return b.ensureCopyIsNotOnReserve(db)
}
