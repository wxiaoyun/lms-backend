package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func CacheMiddleware(t time.Duration) fiber.Handler {
	return cache.New(cache.Config{
		Expiration: t,
	})
}
