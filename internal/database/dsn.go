package database

import (
	"errors"
	"fmt"
	"os"
)

func dnsBuilder() (string, error) {
	dbUsername, ok := os.LookupEnv("DB_USERNAME")
	if !ok {
		return "", errors.New("DB_USERNAME not found")
	}

	dbPassword, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		return "", errors.New("DB_PASSWORD not found")
	}

	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return "", errors.New("DB_HOST not found")
	}

	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return "", errors.New("DB_NAME not found")
	}

	dbPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return "", errors.New("DB_PORT not found")
	}

	sslMode, ok := os.LookupEnv("SSL_MODE")
	if !ok {
		return "", errors.New("SSL_MODE not found")
	}

	return fmt.Sprintf(
		"%s:%s@%s:%s/%s?sslmode=%s",
		dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
		sslMode,
	), nil
}
