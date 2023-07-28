package main

import (
	"fmt"
	"technical-test/internal/config"
	"technical-test/internal/database"
)

func main() {
	err := config.LoadENV()
	if err != nil {
		panic(err)
	}

	//nolint:revive // ignore error
	fmt.Println("Dropping database...")
	err = database.DropDB()
	if err != nil {
		panic(err)
	}
	//nolint:revive // ignore error
	fmt.Println("Database dropped successfully.")
}
