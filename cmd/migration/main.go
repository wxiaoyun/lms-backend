package main

import (
	"auth-practice/internal/config"
	"auth-practice/internal/database"
)

func main() {
	err := config.LoadENV()
	if err != nil {
		panic(err)
	}

	err = database.OpenDataBase()
	if err != nil {
		panic(err)
	}

	err = database.AutoMigration()
	if err != nil {
		panic(err)
	}
}
