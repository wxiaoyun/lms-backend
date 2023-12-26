package database

import (
	"lms-backend/internal/config"

	"github.com/gofiber/storage/redis/v3"
)

var redisStore *redis.Storage

func SetupRedis(cfg *config.Config) {
	redisStore = redis.New(redis.Config{
		Host:     cfg.REDISHost,
		Port:     cfg.REDISPort,
		Username: cfg.REDISUser,
		Password: cfg.REDISPassword,
		URL:      cfg.REDISURL,
		Reset:    true,
	})
}

func GetRedisStore() *redis.Storage {
	return redisStore
}
