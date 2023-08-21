package auditloghandler

import (
	"fmt"
	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	audlog "lms-backend/internal/dataaccess/auditlog"
	"lms-backend/internal/params/auditlogparams"
	"lms-backend/internal/session"
	"lms-backend/internal/view/auditlogview"

	"github.com/gofiber/fiber/v2"
)

// @Summary list audit logs
// @Description list relevang audit logs
// @Tags audit log
// @Accept application/json
// @Param audit_log body auditlogparams.BaseParams true "Audit log creation request"
// @Produce application/json
// @Success 200 {object} api.SwgMsgResponse[auditlogview.View]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/audit_log/ [post]
func HandleCreate(c *fiber.Ctx) error {
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
	tx, rollBackOrCommit := audit.Begin(c, log.Action)
	defer func() { rollBackOrCommit(err) }()

	log, err = audlog.Create(tx, log)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(api.Response{
		Data: auditlogview.ToView(log),
		Messages: api.Messages(
			api.SuccessMessage(fmt.Sprintf(
				"Entry in audit log created successfully: %s", log.Action,
			))),
	})
}
