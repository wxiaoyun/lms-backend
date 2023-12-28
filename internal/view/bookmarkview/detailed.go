package bookmarkview

import (
	"lms-backend/internal/model"
	"lms-backend/internal/view/sharedview"
)

type DetailedView struct {
	BaseView
	Book *sharedview.BookView `json:"book"`
}

func ToDetailedView(bookmark *model.Bookmark) *DetailedView {
	return &DetailedView{
		BaseView: *ToView(bookmark),
		Book:     sharedview.ToBookView(bookmark.Book),
	}
}
