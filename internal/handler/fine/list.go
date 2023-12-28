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

	dbFiltered := cq.Filter(db, fine.Filters(), fine.JoinLoan)

	filteredCount, err := fine.Count(dbFiltered)
	if err != nil {
		return err
	}

	dbSorted := cq.Sort(dbFiltered, fine.Sorters())
	dbPaginated := cq.Paginate(dbSorted)

	fns, err := fine.ListDetailed(dbPaginated)
	if err != nil {
		return err
	}

	var view = []*fineview.DetailedView{}
	for _, f := range fns {
		//nolint:gosec // loop does not modify struct
		view = append(view, fineview.ToDetailedView(&f))
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
