package router

import (
	"github.com/gofiber/fiber/v2"
)

type Router func(fiber.Router)

func Route(parent fiber.Router, prefix string, r Router) {
	router := parent.Group(prefix)
	r(router)
}
