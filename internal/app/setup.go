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

	database.SetupRedis(cfg)
	err = database.SetupPostgres(cfg)
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

	c := cron.RunJobs()
	defer c.Stop()

	// get the port and start
	return app.Listen(":" + cfg.Port)
}

// LoadEnvAndConnectToDB loads the environment variables and connects to the database
// Used for running sql migrations and seeding data
func LoadEnvAndConnectToDB() error {
	cfg, err := config.LoadEnvAndGetConfig()
	if err != nil {
		return err
	}
	err = database.SetupPostgres(cfg)
	if err != nil {
		return err
	}

	return nil
}
