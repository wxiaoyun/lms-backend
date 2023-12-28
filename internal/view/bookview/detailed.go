package bookview

import (
	"lms-backend/internal/model"
	"lms-backend/internal/view/sharedview"
)

type DetailedView struct {
	View
	BookCopies []sharedview.BookCopyView `json:"book_copies"`
}

func ToDetailedView(b *model.Book) *DetailedView {
	copies := []sharedview.BookCopyView{}
	for _, copy := range b.BookCopies {
		//nolint:gosec // loop does not modify struct
		copies = append(copies, *sharedview.ToBookCopyView(&copy))
	}

	return &DetailedView{
		View:       *ToView(b),
		BookCopies: copies,
	}
}
