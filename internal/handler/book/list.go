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
// @Description Lists books in the library
// @Tags book
// @Accept application/json
// @Param offset query int false "Offset for pagination"
// @Param limit query int false "Limit for pagination"
// @Param filter[value] query string false "Filter by value"
// @Param sortBy query string false "Sort by column name (e.g. title)"
// @Param orderBy query string false "Order by direction (asc or desc)"
// @Produce application/json
// @Success 200 {object} api.SwgResponse[[]bookview.BaseView]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/book [get]
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

	dbFiltered := cq.Filter(db, book.Filters())

	filteredCount, err := book.Count(dbFiltered)
	if err != nil {
		return err
	}

	dbSorted := cq.Sort(dbFiltered, book.Sorters())
	dbPaginated := cq.Paginate(dbSorted)
	books, err := book.List(dbPaginated)
	if err != nil {
		return err
	}

	var view = []*bookview.BaseView{}
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
