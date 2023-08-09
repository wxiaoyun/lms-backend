package auditloghandler

import (
	"technical-test/internal/api"
	"technical-test/internal/dataaccess/auditlog"
	"technical-test/internal/database"
	"technical-test/internal/view/auditlogview"
	collection "technical-test/pkg/collectionquery"

	"github.com/gofiber/fiber/v2"
)

// @Summary list audit logs
// @Description list relevang audit logs
// @Tags audit log
// @Accept */*
// @Produce application/json
// @Success 200 "OK"
// @Router /api/v1/audit_log/ [get]
func HandleList(c *fiber.Ctx) error {
	db := database.GetDB()

	cq := collection.GetCollectionQueryFromParam(c)

	totalCount, err := auditlog.Count(db)
	if err != nil {
		return err
	}

	logs, err := auditlog.List(db, cq)
	if err != nil {
		return err
	}

	filteredCount, err := auditlog.CountFiltered(db, cq)
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
		Messages: []api.Message{
			api.SuccessMessage("auditlog listed successfully"),
		},
	})
}
