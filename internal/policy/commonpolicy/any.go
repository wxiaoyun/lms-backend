package commonpolicy

import (
	"lms-backend/internal/policy"

	"github.com/gofiber/fiber/v2"
)

type AnyOf struct {
	Policies []policy.Policy
}

func Any(polices ...policy.Policy) *AnyOf {
	return &AnyOf{
		Policies: polices,
	}
}

func (a *AnyOf) Validate(c *fiber.Ctx) (policy.Decision, error) {
	for _, p := range a.Policies {
		decision, err := p.Validate(c)
		if err != nil {
			return policy.Deny, err
		}

		if decision == policy.Allow {
			return policy.Allow, nil
		}
	}

	return policy.Deny, nil
}
