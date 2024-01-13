package router

import (
	bookcopyhandler "lms-backend/internal/handler/bookcopy"

	"github.com/gofiber/fiber/v2"
)

func BookcopyRoutes(r fiber.Router) {
	Route(r, "/:bookcopy_id", func(r fiber.Router) {
		r.Delete("/", bookcopyhandler.HandleDelete)
		r.Get("/qrcode", bookcopyhandler.HandleGenerateQRCode)
	})
}
