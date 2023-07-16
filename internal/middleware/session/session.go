package sessionmiddleware

import (
	"auth-practice/internal/session"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func SessionMiddleware(c *fiber.Ctx) error {
	// skip auth routes - /api/v1/auth/*
	paths := strings.Split(c.Path(), "/")
	if paths[1] == "swagger" || (len(paths) >= 3 && paths[3] == "auth") {
		return c.Next()
	}

	sess, err := session.Store.Get(c)
	if err != nil {
		return err
	}

	token := sess.Get(session.CookieKey)
	if token == nil {
		return fiber.ErrUnauthorized
	}

	return c.Next()
}
