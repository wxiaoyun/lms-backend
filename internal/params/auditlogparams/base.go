package auditlogparams

import (
	"lms-backend/internal/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

type BaseParams struct {
	Action string `json:"action"`
	Date   string `json:"date"`
}

func (b *BaseParams) Validate() error {
	if b.Action == "" {
		return fiber.NewError(fiber.StatusBadRequest, "action is required")
	}

	if _, err := time.Parse(time.RFC3339, b.Date); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "date does not match RFC3339 format")
	}

	return nil
}

func (b *BaseParams) ToModel(userID int64) *model.AuditLog {
	//nolint // err is checked in Validate()
	date, _ := time.Parse(time.RFC3339, b.Date)
	return &model.AuditLog{
		UserID: uint(userID),
		Action: b.Action,
		Date:   date,
	}
}
