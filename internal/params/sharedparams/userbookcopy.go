package sharedparams

import (
	"github.com/gofiber/fiber/v2"
)

type UserBookcopyParams struct {
	UserID     int64 `json:"user_id"`
	BookCopyID int64 `json:"book_copy_id"`
}

func (params *UserBookcopyParams) Validate() error {
	if params.UserID <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "user_id is required")
	}

	if params.BookCopyID <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "book_copy_id is required")
	}

	return nil
}
