package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func SetupStaticFile(app *fiber.App) {
	app.Use("/v1/file", filesystem.New(filesystem.Config{
		Root: http.Dir("./file_storage"),
	}))
}
