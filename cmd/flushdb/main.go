package main

import (
	"fmt"
	"lms-backend/internal/config"
	"lms-backend/internal/database"
	"log"
)

func main() {
	cf, err := config.LoadEnvAndGetConfig()
	if err != nil {
		panic(err)
	}

	//nolint:revive // ignore error
	fmt.Println("Dropping all tables...")
	err = database.DropAllTables(cf)
	if err != nil {
		//nolint:revive // ignore lint
		log.Fatalf("Failed to drop all tables in database : %v\n", err)
	}
	//nolint:revive // ignore error
	fmt.Println("All tables dropped successfully.")
}
