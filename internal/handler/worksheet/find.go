package worksheethandler

import (
	"fmt"
	"lms-backend/internal/api"
	"lms-backend/internal/dataaccess/worksheet"
	"lms-backend/internal/database"
	"lms-backend/internal/view/worksheetview"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// @Summary Find first few worksheets with matching title to the search query
// @Description  find first few worksheets with matching title to the search query
// @Tags worksheet
// @Accept */*
// @Produce application/json
// @Success 200 "OK"
// @Router /api/v1/worksheet/find [get]
func HandleFind(c *fiber.Ctx) error {
	search := c.Query("search")
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Response{
			Messages: []api.Message{api.ErrorMessage(
				fmt.Sprintf("%s is not a valid limit.", c.Query("limit"))),
			},
			Error: err.Error(),
		})
	}

	db := database.GetDB()
	ws, err := worksheet.Find(db, search, int(limit))
	if err != nil {
		return err
	}

	views := make([]worksheetview.FindView, len(ws))
	for i, w := range ws {
		//nolint:gosec // struct is not modified
		views[i] = *worksheetview.ToFindView(&w)
	}

	return c.JSON(api.Response{
		Data: views,
		Messages: []api.Message{api.SilentMessage(
			fmt.Sprintf("Worksheets matching %s retrieved.", search)),
		},
	})
}
