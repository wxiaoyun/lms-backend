package main

import (
	"fmt"
	"math/rand"
	"technical-test/internal/app"
	"technical-test/internal/dataaccess/user"
	"technical-test/internal/database"
	"technical-test/internal/model"

	"github.com/go-loremipsum/loremipsum"

	"gorm.io/gorm"
)

func main() {
	var err error

	// Load environment variables and connect to database
	err = app.LoadEnvAndConnectToDB()
	if err != nil {
		panic(err)
	}

	db := database.GetDB()

	//nolint:revive // ignore error
	fmt.Println("Seeding database...")

	//nolint:revive // ignore error
	fmt.Println("Seeding users...")
	user1, err := seedUser(db)
	if err != nil {
		panic(err)
	}

	//nolint:revive // ignore error
	fmt.Println("Seeding worksheets...")
	err = seedWorkSheets(db, user1)
	if err != nil {
		panic(err)
	}

	//nolint:revive // ignore error
	fmt.Println("Seeding questions...")
	err = seedQuestions(db)
	if err != nil {
		panic(err)
	}

	//nolint:revive // ignore error
	fmt.Println("Seeding complete!")
}

func seedUser(db *gorm.DB) (*model.User, error) {
	user1 := model.User{
		Email:    "admin@gmail.com",
		Password: "password",
	}

	var exists int64

	result := db.Model(&model.User{}).
		Where("email = ?", user1.Email).
		Count(&exists)
	if result.Error != nil {
		return nil, result.Error
	}

	if exists == 0 {
		if err := user1.Create(db); err != nil {
			return nil, result.Error
		}
		return &user1, nil
	}

	userPtr, err := user.ReadByEmail(db, user1.Email)
	if err != nil {
		return nil, err
	}

	return userPtr, nil
}

func seedWorkSheets(db *gorm.DB, user1 *model.User) error {
	var count int64

	result := db.Model(&model.Worksheet{}).Count(&count)
	if result.Error != nil {
		return result.Error
	}

	if count >= 1000 {
		return nil
	}

	loremIpsumGenerator := loremipsum.New()

	workSheets := make([]model.Worksheet, 1000)
	for i := 1; i <= 1000; i++ {
		workSheets[i-1] = model.Worksheet{
			//nolint:gosec // title does not need to be secure
			Title: loremIpsumGenerator.Words(rand.Intn(10) + 1),
			//nolint:gosec // description does not need to be secure
			Description: loremIpsumGenerator.Paragraphs(rand.Intn(10) + 1),
			//nolint:gosec // cost does not need to be secure
			Cost: rand.Float64() * 10,
			//nolint:gosec // cost does not need to be secure
			Price:  rand.Float64() * 10,
			UserID: user1.ID,
		}
	}

	return db.Create(&workSheets).Error
}

func seedQuestions(db *gorm.DB) error {
	var count int64

	result := db.Model(&model.Question{}).Count(&count)
	if result.Error != nil {
		return result.Error
	}

	if count >= 1000 {
		return nil
	}

	questions := make([]model.Question, 3000)
	for i := 1; i <= 3000; i++ {
		questions[i-1] = model.Question{
			Description: fmt.Sprintf("Description - %d", i),
			Answer:      fmt.Sprintf("Answer - %d", i),
			//nolint:gosec // cost does not need to be secure
			Cost: rand.Float64() * 10,
			//nolint:gosec // cost does not need to be secure
			WorksheetID: int64(rand.Intn(1000) + 1),
		}
	}

	return db.Create(&questions).Error
}
