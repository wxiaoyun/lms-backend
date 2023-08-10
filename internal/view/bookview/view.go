package bookview

import (
	"lms-backend/internal/model"
	"time"
)

type View struct {
	ID              uint   `json:"id,omitempty"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	ISBN            string `json:"isbn"`
	Publisher       string `json:"publisher"`
	PublicationDate string `json:"publication_date"`
	Genre           string `json:"genre"`
	Language        string `json:"language"`
}

func ToView(book *model.Book) *View {
	return &View{
		ID:              book.ID,
		Title:           book.Title,
		Author:          book.Author,
		ISBN:            book.ISBN,
		Publisher:       book.Publisher,
		PublicationDate: book.PublicationDate.Format(time.RFC3339),
		Genre:           book.Genre,
		Language:        book.Language,
	}
}
