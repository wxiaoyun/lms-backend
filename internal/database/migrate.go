package database

import (
	"errors"

	"technical-test/internal/model"
)

func AutoMigration() error {
	if db == nil {
		return errors.New("database is not initialized")
	}

	err := db.AutoMigrate(
		&model.User{},
		&model.Worksheet{},
	)
	if err != nil {
		return err
	}

	return nil
}
