package router

import (
	bookmarkhandler "lms-backend/internal/handler/bookmark"
	"lms-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func BookmarkRoutes(r fiber.Router) {
	r.Get("/", middleware.CacheMiddleware(middleware.VShortExp), bookmarkhandler.HandleList)
	r.Post("/", bookmarkhandler.HandleCreate)

	Route(r, "/:bookmark_id", func(r fiber.Router) {
		r.Delete("/", bookmarkhandler.HandleDelete)
	})
}
