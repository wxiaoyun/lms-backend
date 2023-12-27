package bookmarkview

import (
	"lms-backend/internal/model"
	"lms-backend/internal/view/bookview"
)

type DetailedView struct {
	BaseView
	Book *bookview.BaseView `json:"book"`
}

func ToDetailedView(bookmark *model.Bookmark) *DetailedView {
	return &DetailedView{
		BaseView: *ToView(bookmark),
		Book:     bookview.ToView(bookmark.Book),
	}
}
