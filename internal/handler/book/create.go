package bookhandler

import (
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/bookpolicy"

	"github.com/gofiber/fiber/v2"
)

const (
	createBookAction = "create book"
)

func HandleCreate(c *fiber.Ctx) error {
	err := policy.Authorize(c, createBookAction, bookpolicy.CreatePolicy())
	if err != nil {
		return err
	}

	return nil
}
