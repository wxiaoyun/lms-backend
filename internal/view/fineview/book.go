package fineview

import (
	"lms-backend/internal/model"
	"lms-backend/internal/view/sharedview"
)

type WithBookView struct {
	BaseView
	Book *sharedview.BookView `json:"book"`
}

func ToBookView(fine *model.Fine) *WithBookView {
	return &WithBookView{
		BaseView: *ToBaseView(fine),
		Book:     sharedview.ToBookView(fine.Loan.BookCopy.Book),
	}
}
