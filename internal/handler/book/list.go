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

func HandleList(c *fiber.Ctx) error {
	err := policy.Authorize(c, readBookAction, bookpolicy.ListPolicy())
	if err != nil {
		return err
	}

	cq := collection.GetCollectionQueryFromParam(c)
	db := database.GetDB()

	totalCount, err := book.Count(db)
	if err != nil {
		return err
	}

	dbFiltered := cq.Filter(db, book.Filters())

	filteredCount, err := book.Count(dbFiltered)
	if err != nil {
		return err
	}

	dbSorted := cq.Sort(dbFiltered, book.Sorters())
	dbPaginated := cq.Paginate(dbSorted)
	books, err := book.ListDetailed(dbPaginated)
	if err != nil {
		return err
	}

	var view = []bookview.DetailedView{}
	for _, w := range books {
		//nolint:gosec // loop does not modify struct
		view = append(view, *bookview.ToDetailedView(&w))
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
