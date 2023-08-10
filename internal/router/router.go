package router

import (
	"github.com/gofiber/fiber/v2"

	auditloghandler "lms-backend/internal/handler/auditlog"
	"lms-backend/internal/handler/auth"
	bookhandler "lms-backend/internal/handler/book"
	"lms-backend/internal/handler/health"
	userhandler "lms-backend/internal/handler/user"
	worksheethandler "lms-backend/internal/handler/worksheet"
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

	auditlogRoutes := privateRoutes.Group("/audit_log")
	auditlogRoutes.Get("/", auditloghandler.HandleList)
	auditlogRoutes.Post("/", auditloghandler.HandleCreate)

	bookRoutes := privateRoutes.Group("/book")
	bookRoutes.Post("/", bookhandler.HandleCreate)

	return nil
}
