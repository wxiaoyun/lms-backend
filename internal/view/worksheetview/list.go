package worksheetview

import (
	"technical-test/internal/model"
)

type WorkSheetListView struct {
	ID          uint    `json:"id,omitempty"`
	Title       string  `json:"title"`
	UserID      uint    `json:"user_id"`
	Cost        float64 `json:"cost"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

func ToListView(workSheet *model.Worksheet) *WorkSheetListView {
	return &WorkSheetListView{
		ID:          workSheet.ID,
		Title:       workSheet.Title,
		UserID:      workSheet.UserID,
		Cost:        workSheet.Cost,
		Price:       workSheet.Price,
		Description: workSheet.Description,
	}
}
