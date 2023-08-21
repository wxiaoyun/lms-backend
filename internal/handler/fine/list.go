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

// @Summary List fine
// @Description List fines belonging to a loan
// @Tags fine
// @Accept */*
// @Produce application/json
// @Success 200 {object} api.SwgResponse[[]fineview.View]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/fine?offset=0&limit=25&filter[user_id]=1&sortBy=created_at&orderBy=desc [get]
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
