package loanview

import (
	"lms-backend/internal/model"
	fineview "lms-backend/internal/view/fine"
	loanhistoryview "lms-backend/internal/view/loanhistory"
	"lms-backend/util/ternary"
	"time"
)

type View struct {
	ID            int64                  `json:"id,omitempty"`
	UserID        int64                  `json:"user_id"`
	BookID        int64                  `json:"book_id"`
	Status        string                 `json:"status"`
	BorrowDate    *time.Time             `json:"borrow_date"`
	DueDate       *time.Time             `json:"due_date"`
	ReturnDate    *time.Time             `json:"return_date"`
	LoanHistories []loanhistoryview.View `json:"loan_histories"`
	Fines         []fineview.View        `json:"fines"`
}

func ToView(loan *model.Loan) *View {
	return &View{
		ID:         int64(loan.ID),
		UserID:     int64(loan.UserID),
		BookID:     int64(loan.BookID),
		Status:     loan.Status,
		BorrowDate: &loan.BorrowDate,
		DueDate:    &loan.DueDate,
		ReturnDate: ternary.If[*time.Time](loan.ReturnDate.Valid).
			Then(&loan.ReturnDate.Time).
			Else(nil),
		LoanHistories: loanhistoryview.ToViews(loan.LoanHistories),
		Fines:         fineview.ToViews(loan.Fines),
	}
}
