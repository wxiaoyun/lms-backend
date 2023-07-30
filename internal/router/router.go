package router

import (
	"github.com/gofiber/fiber/v2"

	"technical-test/internal/handler/auth"
	"technical-test/internal/handler/health"
	userhandler "technical-test/internal/handler/user"
	worksheethandler "technical-test/internal/handler/worksheet"
)

func SetUpRoutes(app *fiber.App) error {
	v1Routes := app.Group("/api/v1")

	publicRoutes := v1Routes.Group("")
	publicRoutes.Get("/health", health.HandleHealth)

	authRoutes := publicRoutes.Group("/auth")
	authRoutes.Get("/currentuser", userhandler.HandleGetCurrentUser)
	authRoutes.Post("/signup", userhandler.HandleCreateUser)
	authRoutes.Post("/login", auth.HandleSignIn)
	authRoutes.Get("/logout", auth.HandleSignOut)

	privateRoutes := v1Routes.Group("")

	worksheetRoutes := privateRoutes.Group("/worksheet")
	worksheetRoutes.Get("/summary", worksheethandler.HandleWorksheetSummary)
	worksheetRoutes.Get("/find", worksheethandler.HandleFind)
	worksheetRoutes.Get("/", worksheethandler.HandleList)
	worksheetRoutes.Get("/:id", worksheethandler.HandleRead)

	return nil
}
