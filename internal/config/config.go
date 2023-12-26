package config

import (
	"lms-backend/pkg/error/internalerror"
	"os"
	"strconv"
)

type Config struct {
	AppName       string
	PGHost        string
	PGPort        string
	PGUsername    string
	PGPassword    string
	PGDatabase    string
	SSLMode       string
	REDISHost     string
	REDISPort     int
	REDISUser     string
	REDISPassword string
	// URL standard format Redis URL. If this is set all other config options, Host, Port, Username, Password, Database have no effect.
	REDISURL    string
	Port        string
	FrontendURL string
}

// Returns a Config struct with the values from the environment variables
//
// Must be ran after loading env
func GetConfig() (*Config, error) {
	rp := os.Getenv("REDIS_PORT")
	redisPort, err := strconv.Atoi(rp)
	if err != nil {
		return nil, internalerror.InternalServerError("Bad Redis Port: " + rp)
	}

	return &Config{
		AppName:       os.Getenv("APP_NAME"),
		PGHost:        os.Getenv("PG_HOST"),
		PGPort:        os.Getenv("PG_PORT"),
		PGUsername:    os.Getenv("PG_USERNAME"),
		PGPassword:    os.Getenv("PG_PASSWORD"),
		PGDatabase:    os.Getenv("PG_DATABASE"),
		SSLMode:       os.Getenv("SSL_MODE"),
		REDISHost:     os.Getenv("REDIS_HOST"),
		REDISPort:     redisPort,
		REDISUser:     os.Getenv("REDIS_USER"),
		REDISPassword: os.Getenv("REDIS_PASSWORD"),
		REDISURL:      os.Getenv("REDIS_URL"),
		Port:          os.Getenv("PORT"),
		FrontendURL:   os.Getenv("FRONTEND_URL"),
	}, nil
}

func LoadEnvAndGetConfig() (*Config, error) {
	err := LoadEnv()
	if err != nil {
		return nil, err
	}

	return GetConfig()
}
