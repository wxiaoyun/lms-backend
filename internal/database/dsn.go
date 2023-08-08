package database

import (
	"strings"
	"technical-test/internal/config"
)

func GormDSNBuilder(config *config.Config) (string, error) {
	dsn := strings.Builder{}

	_, err := dsn.WriteString("user=" + config.DBUsername)
	if err != nil {
		return "", err
	}

	_, err = dsn.WriteString(" password=" + config.DBPassword)
	if err != nil {
		return "", err
	}

	_, err = dsn.WriteString(" host=" + config.DBHost)
	if err != nil {
		return "", err
	}

	_, err = dsn.WriteString(" dbname=" + config.DBName)
	if err != nil {
		return "", err
	}

	_, err = dsn.WriteString(" port=" + config.DBPort)
	if err != nil {
		return "", err
	}

	_, err = dsn.WriteString(" sslmode=" + config.SSLMode)
	if err != nil {
		return "", err
	}

	return dsn.String(), nil
}

func DSNBuilder(config *config.Config) (string, error) {
	dsn := strings.Builder{}

	_, err := dsn.WriteString("user=" + config.DBUsername)
	if err != nil {
		return "", err
	}

	_, err = dsn.WriteString(" password=" + config.DBPassword)
	if err != nil {
		return "", err
	}

	_, err = dsn.WriteString(" host=" + config.DBHost)
	if err != nil {
		return "", err
	}

	_, err = dsn.WriteString(" port=" + config.DBPort)
	if err != nil {
		return "", err
	}

	_, err = dsn.WriteString(" sslmode=" + config.SSLMode)
	if err != nil {
		return "", err
	}

	return dsn.String(), nil
}
