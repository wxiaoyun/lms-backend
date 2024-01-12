package bookview

import (
	"lms-backend/internal/model"
	"lms-backend/internal/view/sharedview"
)

type View struct {
	sharedview.BookView
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
}

func ToView(book *model.Book) *View {
	thumbnailurl := ""
	if book.Thumbnail != nil {
		thumbnailurl = book.Thumbnail.GetImageDownloadURL()
	}

	return &View{
		BookView:     *sharedview.ToBookView(book),
		ThumbnailURL: thumbnailurl,
	}
}
