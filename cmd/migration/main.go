package main

import (
	"flag"
	"fmt"
	"lms-backend/internal/app"
	"lms-backend/internal/config"
	"lms-backend/internal/database"
	"lms-backend/util/ternary"
	"log"
	"math"
	"path/filepath"

	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	dir := flag.String("dir", "up", "Direction of migration")
	step := flag.Int("step", math.MaxInt, "Step of migration")
	flag.Parse()

	if *dir != "up" && *dir != "down" {
		log.Fatal("Direction must be up or down")
	}

	err := app.LoadEnvAndConnectToDB()
	if err != nil {
		log.Fatal(err)
	}

	absPath, err := filepath.Abs("./migrations/")
	if err != nil {
		log.Fatal(err)
	}

	cf, err := config.LoadEnvAndGetConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.ConnectToDB(cf)
	if err != nil {
		log.Fatal(err)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: absPath,
	}

	n, err := migrate.ExecMax(
		db,
		"postgres",
		migrations,
		ternary.If[migrate.MigrationDirection](*dir == "up").
			Then(migrate.Up).
			Else(migrate.Down),
		*step,
	)
	if err != nil {
		log.Fatal(err)
	}

	//nolint
	fmt.Printf("Applied %d migrations!\n", n)
}
