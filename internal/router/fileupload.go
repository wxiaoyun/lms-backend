package router

import (
	fileuploadhandler "lms-backend/internal/handler/fileupload"

	"github.com/gofiber/fiber/v2"
)

func PrivateFileRoutes(r fiber.Router) {
	r.Post("/image", fileuploadhandler.HandleUploadImage)
}
