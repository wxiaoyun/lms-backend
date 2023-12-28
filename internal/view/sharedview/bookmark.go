package sharedview

import (
	"lms-backend/internal/model"
)

type BookmarkView struct {
	ID     uint `json:"id,omitempty"`
	UserID uint `json:"user_id"`
	BookID uint `json:"book_id"`
}

func ToBookmarkView(bookmark *model.Bookmark) *BookmarkView {
	return &BookmarkView{
		ID:     bookmark.ID,
		UserID: bookmark.UserID,
		BookID: bookmark.BookID,
	}
}
