package commonpolicy

import (
	"lms-backend/internal/policy"

	"github.com/gofiber/fiber/v2"
)

type AllOf struct {
	Policies []policy.Policy
}

func All(polices ...policy.Policy) *AllOf {
	return &AllOf{
		Policies: polices,
	}
}

func (a *AllOf) Validate(c *fiber.Ctx) (policy.Decision, error) {
	for _, p := range a.Policies {
		decision, err := p.Validate(c)
		if err != nil {
			return policy.Deny, err
		}

		if decision == policy.Deny {
			return policy.Deny, nil
		}
	}

	return policy.Allow, nil
}
