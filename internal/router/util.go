package router

import (
	"lms-backend/internal/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Router func(fiber.Router)

func Route(parent fiber.Router, prefix string, r Router) {
	router := parent.Group(prefix)
	r(router)
}

func CachedRoute(parent fiber.Router, prefix string, r Router) {
	router := parent.Group(prefix)
	middleware.SetupCache(router, time.Minute)
	r(router)
}
