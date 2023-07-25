package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq" // postgres driver
)

var db *gorm.DB

func OpenDataBase() error {
	var err error

	dsn, err := dnsBuilder(false)
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

func ConnectToDefaultDB() (*sql.DB, error) {
	var err error

	dsn, err := dnsBuilder(true)
	if err != nil {
		return nil, err
	}

	pgdb, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return pgdb, nil
}

func CreateDB() error {
	pgdb, err := ConnectToDefaultDB()
	if err != nil {
		return err
	}

	dbName := os.Getenv("DB_NAME") // Note that CREATE DATABASE cannot be executed within a transaction block.
	_, err = pgdb.Exec(fmt.Sprintf("CREATE DATABASE \"%s\"", dbName))
	if err != nil {
		//nolint:revive // ignore lint
		log.Fatalf("Failed to create database '%s': %v\n", dbName, err)
	}
	log.Printf("Successfully created database '%s'\n", dbName)

	return nil
}

func DropDB() error {
	pgdb, err := ConnectToDefaultDB()
	if err != nil {
		return err
	}

	dbName := os.Getenv("DB_NAME")
	_, err = pgdb.Exec(fmt.Sprintf("DROP DATABASE \"%s\"", dbName))
	if err != nil {
		//nolint:revive // ignore lint
		log.Fatalf("Failed to drop database '%s': %v\n", dbName, err)
	}
	log.Printf("Successfully drop database '%s'\n", dbName)

	return nil
}
