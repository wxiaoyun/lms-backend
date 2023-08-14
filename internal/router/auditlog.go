package router

import (
	auditloghandler "lms-backend/internal/handler/auditlog"

	"github.com/gofiber/fiber/v2"
)

func AuditLogRoutes(r fiber.Router) {
	r.Get("/", auditloghandler.HandleList)
	r.Post("/", auditloghandler.HandleCreate)
}
