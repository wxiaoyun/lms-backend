// nolint
package main

import (
	"fmt"
	"lms-backend/cmd/seed/helper"
	"lms-backend/internal/app"
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/database"
	"lms-backend/internal/model"
	"log"
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

	fmt.Println("Seeding roles...")
	err = seedRolesAbilities(db)
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

func seedRolesAbilities(db *gorm.DB) error {
	roles := []model.Role{
		{
			Name: "System Admin",
			Abilities: []model.Ability{
				{
					Name:        "canManageAll",
					Description: "Master permission",
				},
			},
		},
		{
			Name: "Library Admin",
			Abilities: []model.Ability{
				{
					Name:        "canReadStaff",
					Description: "Read Staff",
				},
				{
					Name:        "canCreateStaff",
					Description: "Create Staff",
				},
				{
					Name:        "canUpdateStaff",
					Description: "Update Staff",
				},
				{
					Name:        "canDeleteStaff",
					Description: "Delete Staff",
				},
				{
					Name:        "canReadBook",
					Description: "Read Book",
				},
				{
					Name:        "canCreateBook",
					Description: "Create Book",
				},
				{
					Name:        "canUpdateBook",
					Description: "Update Book",
				},
				{
					Name:        "canDeleteBook",
					Description: "Delete Book",
				},
				{
					Name:        "canBorrowBook",
					Description: "Borrow Book",
				},
				{
					Name:        "canReturnBook",
					Description: "Return Book",
				},
				{
					Name:        "canExtendBook",
					Description: "Extend Book",
				},
				{
					Name:        "canManageBookRecords",
					Description: "Manage Book Records",
				},
			},
		},
		{
			Name: "Staff",
		},
		{
			Name: "Basic",
		},
		{
			Name: "Guest",
		},
	}

	// Create roles and abilities
	if err := db.Create(&roles).Error; err != nil {
		return err
	}

	_, err := user.UpdateRoles(db, 1, []int64{1})
	if err != nil {
		return err
	}

	return nil
}
