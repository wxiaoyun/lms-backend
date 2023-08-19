package main

import (
	cdbhelper "lms-backend/cmd/createdb/helper"
	"lms-backend/internal/config"
	logger "lms-backend/internal/log"
)

var lgr = logger.StdoutLogger()

func main() {
	cf, err := config.LoadEnvAndGetConfig()
	if err != nil {
		lgr.Println(err)
		panic(err)
	}

	if err := cdbhelper.CreateDB(cf); err != nil {
		lgr.Println(err)
		panic(err)
	}

	lgr.Println("Successfully created database.")
}
