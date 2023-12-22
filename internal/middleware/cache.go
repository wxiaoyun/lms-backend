package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func SetupCache(r fiber.Router, t time.Duration) {
	r.Use(cache.New(cache.Config{
		Expiration: t,
	}))
}
