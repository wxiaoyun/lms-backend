package bookview

import (
	"lms-backend/internal/view/googlebookview"
	"strings"
	"time"
)

func GoogleResponseToView(res *googlebookview.ResponseView) []View {
	views := []View{}

	for _, item := range res.Items {
		var bookview View

		bookview.Title = item.VolumeInfo.Title
		bookview.Author = strings.Join(item.VolumeInfo.Authors, ", ")
		bookview.Publisher = item.VolumeInfo.Publisher

		layout := "2006-01-02"
		// nolint:errcheck,gosec // safe to ignore error since we trust google api
		t, _ := time.Parse(layout, item.VolumeInfo.PublishedDate)
		bookview.PublicationDate = t.Format(time.RFC3339)
		bookview.Genre = strings.Join(item.VolumeInfo.Categories, ", ")
		bookview.Language = item.VolumeInfo.Language

		for _, industryIdentifier := range item.VolumeInfo.IndustryIdentifiers {
			if industryIdentifier.Type == "ISBN_13" {
				bookview.ISBN = industryIdentifier.Identifier
				break
			}

			if industryIdentifier.Type == "ISBN_10" {
				bookview.ISBN = industryIdentifier.Identifier
			}
		}

		views = append(views, bookview)
	}

	return views
}
