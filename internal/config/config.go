package config

import (
	"os"
)

type Config struct {
	AppName     string
	DBUsername  string
	DBPassword  string
	DBHost      string
	DBPort      string
	DBName      string
	SSLMode     string
	Port        string
	FrontendURL string
}

// Returns a Config struct with the values from the environment variables
//
// Must be ran after loading env
func GetConfig() *Config {
	return &Config{
		AppName:     os.Getenv("APP_NAME"),
		DBUsername:  os.Getenv("DB_USERNAME"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      os.Getenv("DB_PORT"),
		DBName:      os.Getenv("DB_NAME"),
		SSLMode:     os.Getenv("SSL_MODE"),
		Port:        os.Getenv("PORT"),
		FrontendURL: os.Getenv("FRONTEND_URL"),
	}
}

func LoadEnvAndGetConfig() (*Config, error) {
	err := LoadENV()
	if err != nil {
		return nil, err
	}

	return GetConfig(), nil
}
