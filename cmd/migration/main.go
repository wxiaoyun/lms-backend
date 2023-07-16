package main

import (
	"auth-practice/internal/database"
)

func main() {
	err := database.OpenDataBase()
	if err != nil {
		panic(err)
	}
	err = database.AutoMigration()
	if err != nil {
		panic(err)
	}
}
