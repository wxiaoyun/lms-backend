package main

import (
	"fmt"
	"technical-test/internal/app"
	"technical-test/internal/database"
)

func main() {
	err := app.LoadEnvAndConnectToDB()
	if err != nil {
		panic(err)
	}
	//nolint:revive // ignore error
	fmt.Println("Migrating database...")
	err = database.AutoMigration()
	if err != nil {
		panic(err)
	}
	//nolint:revive // ignore error
	fmt.Println("Database migrated successfully.")
}
