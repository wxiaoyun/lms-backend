package router

import (
	"lms-backend/internal/handler/googlebook"
	"lms-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func ExternalRoutes(r fiber.Router) {
	// We cache the response for 30 days since the data is not likely to change and we don't want to spam google api
	r.Get("/", middleware.CacheMiddleware(middleware.VVVLongExp), googlebook.HandleQuery)
}
