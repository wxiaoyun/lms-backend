package loanhandler

import (
	"lms-backend/internal/api"
	"lms-backend/internal/dataaccess/loan"
	"lms-backend/internal/database"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/loanpolicy"
	"lms-backend/internal/view/loanview"
	collection "lms-backend/pkg/collectionquery"

	"github.com/gofiber/fiber/v2"
)

// @Summary List loans
// @Description List loans depending on collection query
// @Tags loan
// @Accept application/json
// @Param offset query int false "Offset for pagination"
// @Param limit query int false "Limit for pagination"
// @Param filter[user_id] query int false "Filter by user ID"
// @Param sortBy query string false "Sort by column name"
// @Param orderBy query string false "Order by direction (asc or desc)"
// @Produce application/json
// @Success 200 {object} api.SwgResponse[[]loanview.View]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/loan [get]
func HandleList(c *fiber.Ctx) error {
	err := policy.Authorize(c, readLoanAction, loanpolicy.ReadPolicy())
	if err != nil {
		return err
	}

	cq := collection.GetCollectionQueryFromParam(c)
	db := database.GetDB()

	totalCount, err := loan.Count(db)
	if err != nil {
		return err
	}

	dbFiltered := cq.Filter(db, loan.Filters())

	filteredCount, err := loan.Count(dbFiltered)
	if err != nil {
		return err
	}

	dbSorted := cq.Sort(dbFiltered, loan.Sorters())
	dbPaginated := cq.Paginate(dbSorted)

	lns, err := loan.List(dbPaginated)
	if err != nil {
		return err
	}

	var view = []*loanview.View{}
	for _, ln := range lns {
		//nolint:gosec // loop does not modify struct
		view = append(view, loanview.ToView(&ln))
	}

	return c.JSON(api.Response{
		Data: view,
		Meta: api.Meta{
			TotalCount:    totalCount,
			FilteredCount: filteredCount,
		},
		Messages: api.Messages(
			api.SilentMessage("loans listed successfully"),
		),
	})
}
