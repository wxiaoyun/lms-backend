package fineview

import (
	"lms-backend/internal/model"
	"lms-backend/internal/view/sharedview"
)

type DetailedView struct {
	BaseView
	Book *sharedview.BookView `json:"book"`
	Loan *sharedview.LoanView `json:"loan"`
	User *sharedview.UserView `json:"user"`
}

func ToDetailedView(fine *model.Fine) *DetailedView {
	return &DetailedView{
		BaseView: *ToBaseView(fine),
		Book:     sharedview.ToBookView(fine.Loan.BookCopy.Book),
		Loan:     sharedview.ToLoanView(fine.Loan),
		User:     sharedview.ToUserView(fine.User),
	}
}
