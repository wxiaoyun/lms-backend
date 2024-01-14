package router

import (
	bookhandler "lms-backend/internal/handler/book"
	bookcopyhandler "lms-backend/internal/handler/bookcopy"
	bookmarkhandler "lms-backend/internal/handler/bookmark"
	"lms-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func BookRoutes(r fiber.Router) {
	r.Post("/", bookhandler.HandleCreate)

	Route(r, "/:book_id", func(r fiber.Router) {
		r.Patch("/", bookhandler.HandleUpdate)
		r.Patch("/thumbnail", bookhandler.HandleUpdateThumbnail)
		r.Delete("/", bookhandler.HandleDelete)

		Route(r, "/bookmark", BookBookmarkRoutes)
		Route(r, "/bookcopy", BookBookcopyRoutes)
	})

	Route(r, "/autocomplete", func(r fiber.Router) {
		r.Get("/:value", middleware.CacheMiddleware(middleware.ShortExp), bookhandler.HandleAutoComplete)
	})
}

func BookBookmarkRoutes(r fiber.Router) {
	r.Post("/", bookmarkhandler.HandleCreate)
}

func BookBookcopyRoutes(r fiber.Router) {
	r.Post("/", bookcopyhandler.HandleCreate)
}
