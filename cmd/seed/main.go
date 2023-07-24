package main

import (
	"fmt"
	"math/rand"
	"technical-test/internal/app"
	"technical-test/internal/dataaccess/user"
	"technical-test/internal/database"
	"technical-test/internal/model"

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
	fmt.Println("Seeding work sheets...")
	err = seedWorkSheets(db, user1)
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
	workSheets := make([]model.Worksheet, 1000)
	for i := 0; i < 1000; i++ {
		workSheets[i] = model.Worksheet{
			Title:       fmt.Sprintf("Title - %d", i+1),
			Description: fmt.Sprintf("Description - %d", i+1),
			//nolint:gosec // cost does not need to be secure
			Cost: rand.Float64() * 10,
			//nolint:gosec // cost does not need to be secure
			Price:  rand.Float64() * 10,
			UserID: user1.ID,
		}
	}

	return db.Create(&workSheets).Error
}
