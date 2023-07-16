package router

import (
	"github.com/gofiber/fiber/v2"

	"auth-practice/internal/handler/auth"
	"auth-practice/internal/handler/health"
)

func SetUpRoutes(app *fiber.App) error {
	v1Routes := app.Group("/api/v2")

	publicRoutes := v1Routes.Group("")
	publicRoutes.Get("/heath", health.HandleHealth)
	publicRoutes.Post("/signup")

	privateRoutes := v1Routes.Group("")
	privateRoutes.Get("/auth/signin", auth.HandleSignIn)

	return app.Listen(":3000")
}
