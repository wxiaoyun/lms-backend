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
	fmt.Println("Creating database...")
	err = database.CreateDB()
	if err != nil {
		panic(err)
	}
	//nolint:revive // ignore error
	fmt.Println("Successfully created database.")
}
