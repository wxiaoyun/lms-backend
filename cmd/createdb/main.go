package main

import (
	"fmt"
	"lms-backend/internal/config"
	"lms-backend/internal/database"
)

func main() {
	cf, err := config.LoadEnvAndGetConfig()
	if err != nil {
		panic(err)
	}

	//nolint:revive // ignore error
	fmt.Println("Creating database...")
	err = database.CreateDB(cf)
	if err != nil {
		panic(err)
	}
	//nolint:revive // ignore error
	fmt.Println("Successfully created database.")
}
