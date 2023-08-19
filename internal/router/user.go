package router

import (
	userhandler "lms-backend/internal/handler/user"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(r fiber.Router) {
	r.Get("/currentuser", userhandler.HandleGetCurrentUser)
}
