package reservationview

import (
	"lms-backend/internal/model"
	"lms-backend/internal/view/userview"
	"time"
)

type DetailedView struct {
	BaseView
	User userview.View `json:"user"`
	Book BookView      `json:"book"`
}

type BookView struct {
	ID              uint   `json:"id,omitempty"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	ISBN            string `json:"isbn"`
	Publisher       string `json:"publisher"`
	PublicationDate string `json:"publication_date"`
	Genre           string `json:"genre"`
	Language        string `json:"language"`
}

func ToDetailedView(res *model.Reservation) *DetailedView {
	return &DetailedView{
		BaseView: *ToView(res),
		User:     *userview.ToView(res.User),
		Book: BookView{
			ID:              res.Book.ID,
			Title:           res.Book.Title,
			Author:          res.Book.Author,
			ISBN:            res.Book.ISBN,
			Publisher:       res.Book.Publisher,
			PublicationDate: res.Book.PublicationDate.Format(time.RFC3339),
			Genre:           res.Book.Genre,
			Language:        res.Book.Language,
		},
	}
}
