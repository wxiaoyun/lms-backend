package router

import (
	"lms-backend/internal/handler/auth"
	userhandler "lms-backend/internal/handler/user"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(r fiber.Router) {
	r.Post("/signup", userhandler.HandleCreateUser)
	r.Post("/login", auth.HandleSignIn)
	r.Get("/logout", auth.HandleSignOut)
}
