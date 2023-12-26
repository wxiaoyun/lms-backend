package router

import (
	auditloghandler "lms-backend/internal/handler/auditlog"
	"lms-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuditLogRoutes(r fiber.Router) {
	r.Get("/", middleware.CacheMiddleware(middleware.MedExp), auditloghandler.HandleList)
	r.Post("/", auditloghandler.HandleCreate)
}
