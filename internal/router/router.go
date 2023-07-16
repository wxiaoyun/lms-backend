package router

import (
	"github.com/gofiber/fiber/v2"

	"auth-practice/internal/handler/auth"
	"auth-practice/internal/handler/health"
	userhandler "auth-practice/internal/handler/user"
)

func SetUpRoutes(app *fiber.App) error {
	v1Routes := app.Group("/api/v1")

	publicRoutes := v1Routes.Group("")
	publicRoutes.Get("/heath", health.HandleHealth)

	privateRoutes := v1Routes.Group("")
	privateRoutes.Post("/auth/signup", userhandler.HandleCreateUser)
	privateRoutes.Post("/auth/login", auth.HandleSignIn)
	privateRoutes.Get("/auth/logout", auth.HandleSignOut)

	return nil
}
