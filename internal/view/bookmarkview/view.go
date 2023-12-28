package bookmarkview

import (
	"lms-backend/internal/model"
	"lms-backend/internal/view/sharedview"
)

type BaseView struct {
	sharedview.BookmarkView
}

func ToView(bookmark *model.Bookmark) *BaseView {
	return &BaseView{
		BookmarkView: *sharedview.ToBookmarkView(bookmark),
	}
}
