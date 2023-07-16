package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var DB *gorm.DB

func OpenDataBase() error {
	var err error

	dsn, err := dnsBuilder()
	if err != nil {
		return err
	}

	DB, err = gorm.Open(postgres.Open(dsn), GetConfig())
	if err != nil {
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return DB
}
