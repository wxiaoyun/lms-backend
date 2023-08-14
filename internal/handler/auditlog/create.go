package auditloghandler

import (
	"fmt"
	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	audlog "lms-backend/internal/dataaccess/auditlog"
	"lms-backend/internal/database"
	"lms-backend/internal/params/auditlogparams"
	"lms-backend/internal/session"

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
	defer func() { rollBackOrCommit(err) }()

	log, err = audlog.Create(tx, log)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(api.Response{
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"Entry in audit log created successfully: %s", log.Action,
			))),
	})
}
