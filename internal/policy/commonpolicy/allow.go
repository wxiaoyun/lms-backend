package commonpolicy

import (
	"lms-backend/internal/policy"

	"github.com/gofiber/fiber/v2"
)

type Allow struct{}

func AllowAll() Allow { return Allow{} }

func (Allow) Validate(_ *fiber.Ctx) (policy.Decision, error) {
	return policy.Allow, nil
}

func (Allow) Reason() string {
	return "Everyone is allowed to do this."
}
