package config

import (
	"lms-backend/pkg/error/internalerror"
	"os"
	"path/filepath"
	"strconv"
)

var (
	//nolint:errcheck
	cwd, _ = os.Getwd()
	// Path to current working directory, with symlinks evaluated.
	//nolint:errcheck
	RuntimeWorkingDirectory, _ = filepath.EvalSymlinks(cwd)

	GoogleAPIKey string
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
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		return nil, internalerror.InternalServerError("APP_NAME not set")
	}

	pgHost := os.Getenv("PG_HOST")
	if pgHost == "" {
		return nil, internalerror.InternalServerError("PG_HOST not set")
	}

	pgPort := os.Getenv("PG_PORT")
	if pgPort == "" {
		return nil, internalerror.InternalServerError("PG_PORT not set")
	}

	pgUsername := os.Getenv("PG_USERNAME")
	if pgUsername == "" {
		return nil, internalerror.InternalServerError("PG_USERNAME not set")
	}

	pgPassword := os.Getenv("PG_PASSWORD")
	if pgPassword == "" {
		return nil, internalerror.InternalServerError("PG_PASSWORD not set")
	}

	pgDatabase := os.Getenv("PG_DATABASE")
	if pgDatabase == "" {
		return nil, internalerror.InternalServerError("PG_DATABASE not set")
	}

	sslMode := os.Getenv("SSL_MODE")
	if sslMode == "" {
		return nil, internalerror.InternalServerError("SSL_MODE not set")
	}

	var redisHost, redisPassword, redisUser string
	var redisPort int
	redisURL := os.Getenv("REDIS_URL")

	// If REDIS_URL is set, all other config options have no effect
	if redisURL == "" {
		redisHost = os.Getenv("REDIS_HOST")
		if redisHost == "" {
			return nil, internalerror.InternalServerError("REDIS_HOST not set")
		}

		rp := os.Getenv("REDIS_PORT")

		if rp != "" {
			port, err := strconv.Atoi(rp)
			if err != nil {
				return nil, internalerror.InternalServerError("Bad Redis Port: " + rp)
			}
			redisPort = port
		}

		redisUser = os.Getenv("REDIS_USER")

		redisPassword = os.Getenv("REDIS_PASSWORD")
	}

	GoogleAPIKey := os.Getenv("GOOGLE_API_KEY")
	if GoogleAPIKey == "" {
		return nil, internalerror.InternalServerError("GOOGLE_API_KEY not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		return nil, internalerror.InternalServerError("PORT not set")
	}

	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		return nil, internalerror.InternalServerError("FRONTEND_URL not set")
	}

	return &Config{
		AppName:       appName,
		PGHost:        pgHost,
		PGPort:        pgPort,
		PGUsername:    pgUsername,
		PGPassword:    pgPassword,
		PGDatabase:    pgDatabase,
		SSLMode:       sslMode,
		REDISHost:     redisHost,
		REDISPort:     redisPort,
		REDISUser:     redisUser,
		REDISPassword: redisPassword,
		REDISURL:      redisURL,
		Port:          port,
		FrontendURL:   frontendURL,
	}, nil
}

func LoadEnvAndGetConfig() (*Config, error) {
	err := LoadEnv()
	if err != nil {
		return nil, err
	}

	return GetConfig()
}
