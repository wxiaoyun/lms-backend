package router

import (
	userhandler "lms-backend/internal/handler/user"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(r fiber.Router) {
	r.Get("/", userhandler.HandleList)
	r.Post("/", userhandler.HandleCreate)

	Route(r, "/:user_id", func(r fiber.Router) {
		r.Get("/", userhandler.HandleRead)
		r.Patch("/", userhandler.HandleUpdate)
		r.Delete("/", userhandler.HandleDelete)

		r.Patch("/role", userhandler.HandleChangeRole)
	})

	Route(r, "/autocomplete", func(r fiber.Router) {
		r.Get("/:value", userhandler.HandleAutoComplete)
	})
}
