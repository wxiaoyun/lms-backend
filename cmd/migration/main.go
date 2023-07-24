package main

import (
	"technical-test/internal/app"
	"technical-test/internal/database"
)

func main() {
	err := app.LoadEnvAndConnectToDB()
	if err != nil {
		panic(err)
	}

	err = database.AutoMigration()
	if err != nil {
		panic(err)
	}
}
