package main

import (
	"flag"
	"fmt"
	"math"
	"path/filepath"
	migration "technical-test/cmd/migration/helper"
	"technical-test/internal/app"
)

func main() {
	dir := flag.String("dir", "up", "Direction of migration")
	step := flag.Int("step", math.MaxInt, "Step of migration")
	flag.Parse()

	if *dir != "up" && *dir != "down" {
		panic("Direction must be up or down")
	}

	err := app.LoadEnvAndConnectToDB()
	if err != nil {
		panic(err)
	}

	//nolint:revive // ignore error
	fmt.Println("Retrieving migration Files")
	absPath, err := filepath.Abs("./migrations/")
	if err != nil {
		panic(err)
	}
	files, err := migration.ListSQLFiles(absPath)
	if err != nil {
		panic(err)
	}

	for i := range files {
		if i >= *step {
			break
		}

		var fileName string
		if *dir == "up" {
			fileName = files[i]
		} else {
			fileName = files[len(files)-1-i]
		}

		//nolint:revive // ignore error
		fmt.Printf("Running migration file %d: %s\n", i+1, fileName)
		err = migration.RunMigration(fileName, *dir)
		if err != nil {
			panic(err)
		}
	}
	//nolint:revive // ignore error
	fmt.Println("Database migrated successfully.")
}
