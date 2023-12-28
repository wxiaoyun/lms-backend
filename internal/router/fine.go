package router

import (
	finehandler "lms-backend/internal/handler/fine"
	"lms-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func FineRoutes(r fiber.Router) {
	r.Get("/", middleware.CacheMiddleware(middleware.VShortExp), finehandler.HandleList)

	Route(r, "/:fine_id", func(r fiber.Router) {
		r.Patch("/settle", finehandler.HandleSettle)
		r.Delete("/", finehandler.HandleDelete)
	})
}
