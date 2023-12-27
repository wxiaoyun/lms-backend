package bookmarkview

import (
	"lms-backend/internal/model"
)

type BaseView struct {
	ID     uint `json:"id,omitempty"`
	UserID uint `json:"user_id"`
	BookID uint `json:"book_id"`
}

func ToView(bookmark *model.Bookmark) *BaseView {
	return &BaseView{
		ID:     bookmark.ID,
		UserID: bookmark.UserID,
		BookID: bookmark.BookID,
	}
}
