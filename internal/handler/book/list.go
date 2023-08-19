package bookhandler

import (
	"lms-backend/internal/api"
	"lms-backend/internal/dataaccess/book"
	"lms-backend/internal/database"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/bookpolicy"
	"lms-backend/internal/view/bookview"
	collection "lms-backend/pkg/collectionquery"

	"github.com/gofiber/fiber/v2"
)

// @Summary List books
// @Description lists books in the library
// @Tags book
// @Accept */*
// @Produce application/json
// @Success 200 {object} api.SwgResponse[[]bookview.View]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/book/ [get]
func HandleList(c *fiber.Ctx) error {
	err := policy.Authorize(c, readBookAction, bookpolicy.ReadPolicy())
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
