package router

import (
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	v1Routes := app.Group("/api/v1")

	Route(v1Routes, "/health", HealthRoutes)
	Route(v1Routes, "/auth", AuthRoutes)
	Route(v1Routes, "/user", UserRoutes)
	Route(v1Routes, "/audit_log", AuditLogRoutes)
	Route(v1Routes, "/book", BookRoutes)
}
