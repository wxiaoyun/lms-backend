package main

import (
	"technical-test/internal/config"
	"technical-test/internal/database"
)

func main() {
	err := config.LoadENV()
	if err != nil {
		panic(err)
	}

	err = database.DropDB()
	if err != nil {
		panic(err)
	}
}
