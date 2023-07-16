package database

import (
	"errors"

	"auth-practice/internal/model"
)

func AutoMigration() error {
	if db == nil {
		return errors.New("database is not initialized")
	}

	err := db.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}

	return nil
}
