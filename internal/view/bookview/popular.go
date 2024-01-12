package bookview

import (
	"fmt"
	"lms-backend/internal/config"
	"lms-backend/internal/model"
	"lms-backend/internal/viewmodel"
)

type PopularView struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	LoanCount    int64  `json:"loan_count"`
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
}

func ToPopularView(b *viewmodel.PopularBookViewModel) *PopularView {
	thumbnailURL := ""
	if b.ThumbnailFilename != "" {
		thumbnailURL = fmt.Sprintf(model.ImageDownloadURL, config.BackendURL, b.ThumbnailFilename)
	}
	return &PopularView{
		ID:           b.ID,
		Title:        b.Title,
		LoanCount:    b.LoanCount,
		ThumbnailURL: thumbnailURL,
	}
}
