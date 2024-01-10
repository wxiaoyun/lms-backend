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
	VShortExp  = 1 * time.Minute     // 1 Minute
	ShortExp   = 5 * time.Minute     // 5 Minutes
	MedExp     = 30 * time.Minute    // 30 Minutes
	LongExp    = 1 * time.Hour       // 1 Hour
	VLongExp   = 6 * time.Hour       // 6 Hours
	VVLongExp  = 24 * time.Hour      // 24 Hours
	VVVLongExp = 30 * 24 * time.Hour // 30 Days
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
