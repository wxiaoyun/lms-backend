package database

import (
	"lms-backend/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupPostgres(cfg *config.Config) error {
	var err error

	dsn, err := PGDSNBuilder(cfg)
	if err != nil {
		return err
	}

	DB, err = gorm.Open(postgres.Open(dsn), GetConfig())
	if err != nil {
		return err
	}

	return nil
}
