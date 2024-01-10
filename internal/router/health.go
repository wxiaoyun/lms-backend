package router

import (
	"lms-backend/internal/handler/health"

	"github.com/gofiber/fiber/v2"
)

func HealthRoutes(r fiber.Router) {
	r.Get("/", health.HandleHealth)
}
