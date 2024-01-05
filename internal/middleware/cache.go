package middleware

import (
	"fmt"
	"lms-backend/internal/database"
	"sort"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/utils"
)

var (
	// VShortExp = 1 * time.Minute
	// ShortExp  = 5 * time.Minute
	// MedExp    = 30 * time.Minute
	// LongExp   = 1 * time.Hour
	// VLongExp  = 6 * time.Hour
	// VVLongExp = 24 * time.Hour

	// testing
	VShortExp = 1 * time.Second
	ShortExp  = 1 * time.Second
	MedExp    = 1 * time.Second
	LongExp   = 1 * time.Second
	VLongExp  = 1 * time.Second
	VVLongExp = 1 * time.Second
)

func KeyGenerator(c *fiber.Ctx) string {
	path := c.Path()
	queryParams := c.Queries()

	if len(queryParams) == 0 {
		return utils.CopyString(path)
	}

	var params []string
	for key, value := range queryParams {
		params = append(params, key+"="+value)
	}

	sort.Strings(params)

	return fmt.Sprintf("%s?%s", path, strings.Join(params, "&"))
}

func CacheMiddleware(t time.Duration) fiber.Handler {
	return cache.New(cache.Config{
		Storage:      database.GetRedisStore(),
		KeyGenerator: KeyGenerator,
		Expiration:   t,
	})
}
