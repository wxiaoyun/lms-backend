package cdbhelper

import (
	"lms-backend/internal/config"
	"lms-backend/internal/database"
	logger "lms-backend/internal/log"
)

var lgr = logger.StdoutLogger()

func CreateDB(cf *config.Config) error {
	lgr.Println("Creating database...")
	err := database.CreateDB(cf)
	if err != nil {
		return err
	}

	return nil
}
