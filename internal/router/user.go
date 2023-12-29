package router

import (
	userhandler "lms-backend/internal/handler/user"
	"lms-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(r fiber.Router) {
	r.Get("/", middleware.CacheMiddleware(middleware.ShortExp), userhandler.HandleList)

	Route(r, "/:user_id", func(r fiber.Router) {
		r.Get("/", userhandler.HandleRead)
		r.Patch("/", userhandler.HandleUpdate)
		r.Delete("/", userhandler.HandleDelete)

		r.Patch("/role", userhandler.HandleChangeRole)
	})

	Route(r, "/autocomplete", func(r fiber.Router) {
		r.Get("/:value", middleware.CacheMiddleware(middleware.ShortExp), userhandler.HandleAutoComplete)
	})
}
