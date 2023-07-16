package database

import (
	"gorm.io/gorm"
)

func GetConfig() *gorm.Config {
	return &gorm.Config{}
}
