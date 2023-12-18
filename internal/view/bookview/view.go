package bookview

import (
	"lms-backend/internal/model"
	"lms-backend/internal/view/loanview"
	"lms-backend/internal/view/reservationview"
	"time"
)

type BaseView struct {
	ID              uint   `json:"id,omitempty"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	ISBN            string `json:"isbn"`
	Publisher       string `json:"publisher"`
	PublicationDate string `json:"publication_date"`
	Genre           string `json:"genre"`
	Language        string `json:"language"`
}

func ToView(book *model.Book) *BaseView {
	return &BaseView{
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

type DetailedView struct {
	BaseView
	Loans        []loanview.View        `json:"loans"`
	Reservations []reservationview.View `json:"reservations"`
}

func ToDetailedView(book *model.Book) *DetailedView {
	loanviews := make([]loanview.View, len(book.Loans))
	for i, loan := range book.Loans {
		//nolint:gosec // loop does not modify struct
		loanviews[i] = *loanview.ToView(&loan)
	}

	reservationviews := make([]reservationview.View, len(book.Reservations))
	for i, reservation := range book.Reservations {
		//nolint:gosec // loop does not modify struct
		reservationviews[i] = *reservationview.ToView(&reservation)
	}

	return &DetailedView{
		BaseView:     *ToView(book),
		Loans:        loanviews,
		Reservations: reservationviews,
	}
}
