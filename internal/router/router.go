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
	publicRoutes.Post("/signup", userhandler.HandleCreateUser)

	privateRoutes := v1Routes.Group("")
	privateRoutes.Get("/auth/signin", auth.HandleSignIn)

	return nil
}
