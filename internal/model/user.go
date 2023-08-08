package model

import (
	"regexp"
	"time"

	"github.com/dlclark/regexp2"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Email              string `gorm:"unique;not null"`
	EncryptedPassword  string `gorm:"not null"`
	CurrentSignInCount int    `gorm:"not null;default:0"`
	CurrentSignIn      time.Time
	LastSignInAt       time.Time
}

var (
	emailReg    = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	passwordReg = regexp2.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%^&*]).{8,32}$`, regexp2.None)
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
	if len(u.EncryptedPassword) < MinimumPasswordLength {
		return fiber.NewError(fiber.StatusBadRequest, "password must be at least 8 characters")
	}

	if len(u.EncryptedPassword) > MaximumPasswordLength {
		return fiber.NewError(fiber.StatusBadRequest, "password must be at most 32 characters")
	}

	if ok, err := passwordReg.MatchString(u.EncryptedPassword); !ok || err != nil {
		return fiber.NewError(fiber.StatusBadRequest,
			"password must contain at least one lowercase letter, "+
				"one uppercase letter, one digit, and one special character",
		)
	}

	if u.Email == "" {
		return fiber.NewError(fiber.StatusBadRequest, "email is required")
	}

	if !emailReg.MatchString(u.Email) {
		return fiber.NewError(fiber.StatusBadRequest, "invalid email")
	}

	return u.ensureEmailIsUnique(db)
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

	return u.HashPassword()
}

func (u *User) BeforeUpdate(db *gorm.DB) error {
	return u.Validate(db)
}

func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.EncryptedPassword), DefaultCost)
	if err != nil {
		return err
	}
	u.EncryptedPassword = string(bytes)

	return nil
}
