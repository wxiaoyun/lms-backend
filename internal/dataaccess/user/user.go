package user

import (
	"errors"
	"fmt"
	"lms-backend/internal/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func preloadAssociations(db *gorm.DB) *gorm.DB {
	return db.Preload("Person")
}

func Login(db *gorm.DB, user *model.User) (*model.User, error) {
	var userInDB model.User
	result := db.Model(&model.User{}).
		Scopes(preloadAssociations).
		Where("email = ?", user.Email).
		First(&userInDB)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "user not found or invalid password")
		}
		return nil, err
	}

	err := bcrypt.CompareHashAndPassword([]byte(userInDB.EncryptedPassword), []byte(user.EncryptedPassword))
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "user not found or invalid password")
	}

	userInDB.LastSignInAt = userInDB.CurrentSignInAt
	userInDB.CurrentSignInAt = time.Now()
	userInDB.SignInCount++

	return Update(db, &userInDB)
}

func Read(db *gorm.DB, id int64) (*model.User, error) {
	var user model.User
	result := db.Model(&model.User{}).
		Scopes(preloadAssociations).
		Where("id = ?", id).
		First(&user)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound,
				fmt.Sprintf("user %d not found", id),
			)
		}
		return nil, err
	}

	return &user, nil
}

func Update(db *gorm.DB, user *model.User) (*model.User, error) {
	if user.PersonID != 0 {
		if err := user.Person.Update(db); err != nil {
			return nil, err
		}
	} else {
		if err := user.Person.Create(db); err != nil {
			return nil, err
		}
	}

	if err := user.Update(db); err != nil {
		return nil, err
	}

	return Read(db, int64(user.ID))
}

func ReadByEmail(db *gorm.DB, email string) (*model.User, error) {
	var user model.User
	result := db.Model(&model.User{}).
		Scopes(preloadAssociations).
		Where("email = ?", email).
		First(&user)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound,
				fmt.Sprintf("user %s not found", email),
			)
		}
		return nil, err
	}

	return &user, nil
}
