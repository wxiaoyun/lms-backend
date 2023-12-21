// Package app
//
// Runs the app
package app

import (
	"lms-backend/internal/api"
	"lms-backend/internal/config"
	"lms-backend/internal/cron"
	"lms-backend/internal/database"
	"lms-backend/internal/router"

	"github.com/gofiber/fiber/v2"
)

// SetupAndRunApp sets up the app and runs it
func SetupAndRunApp() error {
	cfg, err := config.LoadEnvAndGetConfig()
	if err != nil {
		return err
	}

	err = database.OpenDataBase(cfg)
	if err != nil {
		return err
	}

	// create app
	app := fiber.New(fiber.Config{
		AppName:      cfg.AppName,
		ErrorHandler: api.ErrorHandler,
	})

	// setup routes
	router.SetUpRoutes(app, cfg)

	// attach swagger
	AddSwaggerRoutes(app)

	c := cron.RunJobs()
	defer c.Stop()

	// get the port and start
	return app.Listen(":" + cfg.Port)
}

// LoadEnvAndConnectToDB loads the environment variables and connects to the database
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
