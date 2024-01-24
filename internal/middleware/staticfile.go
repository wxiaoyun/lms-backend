package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func SetupStaticFile(app *fiber.App) {
	app.Static("/assets", "./frontend/assets", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    false,
	})
	app.Use("/file", filesystem.New(filesystem.Config{
		Root: http.Dir("./file_storage"),
	}))
}
