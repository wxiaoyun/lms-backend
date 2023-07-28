package testutil

// Utilites for testing

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" || goEnv == "development" {
		err := godotenv.Load("../../.env.development")
		if err != nil {
			return nil, err
		}
	}

	dsn := strings.Builder{}

	dbUsername, ok := os.LookupEnv("DB_USERNAME")
	if ok {
		_, err := dsn.WriteString("user=" + dbUsername)
		if err != nil {
			return nil, err
		}
	}

	dbPassword, ok := os.LookupEnv("DB_PASSWORD")
	if ok {
		_, err := dsn.WriteString(" password=" + dbPassword)
		if err != nil {
			return nil, err
		}
	}

	dbHost, ok := os.LookupEnv("DB_HOST")
	if ok {
		_, err := dsn.WriteString(" host=" + dbHost)
		if err != nil {
			return nil, err
		}
	}

	dbName, ok := os.LookupEnv("DB_NAME")
	if ok {
		_, err := dsn.WriteString(" dbname=" + dbName)
		if err != nil {
			return nil, err
		}
	}

	dbPort, ok := os.LookupEnv("DB_PORT")
	if ok {
		_, err := dsn.WriteString(" port=" + dbPort)
		if err != nil {
			return nil, err
		}
	}

	sslMode, ok := os.LookupEnv("SSL_MODE")
	if ok {
		_, err := dsn.WriteString(" sslmode=" + sslMode)
		if err != nil {
			return nil, err
		}
	}

	return gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})
}
