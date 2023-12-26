package config

import (
	"os"

	"github.com/joho/godotenv"
)

// This function will load the .env file if the GO_ENV environment variable is not set
func LoadEnv() error {
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" || goEnv == "development" {
		err := godotenv.Load(".env.development")
		if err != nil {
			return err
		}
	}
	return nil
}
