package sessionmiddleware

import (
	"auth-practice/internal/session"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func SessionMiddleware(c *fiber.Ctx) error {
	// skip auth routes - /api/v1/auth/*
	fmt.Println(c.Path())
	if strings.Split(c.Path(), "/")[3] == "auth" {
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
