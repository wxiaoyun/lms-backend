package auditlogparams

import (
	"lms-backend/internal/model"

	"github.com/gofiber/fiber/v2"
)

type BaseParams struct {
	Action string `json:"action"`
}

func (b *BaseParams) Validate() error {
	if b.Action == "" {
		return fiber.NewError(fiber.StatusBadRequest, "action is required")
	}

	return nil
}

func (b *BaseParams) ToModel(userID int64) *model.AuditLog {
	return &model.AuditLog{
		Action: b.Action,
		UserID: uint(userID),
	}
}
