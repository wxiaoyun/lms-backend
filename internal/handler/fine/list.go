package finehandler

import (
	"lms-backend/internal/api"
	"lms-backend/internal/dataaccess/fine"
	"lms-backend/internal/database"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/finepolicy"
	"lms-backend/internal/view/fineview"
	collection "lms-backend/pkg/collectionquery"

	"github.com/gofiber/fiber/v2"
)

const (
	readFineAction = "read fine"
)

// @Summary List fines
// @Description List fines belonging to a loan
// @Tags fine
// @Accept application/json
// @Param offset query int false "Offset for pagination"
// @Param limit query int false "Limit for pagination"
// @Param filter[user_id] query int false "Filter by user ID"
// @Param sortBy query string false "Sort by column name"
// @Param orderBy query string false "Order by direction (asc or desc)"
// @Produce application/json
// @Success 200 {object} api.SwgResponse[[]fineview.View]
// @Failure 400 {object} api.SwgErrResponse
// @Router /v1/fine [get]
func HandleList(c *fiber.Ctx) error {
	err := policy.Authorize(c, readFineAction, finepolicy.ReadPolicy())
	if err != nil {
		return err
	}

	cq := collection.GetCollectionQueryFromParam(c)
	db := database.GetDB()

	totalCount, err := fine.Count(db)
	if err != nil {
		return err
	}

	dbFiltered := cq.Filter(db, fine.Filters())

	filteredCount, err := fine.Count(dbFiltered)
	if err != nil {
		return err
	}

	dbSorted := cq.Sort(dbFiltered, fine.Sorters())
	dbPaginated := cq.Paginate(dbSorted)

	fns, err := fine.List(dbPaginated)
	if err != nil {
		return err
	}

	var view = []*fineview.View{}
	for _, w := range fns {
		//nolint:gosec // loop does not modify struct
		view = append(view, fineview.ToView(&w))
	}

	return c.JSON(api.Response{
		Data: view,
		Meta: api.Meta{
			TotalCount:    totalCount,
			FilteredCount: filteredCount,
		},
		Messages: api.Messages(
			api.SilentMessage("fines listed successfully"),
		),
	})
}
