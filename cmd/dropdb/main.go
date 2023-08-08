package main

import (
	"fmt"
	"technical-test/internal/config"
	"technical-test/internal/database"
)

func main() {
	cf, err := config.LoadEnvAndGetConfig()
	if err != nil {
		panic(err)
	}

	//nolint:revive // ignore error
	fmt.Println("Dropping database...")
	err = database.DropDB(cf)
	if err != nil {
		panic(err)
	}
	//nolint:revive // ignore error
	fmt.Println("Database dropped successfully.")
}
