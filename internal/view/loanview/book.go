package loanview

import (
	"lms-backend/internal/model"
	"lms-backend/internal/view/sharedview"
)

type WithBookView struct {
	sharedview.LoanView
	Book *sharedview.BookView `json:"book"`
}

func ToBookView(loan *model.Loan) *WithBookView {
	return &WithBookView{
		LoanView: *sharedview.ToLoanView(loan),
		Book:     sharedview.ToBookView(loan.BookCopy.Book),
	}
}
