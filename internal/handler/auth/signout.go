package auth

import (
	"lms-backend/internal/api"
	"lms-backend/internal/session"

	"github.com/gofiber/fiber/v2"
)

func HandleSignOut(c *fiber.Ctx) error {
	if !session.HasSession(c) {
		return c.Status(fiber.StatusOK).JSON(api.Response{
			Messages: []api.Message{api.ErrorMessage("User is not logged in")},
		})
	}

	// if handler, ok := c.Locals(csrf.ConfigDefault.HandlerContextKey).(*csrf.CSRFHandler); ok {
	// 	if err := handler.DeleteToken(c); err != nil {
	// 		return err
	// 	}
	// }

	sess, err := session.Store.Get(c)
	if err != nil {
		return err
	}

	err = sess.Destroy()
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(api.Response{
		Messages: api.Messages(
			api.SilentMessage("User is logged out successfully"),
		),
	})
}
