package model

import (
	"database/sql"
	"lms-backend/pkg/error/externalerrors"
	"lms-backend/util/sliceutil"
	"time"

	"gorm.io/gorm"
)

type LoanStatus = string

type Loan struct {
	gorm.Model

	UserID        uint          `gorm:"not null"`
	User          *User         `gorm:"->"`
	BookID        uint          `gorm:"not null"`
	Book          *Book         `gorm:"->"`
	Status        LoanStatus    `gorm:"not null"`
	BorrowDate    time.Time     `gorm:"not null"` // Date when the book is borrowed
	DueDate       time.Time     `gorm:"not null"` // Date when the book is due
	ReturnDate    sql.NullTime  // Date when the book is returned
	LoanHistories []LoanHistory `gorm:"->;<-:create"`
	Fines         []Fine        `gorm:"->;<-:create"`
}

const (
	LoanModelName = "loan"
	LoanTableName = "loans"
)

const (
	LoanStatusBorrowed LoanStatus = "borrowed"
	LoanStatusReturned LoanStatus = "returned"
)

const (
	LoanDuration = 7 * 24 * time.Hour
	MaximumLoans = 5
)

func (l *Loan) Create(db *gorm.DB) error {
	return db.Create(l).Error
}

func (l *Loan) Update(db *gorm.DB) error {
	for _, hist := range l.LoanHistories {
		if hist.ID == 0 {
			if err := hist.Create(db); err != nil {
				return err
			}
		}
	}

	return db.Updates(l).Error
}

// Need to call preloadAssociations	before calling this method.
func (l *Loan) Delete(db *gorm.DB) error {
	for _, hist := range l.LoanHistories {
		if err := hist.Delete(db); err != nil {
			return err
		}
	}

	for _, fine := range l.Fines {
		if err := fine.Delete(db); err != nil {
			return err
		}
	}

	return db.Delete(l).Error
}

func (l *Loan) ensureUserExistsAndPresent(db *gorm.DB) error {
	if l.UserID == 0 {
		return externalerrors.BadRequest("user id is required")
	}

	var exists int64
	result := db.Model(&User{}).Where("id = ?", l.UserID).Count(&exists)
	if err := result.Error; err != nil {
		return err
	}

	if exists == 0 {
		return externalerrors.BadRequest("user does not exist")
	}

	return nil
}

func (l *Loan) ensureBookExistsAndPresent(db *gorm.DB) error {
	if l.BookID == 0 {
		return externalerrors.BadRequest("book id is required")
	}

	var exists int64
	result := db.Model(&Book{}).Where("id = ?", l.BookID).Count(&exists)
	if err := result.Error; err != nil {
		return err
	}

	if exists == 0 {
		return externalerrors.BadRequest("book does not exist")
	}

	return nil
}

func (l *Loan) ValidateStatus() error {
	if l.Status == "" {
		return externalerrors.BadRequest("status is required")
	}

	if !sliceutil.Contains([]LoanStatus{
		LoanStatusBorrowed,
		LoanStatusReturned,
	}, l.Status) {
		return externalerrors.BadRequest("invalid loan status")
	}

	if l.Status == LoanStatusReturned && !l.ReturnDate.Valid {
		return externalerrors.BadRequest("return date is required")
	}

	return nil
}

func (l *Loan) Validate(db *gorm.DB) error {
	if l.BorrowDate.IsZero() {
		return externalerrors.BadRequest("borrow date is required")
	}

	if l.DueDate.IsZero() {
		return externalerrors.BadRequest("due date is required")
	}

	if l.DueDate.Before(l.BorrowDate) {
		return externalerrors.BadRequest("due date must be after borrow date")
	}

	if l.ReturnDate.Valid && l.ReturnDate.Time.Before(l.BorrowDate) {
		return externalerrors.BadRequest("return date must be after borrow date")
	}

	if err := l.ensureUserExistsAndPresent(db); err != nil {
		return err
	}

	if err := l.ensureBookExistsAndPresent(db); err != nil {
		return err
	}

	return l.ValidateStatus()
}

func (l *Loan) BeforeCreate(db *gorm.DB) error {
	return l.Validate(db)
}

func (l *Loan) BeforeUpdate(db *gorm.DB) error {
	return l.Validate(db)
}
