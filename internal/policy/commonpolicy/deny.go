package commonpolicy

import (
	"lms-backend/internal/policy"

	"github.com/gofiber/fiber/v2"
)

type Deny struct{}

func DenyAll() Deny {
	return Deny{}
}

func (Deny) Validate(_ *fiber.Ctx) (policy.Decision, error) {
	return policy.Deny, nil
}

func (Deny) Reason() string {
	return "This action is not allowed by anyone."
}
