package router

import (
	"lms-backend/internal/handler/auth"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(r fiber.Router) {
	r.Post("/signin", auth.HandleSignIn)
}
