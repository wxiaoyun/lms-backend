package router

import (
	finehandler "lms-backend/internal/handler/fine"

	"github.com/gofiber/fiber/v2"
)

func FineRoutes(r fiber.Router) {
	r.Get("/", finehandler.HandleList)
}
