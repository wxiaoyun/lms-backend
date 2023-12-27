package userhandler

import (
	"lms-backend/internal/api"
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/database"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/userpolicy"
	"lms-backend/internal/view/userview"
	collection "lms-backend/pkg/collectionquery"

	"github.com/gofiber/fiber/v2"
)

func HandleList(c *fiber.Ctx) error {
	err := policy.Authorize(c, readUserAction, userpolicy.ListPolicy())
	if err != nil {
		return err
	}

	cq := collection.GetCollectionQueryFromParam(c)
	db := database.GetDB()

	totalCount, err := user.Count(db)
	if err != nil {
		return err
	}

	dbFiltered := cq.Filter(db, user.Filters())

	filteredCount, err := user.Count(dbFiltered)
	if err != nil {
		return err
	}

	dbSorted := cq.Sort(dbFiltered, user.Sorters())
	dbPaginated := cq.Paginate(dbSorted)

	rvs, err := user.List(dbPaginated)
	if err != nil {
		return err
	}

	var view = []*userview.View{}
	for _, r := range rvs {
		//nolint:gosec // loop does not modify struct
		view = append(view, userview.ToView(&r))
	}

	return c.JSON(api.Response{
		Data: view,
		Meta: api.Meta{
			TotalCount:    totalCount,
			FilteredCount: filteredCount,
		},
		Messages: api.Messages(
			api.SilentMessage("users listed successfully"),
		),
	})
}
