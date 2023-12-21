package sessionmiddleware

import (
	"lms-backend/internal/api"
	"lms-backend/internal/session"

	"github.com/gofiber/fiber/v2"
)

func SessionMiddleware(c *fiber.Ctx) error {
	sess, err := session.Store.Get(c)
	if err != nil {
		return err
	}

	token := sess.Get(session.CookieKey)
	if token == nil {
		err := c.JSON(api.Response{
			Messages: api.Messages(api.InfoMessage("User is not logged in")),
		})
		if err != nil {
			return err
		}
		if err := sess.Destroy(); err != nil {
			return err
		}
		return fiber.NewError(fiber.StatusUnauthorized, "User is not logged in")
	}

	c.Locals(session.UserIDKey, token)

	return c.Next()
}
