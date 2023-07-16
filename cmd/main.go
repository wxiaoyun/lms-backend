package main

import (
	"auth-practice/internal/database"
	"auth-practice/internal/routes"

	"github.com/pkg/errors"
)

func main() {
	var err error

	err = database.OpenDataBase()
	if err != nil {
		panic(errors.Wrap(err, "Unavailable to open database\n"))
	}

	err = routes.SetUpRoutes()
	if err != nil {
		panic(errors.Wrap(err, "Unavailable to set up routes\n"))
	}

}
