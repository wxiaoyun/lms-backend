package bookview

import (
	"lms-backend/internal/model"
)

type SimpleView struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

func ToSimpleView(book *model.Book) *SimpleView {
	return &SimpleView{
		ID:    book.ID,
		Title: book.Title,
	}
}
