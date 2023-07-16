package app

import (
	"os"

	"github.com/gofiber/fiber/v2"

	"auth-practice/internal/api"
	"auth-practice/internal/config"
	"auth-practice/internal/database"
	"auth-practice/internal/middleware"
	"auth-practice/internal/router"
)

func SetupAndRunApp() error {
	var err error

	// load ENV
	err = config.LoadENV()
	if err != nil {
		return err
	}

	// open database
	err = database.OpenDataBase()
	if err != nil {
		return err
	}

	// create app
	app := fiber.New(fiber.Config{
		AppName:      "auth-practice",
		ErrorHandler: api.ErrorHandler,
	})

	// attach app middleware
	middleware.SetupAppMiddleware(app)

	// setup routes
	router.SetUpRoutes(app)

	// attach swagger
	config.AddSwaggerRoutes(app)

	// get the port and start
	port := os.Getenv("PORT")
	return app.Listen(":" + port)
}
