package database

import (
	"errors"
	"technical-test/internal/model"
)

func AutoMigration() error {
	if DB == nil {
		return errors.New("database is not initialized")
	}

	err := DB.AutoMigrate(
		&model.User{},
		&model.Worksheet{},
		&model.Question{},
	)
	if err != nil {
		return err
	}

	return nil
}
