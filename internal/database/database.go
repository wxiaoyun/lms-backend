package database

import (
	"database/sql"
	"fmt"
	"log"
	"technical-test/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq" // postgres driver
)

var DB *gorm.DB

func OpenDataBase(cf *config.Config) error {
	var err error

	dsn, err := DSNBuilder(cf)
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
	return DB.Session(&gorm.Session{NewDB: true})
}

func ConnectToDB(cf *config.Config) (*sql.DB, error) {
	var err error

	dsn, err := DSNBuilder(cf)
	if err != nil {
		return nil, err
	}

	pgdb, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return pgdb, nil
}

func ConnectToDefaultDB(cf *config.Config) (*sql.DB, error) {
	var err error

	cf.DBName = "" // connect to default database

	dsn, err := DSNBuilder(cf)
	if err != nil {
		return nil, err
	}

	pgdb, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return pgdb, nil
}

func CreateDB(cf *config.Config) error {
	dbName := cf.DBName // Note that CREATE DATABASE cannot be executed within a transaction block.

	pgdb, err := ConnectToDefaultDB(cf)
	if err != nil {
		return err
	}

	_, err = pgdb.Exec(fmt.Sprintf("CREATE DATABASE \"%s\"", dbName))
	if err != nil {
		//nolint:revive // ignore lint
		log.Fatalf("Failed to create database '%s': %v\n", dbName, err)
	}
	log.Printf("Successfully created database '%s'\n", dbName)

	return nil
}

func DropDB(cf *config.Config) error {
	dbName := cf.DBName

	pgdb, err := ConnectToDefaultDB(cf)
	if err != nil {
		return err
	}

	_, err = pgdb.Exec(fmt.Sprintf("DROP DATABASE \"%s\"", dbName))
	if err != nil {
		//nolint:revive // ignore lint
		log.Fatalf("Failed to drop database '%s': %v\n", dbName, err)
	}
	log.Printf("Successfully drop database '%s'\n", dbName)

	return nil
}
