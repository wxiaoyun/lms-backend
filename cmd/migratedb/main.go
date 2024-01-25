package main

import (
	"flag"
	"fmt"
	migratedb "lms-backend/cmd/migratedb/migrate"
	"lms-backend/internal/app"
	"lms-backend/internal/config"
	"lms-backend/internal/database"
	"log"
	"math"
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

	cfg, err := config.LoadEnvAndGetConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.ConnectToDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	n, err := migratedb.Migrate(db, *dir, *step)
	if err != nil {
		log.Fatal(err)
	}

	//nolint
	fmt.Printf("Applied %d migrations!\n", n)
}
