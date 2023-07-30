package worksheetview

import (
	"technical-test/internal/model"
)

type FindView struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

func ToFindView(workSheets *model.Worksheet) *FindView {
	return &FindView{
		ID:    workSheets.ID,
		Title: workSheets.Title,
	}
}
