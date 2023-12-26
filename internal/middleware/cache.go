package middleware

import (
	"lms-backend/internal/database"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

var (
	VShortExp = 1 * time.Minute
	ShortExp  = 5 * time.Minute
	MedExp    = 30 * time.Minute
	LongExp   = 1 * time.Hour
	VLongExp  = 6 * time.Hour
	VVLongExp = 24 * time.Hour
)

func CacheMiddleware(t time.Duration) fiber.Handler {
	return cache.New(cache.Config{
		Storage:    database.GetRedisStore(),
		Expiration: t,
	})
}
