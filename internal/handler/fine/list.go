package finehandler

import (
	"lms-backend/internal/api"
	"lms-backend/internal/dataaccess/book"
	"lms-backend/internal/database"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/finepolicy"
	"lms-backend/internal/view/bookview"
	collection "lms-backend/pkg/collectionquery"

	"github.com/gofiber/fiber/v2"
)

const (
	readFineAction = "read fine"
)

// @Summary List fine
// @Description List fines belonging to a loan
// @Tags fine
// @Accept */*
// @Produce application/json
// @Success 200 {object} api.SwgResponse[[]fineview.View]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/book/{book_id}/loan/{loan_id}/fine/ [get]
func HandleList(c *fiber.Ctx) error {
	err := policy.Authorize(c, readFineAction, finepolicy.ReadPolicy())
	if err != nil {
		return err
	}

	cq := collection.GetCollectionQueryFromParam(c)
	db := database.GetDB()

	totalCount, err := book.Count(db)
	if err != nil {
		return err
	}

	filteredCount, err := book.CountFiltered(db, cq)
	if err != nil {
		return err
	}

	books, err := book.List(db, cq)
	if err != nil {
		return err
	}

	var view = []*bookview.View{}
	for _, w := range books {
		//nolint:gosec // loop does not modify struct
		view = append(view, bookview.ToView(&w))
	}

	return c.JSON(api.Response{
		Data: view,
		Meta: api.Meta{
			TotalCount:    totalCount,
			FilteredCount: filteredCount,
		},
		Messages: api.Messages(
			api.SilentMessage("books listed successfully"),
		),
	})
}
