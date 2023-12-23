package auditloghandler

import (
	"lms-backend/internal/api"
	"lms-backend/internal/dataaccess/auditlog"
	"lms-backend/internal/database"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/auditlogpolicy"
	"lms-backend/internal/view/auditlogview"
	collection "lms-backend/pkg/collectionquery"

	"github.com/gofiber/fiber/v2"
)

const (
	readAuditLogAction = "read audit logs"
)

// @Summary List audit logs
// @Description List relevant audit logs
// @Tags audit log
// @Accept application/json
// @Param offset query int false "Offset for pagination"
// @Param limit query int false "Limit for pagination"
// @Param filter[action] query string false "Filter by action"
// @Param sortBy query string false "Sort by column name"
// @Param orderBy query string false "Order by direction (asc or desc)"
// @Produce application/json
// @Success 200 {object} api.SwgResponse[[]auditlogview.View]
// @Failure 400 {object} api.SwgErrResponse
// @Router /v1/audit_log [get]
func HandleList(c *fiber.Ctx) error {
	err := policy.Authorize(c, readAuditLogAction, auditlogpolicy.ReadPolicy())
	if err != nil {
		return err
	}

	db := database.GetDB()
	cq := collection.GetCollectionQueryFromParam(c)

	totalCount, err := auditlog.Count(db)
	if err != nil {
		return err
	}

	dbFiltered := cq.Filter(db, auditlog.Filters(), auditlog.JoinUser)

	filteredCount, err := auditlog.Count(dbFiltered)
	if err != nil {
		return err
	}

	dbSorted := cq.Sort(dbFiltered, auditlog.Sorters())
	dbPaginated := cq.Paginate(dbSorted)
	logs, err := auditlog.ListDetailed(dbPaginated)
	if err != nil {
		return err
	}

	var view = []*auditlogview.DetailedView{}
	for _, log := range logs {
		//nolint:gosec // loop does not modify struct
		view = append(view, auditlogview.ToDetailedView(&log))
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
