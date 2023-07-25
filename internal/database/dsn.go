package database

import (
	"os"
	"strings"
)

func dnsBuilder(isDefault bool) (string, error) {
	dsn := strings.Builder{}

	dbUsername, ok := os.LookupEnv("DB_USERNAME")
	if ok {
		_, err := dsn.WriteString("user=" + dbUsername)
		if err != nil {
			return "", err
		}
	}

	dbPassword, ok := os.LookupEnv("DB_PASSWORD")
	if ok {
		_, err := dsn.WriteString(" password=" + dbPassword)
		if err != nil {
			return "", err
		}
	}

	dbHost, ok := os.LookupEnv("DB_HOST")
	if ok {
		_, err := dsn.WriteString(" host=" + dbHost)
		if err != nil {
			return "", err
		}
	}

	dbName, ok := os.LookupEnv("DB_NAME")
	if ok && !isDefault {
		_, err := dsn.WriteString(" dbname=" + dbName)
		if err != nil {
			return "", err
		}
	}

	dbPort, ok := os.LookupEnv("DB_PORT")
	if ok {
		_, err := dsn.WriteString(" port=" + dbPort)
		if err != nil {
			return "", err
		}
	}

	sslMode, ok := os.LookupEnv("SSL_MODE")
	if ok {
		_, err := dsn.WriteString(" sslmode=" + sslMode)
		if err != nil {
			return "", err
		}
	}

	return dsn.String(), nil
}
