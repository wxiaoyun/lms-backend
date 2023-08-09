package auditloghandler

import (
	"fmt"
	"technical-test/internal/api"
	audit "technical-test/internal/auditlog"
	audlog "technical-test/internal/dataaccess/auditlog"
	"technical-test/internal/database"
	"technical-test/internal/params/auditlogparams"
	"technical-test/internal/session"

	"github.com/gofiber/fiber/v2"
)

// @Summary post audit logs
// @Description create an entry in the audit log
// @Tags audit log
// @Accept */*
// @Produce application/json
// @Success 200 "OK"
// @Router /api/v1/audit_log/ [post]
func HandleCreate(c *fiber.Ctx) error {
	var params auditlogparams.BaseParams
	err := c.BodyParser(&params)
	if err != nil {
		return err
	}

	err = params.Validate()
	if err != nil {
		return err
	}

	userID, err := session.GetLoginSession(c)
	if err != nil {
		return err
	}

	log := params.ToModel(userID)
	db := database.GetDB()
	tx, rollBackOrCommit := audit.Begin(c, db, log.Action)
	defer rollBackOrCommit()

	log, err = audlog.Create(tx, log)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(api.Response{
		Messages: []api.Message{
			api.SuccessMessage(fmt.Sprintf(
				"Entry in audit log created successfully: %s", log.Action,
			))},
	})
}
