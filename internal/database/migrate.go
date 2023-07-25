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
		&model.Question{},
	)
	if err != nil {
		return err
	}

	return nil
}
