package worksheethandler

import (
	"technical-test/internal/api"
	"technical-test/internal/dataaccess/worksheet"
	"technical-test/internal/database"
	"technical-test/internal/view/worksheetview"

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

	worksheets, err := worksheet.List(db)
	if err != nil {
		return err
	}

	var view = []*worksheetview.WorkSheetListView{}
	for _, w := range worksheets {
		//nolint:gosec // loop does not modify struct
		view = append(view, worksheetview.ToListView(&w))
	}

	return c.JSON(api.Response{
		Data: view,
		Messages: []api.Message{
			api.SuccessMessage("worksheets listed successfully"),
		},
	})
}
