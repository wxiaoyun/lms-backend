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
	BackendURL   string
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
		appName = "LMS"
	}

	pgHost := os.Getenv("PG_HOST")
	if pgHost == "" {
		pgHost = "localhost"
	}

	pgPort := os.Getenv("PG_PORT")
	if pgPort == "" {
		pgPort = "5432"
	}

	pgUsername := os.Getenv("PG_USERNAME")
	if pgUsername == "" {
		pgUsername = "postgres"
	}

	pgPassword := os.Getenv("PG_PASSWORD")
	if pgPassword == "" {
		pgPassword = "1234"
	}

	pgDatabase := os.Getenv("PG_DATABASE")
	if pgDatabase == "" {
		pgDatabase = "lms-database"
	}

	sslMode := os.Getenv("SSL_MODE")
	if sslMode == "" {
		sslMode = "disable"
	}

	var redisHost, redisPassword, redisUser string
	var redisPort int
	redisURL := os.Getenv("REDIS_URL")

	// If REDIS_URL is set, all other config options have no effect
	if redisURL == "" {
		redisHost = os.Getenv("REDIS_HOST")
		if redisHost == "" {
			redisHost = "localhost"
		}

		rp := os.Getenv("REDIS_PORT")

		if rp != "" {
			port, err := strconv.Atoi(rp)
			if err != nil {
				return nil, internalerror.InternalServerError("Bad Redis Port: " + rp)
			}
			redisPort = port
		} else {
			redisPort = 6379
		}

		redisUser = os.Getenv("REDIS_USER")

		redisPassword = os.Getenv("REDIS_PASSWORD")
	}

	GoogleAPIKey = os.Getenv("GOOGLE_API_KEY")
	if GoogleAPIKey == "" {
		return nil, internalerror.InternalServerError("GOOGLE_API_KEY not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173"
	}

	BackendURL = os.Getenv("BACKEND_URL")
	if BackendURL == "" {
		BackendURL = "http://localhost:3000"
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
