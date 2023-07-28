package worksheethandler

import (
	"technical-test/internal/api"
	"technical-test/internal/dataaccess/worksheet"
	"technical-test/internal/database"
	"technical-test/internal/view/worksheetview"
	collection "technical-test/pkg/collectionquery"

	"github.com/gofiber/fiber/v2"
)

// @Summary List all the existing worksheets
// @Description list all the existing worksheets
// @Tags worksheet
// @Accept */*
// @Produce application/json
// @Success 200 "OK"
// @Router /api/v1/worksheet/ [get]
func HandleList(c *fiber.Ctx) error {
	db := database.GetDB()

	cq := collection.GetCollectionQueryFromParam(c)

	totalCount, err := worksheet.Count(db)
	if err != nil {
		return err
	}

	worksheets, err := worksheet.List(db, cq)
	if err != nil {
		return err
	}

	filteredCount, err := worksheet.CountFiltered(db, cq)
	if err != nil {
		return err
	}

	var view = []*worksheetview.WorkSheetView{}
	for _, w := range worksheets {
		//nolint:gosec // loop does not modify struct
		view = append(view, worksheetview.ToView(&w))
	}

	return c.JSON(api.Response{
		Data: view,
		Meta: api.Meta{
			TotalCount:    totalCount,
			FilteredCount: filteredCount,
		},
		Messages: []api.Message{
			api.SuccessMessage("worksheets listed successfully"),
		},
	})
}
