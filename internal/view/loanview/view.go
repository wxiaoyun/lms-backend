package loanview

import (
	"lms-backend/internal/model"
	"lms-backend/internal/view/sharedview"
)

type DetailedView struct {
	sharedview.LoanView
	User *sharedview.UserView `json:"user"`
	Book *sharedview.BookView `json:"book"`
}

func ToDetailedView(loan *model.Loan) *DetailedView {
	return &DetailedView{
		LoanView: *sharedview.ToLoanView(loan),
		User:     sharedview.ToUserView(loan.User),
		Book:     sharedview.ToBookView(loan.BookCopy.Book),
	}
}
