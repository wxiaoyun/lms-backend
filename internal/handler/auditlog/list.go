package auditloghandler

import (
	"lms-backend/internal/api"
	"lms-backend/internal/dataaccess/auditlog"
	"lms-backend/internal/database"
	"lms-backend/internal/view/auditlogview"
	collection "lms-backend/pkg/collectionquery"

	"github.com/gofiber/fiber/v2"
)

// @Summary list audit logs
// @Description list relevang audit logs
// @Tags audit log
// @Accept */*
// @Produce application/json
// @Success 200 {object} api.SwgResponse[[]auditlogview.View]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/audit_log?offset=&limit=&filter[action]=<action>&sortBy=<col_name>&orderBy=<asc|desc> [get]
func HandleList(c *fiber.Ctx) error {
	db := database.GetDB()
	cq := collection.GetCollectionQueryFromParam(c)

	totalCount, err := auditlog.Count(db)
	if err != nil {
		return err
	}

	dbFiltered := cq.Filter(db, auditlog.Filters())

	filteredCount, err := auditlog.Count(dbFiltered)
	if err != nil {
		return err
	}

	dbSorted := cq.Sort(dbFiltered)
	dbPaginated := cq.Paginate(dbSorted)
	logs, err := auditlog.List(dbPaginated)
	if err != nil {
		return err
	}

	var view = []*auditlogview.View{}
	for _, log := range logs {
		//nolint:gosec // loop does not modify struct
		view = append(view, auditlogview.ToView(&log))
	}

	return c.JSON(api.Response{
		Data: view,
		Meta: api.Meta{
			TotalCount:    totalCount,
			FilteredCount: filteredCount,
		},
		Messages: api.Messages(
			api.SilentMessage("auditlog listed successfully"),
		),
	})
}
