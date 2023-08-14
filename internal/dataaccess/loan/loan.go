package loan

import (
	"database/sql"
	"lms-backend/internal/model"
	"lms-backend/internal/orm"
	"lms-backend/pkg/error/externalerrors"
	"time"

	"gorm.io/gorm"
)

func preloadAssociations(db *gorm.DB) *gorm.DB {
	return db.
		Preload("LoanHistories").
		Preload("Fines")
}

func Read(db *gorm.DB, loanID int64) (*model.Loan, error) {
	var loan model.Loan

	result := db.Model(&model.Loan{}).
		Where("id = ?", loanID).
		First(&loan)

	if err := result.Error; err != nil {
		if orm.IsRecordNotFound(err) {
			return nil, orm.ErrRecordNotFound(model.LoanModelName)
		}
		return nil, result.Error
	}

	return &loan, nil
}

func ReadDetailed(db *gorm.DB, loanID int64) (*model.Loan, error) {
	var loan model.Loan

	result := db.Model(&model.Loan{}).
		Scopes(preloadAssociations).
		Where("id = ?", loanID).
		First(&loan)

	if err := result.Error; err != nil {
		if orm.IsRecordNotFound(err) {
			return nil, orm.ErrRecordNotFound(model.LoanModelName)
		}
		return nil, result.Error
	}

	return &loan, nil
}

func Delete(db *gorm.DB, loanID int64) (*model.Loan, error) {
	ln, err := ReadDetailed(db, loanID)
	if err != nil {
		return nil, err
	}

	if err := ln.Delete(db); err != nil {
		return nil, err
	}

	return ln, nil
}

// Returns the outstanding loan for the given book, sorted by create date.
func ReadOutstandingLoansByBookID(db *gorm.DB, bookID int64) ([]model.Loan, error) {
	var loans []model.Loan

	result := db.Model(&model.Loan{}).
		Where("book_id = ?", bookID).
		Where("status = ?", model.LoanStatusBorrowed).
		Where("return_date IS NULL").
		Order("created_at DESC").
		Find(&loans)

	if result.Error != nil {
		return nil, result.Error
	}

	return loans, nil
}

// Returns the outstanding loan by the given user, sorted by create date.
func ReadOutstandingLoansByUserID(db *gorm.DB, userID int64) ([]model.Loan, error) {
	var loans []model.Loan

	result := db.Model(&model.Loan{}).
		Where("user_id = ?", userID).
		Where("status = ?", model.LoanStatusBorrowed).
		Where("return_date IS NULL").
		Order("created_at DESC").
		Find(&loans)

	if result.Error != nil {
		return nil, result.Error
	}

	return loans, nil
}

// Returns the overdue loan for the given book, sorted by create date.
func ReadOverdueLoansByBookID(db *gorm.DB, bookID int64) ([]model.Loan, error) {
	var loans []model.Loan

	result := db.Model(&model.Loan{}).
		Where("book_id = ?", bookID).
		Where("status = ?", model.LoanStatusBorrowed).
		Where("return_date IS NULL").
		Where("due_date < NOW()").
		Order("created_at DESC").
		Find(&loans)

	if result.Error != nil {
		return nil, result.Error
	}

	return loans, nil
}

// Returns the overdue loan by the given user, sorted by create date.
func ReadOverdueLoansByUserID(db *gorm.DB, userID int64) ([]model.Loan, error) {
	var loans []model.Loan

	result := db.Model(&model.Loan{}).
		Where("user_id = ?", userID).
		Where("status = ?", model.LoanStatusBorrowed).
		Where("return_date IS NULL").
		Where("due_date < NOW()").
		Order("created_at DESC").
		Find(&loans)

	if result.Error != nil {
		return nil, result.Error
	}

	return loans, nil
}

// Assumes that the book is available.
//
// Relevant checks should be done before calling this function.
//
// User should not have more than maximum reservations and loans.
//
// Book should be neither on loan nor on reserve.
func LoanBook(db *gorm.DB, userID, bookID int64) (*model.Loan, error) {
	ln := model.Loan{
		UserID:     uint(userID),
		BookID:     uint(bookID),
		Status:     model.LoanStatusBorrowed,
		BorrowDate: time.Now(),
		DueDate:    time.Now().Add(model.LoanDuration),
		LoanHistories: []model.LoanHistory{
			{
				Action: model.LoanHistoryActionBorrow,
			},
		},
	}
	if err := ln.Create(db); err != nil {
		return nil, err
	}

	return &ln, nil
}

func ReturnBook(db *gorm.DB, loanID int64) (*model.Loan, error) {
	ln, err := ReadDetailed(db, loanID)
	if err != nil {
		return nil, err
	}

	if ln.Status != model.LoanStatusBorrowed {
		return nil, externalerrors.BadRequest("book is not on loan")
	}

	ln.ReturnDate = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	ln.Status = model.LoanStatusReturned
	ln.LoanHistories = append(ln.LoanHistories, model.LoanHistory{
		LoanID: ln.ID,
		Action: model.LoanHistoryActionReturn,
	})

	if err := ln.Update(db); err != nil {
		return nil, err
	}

	return ln, nil
}

func RenewLoan(db *gorm.DB, loanID int64) (*model.Loan, error) {
	ln, err := ReadDetailed(db, loanID)
	if err != nil {
		return nil, err
	}

	if ln.Status != model.LoanStatusBorrowed {
		return nil, externalerrors.BadRequest("book is not on loan")
	}

	ln.DueDate = ln.DueDate.Add(model.LoanDuration)
	ln.LoanHistories = append(ln.LoanHistories, model.LoanHistory{
		LoanID: ln.ID,
		Action: model.LoanHistoryActionExtend,
	})

	if err := ln.Update(db); err != nil {
		return nil, err
	}

	return ln, nil
}
