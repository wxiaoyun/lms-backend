package worksheethandler

import (
	"fmt"
	"strconv"
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
// @Router /api/v1/worksheet/:id [get]
func HandleRead(c *fiber.Ctx) error {
	param := c.Params("id")
	worksheetID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Response{
			Messages: []string{fmt.Sprintf("%s is not a valid worksheet id.", param)},
			Error:    err.Error(),
		})
	}

	db := database.GetDB()
	ws, err := worksheet.Read(db, worksheetID)
	if err != nil {
		return err
	}

	view := worksheetview.ToView(ws)

	return c.JSON(api.Response{
		Data:     view,
		Messages: []string{fmt.Sprintf("Worksheet %s retrieved successfully.", ws.Title)},
	})
}
