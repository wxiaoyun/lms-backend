package commonpolicy

import (
	"fmt"
	"lms-backend/internal/policy"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AnyOf struct {
	Policies  []policy.Policy
	ReasonStr string
}

func Any(polices ...policy.Policy) *AnyOf {
	return &AnyOf{
		Policies: polices,
	}
}

func (a *AnyOf) Validate(c *fiber.Ctx) (policy.Decision, error) {
	builder := strings.Builder{}
	for i, p := range a.Policies {
		decision, err := p.Validate(c)
		if err != nil {
			return policy.Deny, err
		}

		if decision == policy.Allow {
			return policy.Allow, nil
		}
		//nolint
		builder.WriteString(fmt.Sprintf("%d. %s\n", i+1, p.Reason()))
	}

	a.ReasonStr = builder.String()
	return policy.Deny, nil
}

func (a *AnyOf) Reason() string {
	return "At least one of the policies policies must be satisfied. You can't do this because:\n" + a.ReasonStr
}
