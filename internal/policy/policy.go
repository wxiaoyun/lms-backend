package policy

import (
	"fmt"
	"lms-backend/internal/api"
	"lms-backend/util/ternary"

	"github.com/gofiber/fiber/v2"
)

type Decision int

const (
	Allow Decision = iota
	Deny
)

type Policy interface {
	Validate(c *fiber.Ctx) (Decision, error)
}

func Authorize(c *fiber.Ctx, action string, policy Policy) error {
	decision, err := policy.Validate(c)
	if err != nil || decision == Deny {
		//nolint
		c.Status(fiber.StatusForbidden).JSON(api.Response{
			Messages: []api.Message{
				api.ErrorMessage(fmt.Sprintf("You are not allowed to %s.", action)),
			},
		})
		return fiber.NewError(fiber.StatusForbidden,
			ternary.If[string](err != nil).
				LazyThen(func() string { return err.Error() }).
				LazyElse(func() string { return "Unauthorized" }),
		)
	}

	return nil
}
