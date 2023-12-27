package router

import (
	"lms-backend/internal/handler/health"
	"lms-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func HealthRoutes(r fiber.Router) {
	r.Get("/", middleware.CacheMiddleware(middleware.VVLongExp), health.HandleHealth)
}
