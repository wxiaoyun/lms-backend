package policy

import (
	"fmt"
	"lms-backend/pkg/error/externalerrors"

	"github.com/gofiber/fiber/v2"
)

type Decision int

const (
	Allow Decision = iota
	Deny
)

type Policy interface {
	Validate(c *fiber.Ctx) (Decision, error)
	Reason() string
}

func Authorize(c *fiber.Ctx, action string, policy Policy) error {
	decision, err := policy.Validate(c)
	if err != nil {
		return err
	}

	if decision == Deny {
		return externalerrors.Forbidden(
			fmt.Sprintf("You are not authorized to %s. %s", action, policy.Reason()),
		)
	}

	return nil
}
