package sharedview

import (
	"lms-backend/internal/model"

	"time"

	"github.com/ForAeons/ternary"
)

type LoanView struct {
	ID         int64      `json:"id,omitempty"`
	UserID     int64      `json:"user_id"`
	BookCopyID int64      `json:"book_copy_id"`
	Status     string     `json:"status"`
	BorrowDate *time.Time `json:"borrow_date"`
	DueDate    *time.Time `json:"due_date"`
	ReturnDate *time.Time `json:"return_date"`
}

func ToLoanView(loan *model.Loan) *LoanView {
	return &LoanView{
		ID:         int64(loan.ID),
		UserID:     int64(loan.UserID),
		BookCopyID: int64(loan.BookCopyID),
		Status:     loan.Status,
		BorrowDate: &loan.BorrowDate,
		DueDate:    &loan.DueDate,
		ReturnDate: ternary.If[*time.Time](loan.ReturnDate.Valid).
			Then(&loan.ReturnDate.Time).
			Else(nil),
	}
}
