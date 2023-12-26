package router

import (
	"lms-backend/internal/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Router func(fiber.Router)

func Route(parent fiber.Router, prefix string, r Router, handlers ...fiber.Handler) {
	router := parent.Group(prefix, handlers...)
	r(router)
}

func CachedRoute(parent fiber.Router, prefix string, t time.Duration, r Router, handlers ...fiber.Handler) {
	handlers = append(handlers, middleware.CacheMiddleware(t))
	router := parent.Group(prefix, handlers...)
	r(router)
}
