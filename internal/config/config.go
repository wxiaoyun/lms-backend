package config

import (
	"os"
)

type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	SSLMode    string
	Port       string
}

// Returns a Config struct with the values from the environment variables
//
// Must be ran after loading env
func GetConfig() *Config {
	return &Config{
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		SSLMode:    os.Getenv("SSL_MODE"),
		Port:       os.Getenv("PORT"),
	}
}

func LoadEnvAndGetConfig() (*Config, error) {
	err := LoadENV()
	if err != nil {
		return nil, err
	}

	return GetConfig(), nil
}
