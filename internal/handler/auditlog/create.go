package auditloghandler

import (
	"fmt"
	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	audlog "lms-backend/internal/dataaccess/auditlog"
	"lms-backend/internal/params/auditlogparams"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/auditlogpolicy"
	"lms-backend/internal/session"
	"lms-backend/internal/view/auditlogview"

	"github.com/gofiber/fiber/v2"
)

const (
	createAuditLogAction = "create audit log entry"
)

func HandleCreate(c *fiber.Ctx) error {
	err := policy.Authorize(c, createAuditLogAction, auditlogpolicy.CreatePolicy())
	if err != nil {
		return err
	}

	var params auditlogparams.BaseParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	if err := params.Validate(); err != nil {
		return err
	}

	userID, err := session.GetLoginSession(c)
	if err != nil {
		return err
	}

	log := params.ToModel(userID)
	tx, rollBackOrCommit := audit.Begin(c, fmt.Sprintf("User id - /'%d/' creating an entry in audit log", userID))
	defer func() { rollBackOrCommit(err) }()

	log, err = audlog.Create(tx, log)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(api.Response{
		Data: auditlogview.ToDetailedView(log),
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"Entry in audit log created successfully: %s", log.Action,
			))),
	})
}
