package app

import (
	"os"

	"github.com/gofiber/fiber/v2"

	"technical-test/internal/api"
	"technical-test/internal/config"
	"technical-test/internal/database"
	"technical-test/internal/middleware"
	"technical-test/internal/router"
)

func SetupAndRunApp() error {
	var err error

	// load ENV
	err = LoadEnvAndConnectToDB()
	if err != nil {
		return err
	}

	// create app
	app := fiber.New(fiber.Config{
		AppName:      "Library Management System Backend",
		ErrorHandler: api.ErrorHandler,
	})

	// attach app middleware
	middleware.SetupAppMiddleware(app)

	// setup routes
	router.SetUpRoutes(app)

	// attach swagger
	AddSwaggerRoutes(app)

	// get the port and start
	port := os.Getenv("PORT")
	return app.Listen(":" + port)
}

func LoadEnvAndConnectToDB() error {
	cf, err := config.LoadEnvAndGetConfig()
	if err != nil {
		panic(err)
	}
	err = database.OpenDataBase(cf)
	if err != nil {
		return err
	}

	return nil
}
