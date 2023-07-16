package model

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

const (
	MinimumPasswordLength = 8
	DefaultCost           = 10
)

func (u *User) Validate(db *gorm.DB) error {
	if u.Email == "" {
		return fiber.NewError(fiber.StatusBadRequest, "email is required")
	}

	if len(u.Password) < MinimumPasswordLength {
		return fiber.NewError(fiber.StatusBadRequest, "password must be at least 8 characters")
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
	// hash password
	err := u.HashPassword()
	if err != nil {
		return err
	}

	return u.Validate(db)
}

func (u *User) BeforeUpdate(db *gorm.DB) error {
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
