// nolint
package main

import (
	"fmt"
	"log"
	"technical-test/cmd/seed/helper"
	"technical-test/internal/app"
	"technical-test/internal/database"
	"technical-test/internal/model"
	"time"

	"github.com/go-loremipsum/loremipsum"
	"gorm.io/gorm"
)

func main() {
	var err error

	// Load environment variables and connect to database
	err = app.LoadEnvAndConnectToDB()
	if err != nil {
		log.Fatal(err)
	}

	db := database.GetDB()

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil || err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	fmt.Println("Seeding database...")
	fmt.Println("Seeding users and people...")
	err = seedUsersAndPeople(db)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Seeding books...")
	err = seedBooks(db)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Seeding complete!")
}

func seedUsersAndPeople(db *gorm.DB) error {
	var count int64

	result := db.Model(&model.User{}).Count(&count)
	if result.Error != nil {
		return result.Error
	}

	if count >= 100 {
		return nil
	}

	loremIpsumGenerator := loremipsum.New()

	// Generate 100 users
	users := make([]model.User, 100)
	for i := 1; i <= 100; i++ {
		users[i-1] = model.User{
			Email:             fmt.Sprintf("user%d@gmail.com", i),
			EncryptedPassword: "P4ssw0rd!",
			Person: &model.Person{
				FirstName: loremIpsumGenerator.Word(),
				LastName:  loremIpsumGenerator.Word(),
			},
		}
	}

	return db.Create(&users).Error
}

func seedBooks(db *gorm.DB) error {
	var count int64

	result := db.Model(&model.Book{}).Count(&count)
	if result.Error != nil {
		return result.Error
	}

	if count >= 1000 {
		return nil
	}

	loremIpsumGenerator := loremipsum.New()

	books := make([]model.Book, 3000)
	for i := 1; i <= 3000; i++ {
		books[i-1] = model.Book{
			Title:           loremIpsumGenerator.Words(helper.RandInt(4, 11)),
			Author:          loremIpsumGenerator.Words(helper.RandInt(2, 5)),
			ISBN:            helper.GenerateISBN13(),
			Publisher:       loremIpsumGenerator.Words(helper.RandInt(4, 7)),
			PublicationDate: helper.RandomDate(time.Now().AddDate(-10, 0, 0), time.Now()),
			Genre:           loremIpsumGenerator.Words(helper.RandInt(1, 3)),
			Language:        loremIpsumGenerator.Words(helper.RandInt(1, 3)),
		}
	}

	return db.Create(&books).Error
}
