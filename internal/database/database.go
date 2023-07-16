package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq" // postgres driver
)

var db *gorm.DB

func OpenDataBase() error {
	var err error

	dsn, err := dnsBuilder()
	if err != nil {
		return err
	}

	db, err = gorm.Open(postgres.Open(dsn), GetConfig())
	if err != nil {
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}
