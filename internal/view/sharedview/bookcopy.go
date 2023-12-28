package sharedview

import (
	"lms-backend/internal/model"
)

type BookCopyView struct {
	ID     uint   `json:"id,omitempty"`
	BookID uint   `json:"book_id"`
	Status string `json:"status"`
}

func ToBookCopyView(bookCopy *model.BookCopy) *BookCopyView {
	return &BookCopyView{
		ID:     bookCopy.ID,
		BookID: bookCopy.BookID,
		Status: bookCopy.Status,
	}
}
