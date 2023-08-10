package database

import (
	"lms-backend/internal/config"
	"strings"
)

func DSNBuilder(cf *config.Config) (string, error) {
	dsn := strings.Builder{}

	if cf.DBUsername != "" {
		_, err := dsn.WriteString("user=" + cf.DBUsername)
		if err != nil {
			return "", err
		}
	}

	if cf.DBPassword != "" {
		_, err := dsn.WriteString(" password=" + cf.DBPassword)
		if err != nil {
			return "", err
		}
	}

	if cf.DBHost != "" {
		_, err := dsn.WriteString(" host=" + cf.DBHost)
		if err != nil {
			return "", err
		}
	}

	if cf.DBName != "" {
		_, err := dsn.WriteString(" dbname=" + cf.DBName)
		if err != nil {
			return "", err
		}
	}

	if cf.DBPort != "" {
		_, err := dsn.WriteString(" port=" + cf.DBPort)
		if err != nil {
			return "", err
		}
	}

	if cf.SSLMode != "" {
		_, err := dsn.WriteString(" sslmode=" + cf.SSLMode)
		if err != nil {
			return "", err
		}
	}

	return dsn.String(), nil
}
