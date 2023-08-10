package worksheethandler

import (
	"fmt"
	"lms-backend/internal/api"
	"lms-backend/internal/dataaccess/worksheet"
	"lms-backend/internal/database"
	"lms-backend/internal/view/worksheetview"

	"github.com/gofiber/fiber/v2"
)

// @Summary Summarize all the existing worksheets
// @Description summarizes important information about all the existing worksheets
// @Tags worksheet
// @Accept */*
// @Produce application/json
// @Success 200 "OK"
// @Router /api/v1/worksheet/summary [get]
func HandleWorksheetSummary(c *fiber.Ctx) error {
	db := database.GetDB()
	summary, err := worksheet.Summarize(db)
	if err != nil {
		return err
	}

	totalCount, err := worksheet.Count(db)
	if err != nil {
		return err
	}

	view := worksheetview.ToSummaryView(summary)

	return c.JSON(api.Response{
		Data: view,
		Meta: api.Meta{
			TotalCount: totalCount,
		},
		Messages: []api.Message{api.SuccessMessage(
			fmt.Sprintf("Total of %d worksheets summarized successfully.", totalCount)),
		},
	})
}
