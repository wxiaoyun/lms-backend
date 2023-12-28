package main

import (
	shelper "lms-backend/cmd/seed/helper"
	"lms-backend/internal/app"
	"lms-backend/internal/database"
	logger "lms-backend/internal/log"
	"log"
)

const (
	NumberOfUsers = 10
	NumberOfBooks = 3000
)

func main() {
	var err error
	lgr := logger.StdoutLogger()

	// Load environment variables and connect to database
	err = app.LoadEnvAndConnectToDB()
	if err != nil {
		log.Fatal(err)
	}

	db := database.GetDB()

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil || err != nil {
			lgr.Println(r)
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	lgr.Println("Seeding database...")

	lgr.Println("Seeding roles and abilities...")
	err = shelper.SeedRoleAndAbility(tx)
	if err != nil {
		panic(err)
	}

	lgr.Println("Linking roles with abilities...")
	err = shelper.LinkRoleAndAbility(tx)
	if err != nil {
		panic(err)
	}

	lgr.Println("Seeding users and people...")
	err = shelper.SeedUsersAndPeople(tx, NumberOfUsers)
	if err != nil {
		panic(err)
	}

	lgr.Println("Seeding books and copies...")
	err = shelper.SeedBookAndCopies(tx, NumberOfBooks)
	if err != nil {
		panic(err)
	}

	lgr.Println("Linking user with roles...")
	err = shelper.LinkUserWithRoles(tx)
	if err != nil {
		panic(err)
	}

	lgr.Println("Seeding loans and reservations...")
	err = shelper.SeedLoanAndReservations(tx)
	if err != nil {
		panic(err)
	}

	lgr.Println("Seeding complete!")
}
