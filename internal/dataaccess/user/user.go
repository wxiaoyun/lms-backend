package user

import (
	"auth-practice/internal/model"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func VerifyLogin(db *gorm.DB, user *model.User) error {
	var userInDB model.User
	result := db.Model(&model.User{}).
		Where("email = ?", user.Email).
		First(&userInDB)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "user not found or invalid password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(userInDB.Password), []byte(user.Password))
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "user not found or invalid password")
	}

	return nil
}
