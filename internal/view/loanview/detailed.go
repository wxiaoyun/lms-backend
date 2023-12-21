package loanview

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

func ToDetailedView(loan *model.Loan) *DetailedView {
	return &DetailedView{
		BaseView: *ToView(loan),
		User:     *userview.ToView(loan.User),
		Book: BookView{
			ID:              loan.Book.ID,
			Title:           loan.Book.Title,
			Author:          loan.Book.Author,
			ISBN:            loan.Book.ISBN,
			Publisher:       loan.Book.Publisher,
			PublicationDate: loan.Book.PublicationDate.Format(time.RFC3339),
			Genre:           loan.Book.Genre,
			Language:        loan.Book.Language,
		},
	}
}
