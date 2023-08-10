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

// @Summary List all the existing worksheets
// @Description list all the existing worksheets
// @Tags worksheet
// @Accept */*
// @Produce application/json
// @Success 200 "OK"
// @Router /api/v1/worksheet/:id [get]
func HandleRead(c *fiber.Ctx) error {
	param := c.Params("id")
	worksheetID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Response{
			Messages: []api.Message{api.ErrorMessage(
				fmt.Sprintf("%s is not a valid worksheet id.", param)),
			},
			Error: err.Error(),
		})
	}

	db := database.GetDB()
	ws, err := worksheet.Read(db, worksheetID)
	if err != nil {
		return err
	}

	view := worksheetview.ToView(ws)

	return c.JSON(api.Response{
		Data: view,
		Messages: []api.Message{api.SuccessMessage(
			fmt.Sprintf("Worksheet %s retrieved successfully.", ws.Title)),
		},
	})
}
