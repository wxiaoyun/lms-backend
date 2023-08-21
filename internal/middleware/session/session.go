package sessionmiddleware

import (
	"lms-backend/internal/api"
	"lms-backend/internal/session"
	"lms-backend/util/sliceutil"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func SessionMiddleware(c *fiber.Ctx) error {
	paths := strings.Split(c.Path(), "/")
	if sliceutil.Contains(paths, "swagger") {
		return c.Next()
	}
	if sliceutil.Contains(paths, "signin") {
		return c.Next()
	}
	if sliceutil.Contains(paths, "signup") {
		return c.Next()
	}
	if sliceutil.Contains(paths, "health") {
		return c.Next()
	}

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

	return c.Next()
}
