package database

import (
	"lms-backend/internal/config"
	"strings"
)

func PGDSNBuilder(cf *config.Config) (string, error) {
	dsn := strings.Builder{}

	if cf.PGUsername != "" {
		_, err := dsn.WriteString("user=" + cf.PGUsername)
		if err != nil {
			return "", err
		}
	}

	if cf.PGPassword != "" {
		_, err := dsn.WriteString(" password=" + cf.PGPassword)
		if err != nil {
			return "", err
		}
	}

	if cf.PGHost != "" {
		_, err := dsn.WriteString(" host=" + cf.PGHost)
		if err != nil {
			return "", err
		}
	}

	if cf.PGDatabase != "" {
		_, err := dsn.WriteString(" dbname=" + cf.PGDatabase)
		if err != nil {
			return "", err
		}
	}

	if cf.PGPort != "" {
		_, err := dsn.WriteString(" port=" + cf.PGPort)
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
