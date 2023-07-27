package model

import (
	"regexp"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

var (
	emailReg = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

const (
	MinimumPasswordLength = 8
	MaximumPasswordLength = 32
	DefaultCost           = 10
)

func (u *User) ensureEmailIsUnique(db *gorm.DB) error {
	var exists int64

	result := db.Model(&User{}).
		Where("email = ?", u.Email).
		Count(&exists)
	if result.Error != nil {
		return result.Error
	}

	if exists > 0 {
		return fiber.NewError(fiber.StatusBadRequest, "email already exists")
	}

	return nil
}

func (u *User) Validate(db *gorm.DB) error {
	if u.Email == "" {
		return fiber.NewError(fiber.StatusBadRequest, "email is required")
	}

	if err := u.ensureEmailIsUnique(db); err != nil {
		return err
	}

	if len(u.Password) < MinimumPasswordLength {
		return fiber.NewError(fiber.StatusBadRequest, "password must be at least 8 characters")
	}

	if len(u.Password) > MaximumPasswordLength {
		return fiber.NewError(fiber.StatusBadRequest, "password must be at most 32 characters")
	}

	return nil
}

func (u *User) Create(db *gorm.DB) error {
	return db.Create(u).Error
}

func (u *User) Update(db *gorm.DB) error {
	return db.Updates(u).Error
}

func (u *User) Delete(db *gorm.DB) error {
	return db.Delete(u).Error
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	err := u.Validate(db)
	if err != nil {
		return err
	}

	// hash password
	err = u.HashPassword()
	if err != nil {
		return err
	}

	if !emailReg.MatchString(u.Email) {
		return fiber.NewError(fiber.StatusBadRequest, "invalid email")
	}

	return nil
}

func (u *User) BeforeUpdate(db *gorm.DB) error {
	if !emailReg.MatchString(u.Email) {
		return fiber.NewError(fiber.StatusBadRequest, "invalid email")
	}

	return u.Validate(db)
}

func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)

	return nil
}
