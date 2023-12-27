package bookcopyview

import (
	"lms-backend/internal/model"
	"lms-backend/internal/view/sharedview"
)

type DetailedView struct {
	sharedview.BookCopyView
	Book *sharedview.BookView `json:"book"`
}

func ToDetailedView(bookCopy *model.BookCopy) *DetailedView {
	return &DetailedView{
		BookCopyView: *sharedview.ToBookCopyView(bookCopy),
		Book:         sharedview.ToBookView(bookCopy.Book),
	}
}
