package router

import (
	"github.com/gofiber/fiber/v2"

	auditloghandler "lms-backend/internal/handler/auditlog"
	"lms-backend/internal/handler/auth"
	bookhandler "lms-backend/internal/handler/book"
	"lms-backend/internal/handler/health"
	userhandler "lms-backend/internal/handler/user"
)

func SetUpRoutes(app *fiber.App) {
	v1Routes := app.Group("/api/v1")

	publicRoutes := v1Routes.Group("")
	publicRoutes.Get("/health", health.HandleHealth)

	authRoutes := publicRoutes.Group("/auth")
	authRoutes.Get("/currentuser", userhandler.HandleGetCurrentUser)
	authRoutes.Post("/signup", userhandler.HandleCreateUser)
	authRoutes.Post("/login", auth.HandleSignIn)
	authRoutes.Get("/logout", auth.HandleSignOut)

	privateRoutes := v1Routes.Group("")

	auditlogRoutes := privateRoutes.Group("/audit_log")
	auditlogRoutes.Get("/", auditloghandler.HandleList)
	auditlogRoutes.Post("/", auditloghandler.HandleCreate)

	bookRoutes := privateRoutes.Group("/book")
	bookRoutes.Post("/", bookhandler.HandleCreate)
	bookRoutes.Get("/", bookhandler.HandleList)

	bookSpecificRoutes := bookRoutes.Group("/:id")
	bookSpecificRoutes.Get("/", bookhandler.HandleRead)
	bookSpecificRoutes.Patch("/", bookhandler.HandleUpdate)
	bookSpecificRoutes.Delete("/", bookhandler.HandleDelete)

	bookSpecificRoutes.Post("/loan", bookhandler.HandleLoan)
	bookSpecificRoutes.Post("/return", bookhandler.HandleReturn)
	bookSpecificRoutes.Post("/renew", bookhandler.HandleRenew)
	bookSpecificRoutes.Post("/reserve", bookhandler.HandleReserve)
}
