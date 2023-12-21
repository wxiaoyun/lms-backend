package sharedparams

import (
	"github.com/gofiber/fiber/v2"
)

type UserBookParams struct {
	UserID int64 `json:"user_id"`
	BookID int64 `json:"book_id"`
}

func (params *UserBookParams) Validate() error {
	if params.UserID <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "user_id is required")
	}

	if params.BookID <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "book_id is required")
	}

	return nil
}
