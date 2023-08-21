package model

import (
	"lms-backend/internal/orm"
	"lms-backend/pkg/error/externalerrors"
	"regexp"
	"time"
	"unicode/utf8"

	"github.com/dlclark/regexp2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username          string `gorm:"unique;not null"`
	Email             string `gorm:"unique"`
	EncryptedPassword string `gorm:"not null"`
	SignInCount       int    `gorm:"not null;default:0"`
	CurrentSignInAt   time.Time
	LastSignInAt      time.Time
	PersonID          uint    `gorm:"not null"`
	Person            *Person `gorm:"->;<-:create"`
	Roles             []Role  `gorm:"many2many:user_roles;->"`
	Loans             []Loan  `gorm:"->"`
}

var (
	emailReg    = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	passwordReg = regexp2.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%^&*]).{8,32}$`, regexp2.None)
)

const (
	MinimumUsernameLength = 5
	MaximumUsernameLength = 30
	MinimumPasswordLength = 8
	MaximumPasswordLength = 32
	DefaultCost           = 10
)

const (
	UserModelName = "user"
	UserTableName = "users"
)

func (u *User) ensureUsernameIsNew(db *gorm.DB) error {
	if u.Username == "" {
		return externalerrors.BadRequest("username is required")
	}

	var exists int64

	result := db.Model(&User{}).
		Where("username = ?", u.Username).
		Count(&exists)
	if err := result.Error; err != nil {
		if !orm.IsRecordNotFound(err) {
			return err
		}
	}

	if exists > 0 {
		return externalerrors.BadRequest("username already exists")
	}

	return nil
}

func (u *User) ensureEmailIsUniqueIfPresent(db *gorm.DB) error {
	if u.Email == "" {
		return nil
	}

	var exists int64

	result := db.Model(&User{}).
		Where("email = ?", u.Email).
		Count(&exists)
	if err := result.Error; err != nil {
		if !orm.IsRecordNotFound(err) {
			return err
		}
	}

	if exists > 0 {
		return externalerrors.BadRequest("email already exists")
	}

	return nil
}

func (u *User) ensurePersonIsNewOrExists(db *gorm.DB) error {
	if u.PersonID == 0 {
		return nil
	}

	if u.PersonID != u.Person.ID {
		return externalerrors.BadRequest("person id does not match person")
	}

	var exists int64

	result := db.Model(&Person{}).
		Where("id = ?", u.PersonID).
		Count(&exists)
	if result.Error != nil {
		return externalerrors.BadRequest("person not found")
	}

	if exists != 1 {
		return externalerrors.BadRequest("person does not exist")
	}

	return nil
}

func (u *User) ValidateUnencryptedPassword() error {
	if len(u.EncryptedPassword) < MinimumPasswordLength {
		return externalerrors.BadRequest("password must be at least 8 characters")
	}

	if len(u.EncryptedPassword) > MaximumPasswordLength {
		return externalerrors.BadRequest("password must be at most 32 characters")
	}

	if ok, err := passwordReg.MatchString(u.EncryptedPassword); !ok || err != nil {
		return externalerrors.BadRequest(
			"password must contain at least one lowercase letter, " +
				"one uppercase letter, one digit, and one special character",
		)
	}

	return nil
}

func (u *User) ValidateUsername(db *gorm.DB) error {
	// counting the utf8 character length instead of byte length
	if utf8.RuneCountInString(u.Username) < MinimumUsernameLength {
		return externalerrors.BadRequest("username must be at least 5 characters")
	}

	// counting the utf8 character length instead of byte length
	if utf8.RuneCountInString(u.Username) > MaximumUsernameLength {
		return externalerrors.BadRequest("username must be at most 30 characters")
	}

	// New user
	if u.ID == 0 {
		return u.ensureUsernameIsNew(db)
	}

	// Updating user
	var originalUser User
	result := db.Model(&User{}).
		Where("id = ?", u.ID).
		First(&originalUser)
	if err := result.Error; err != nil {
		return err
	}

	if u.Username != originalUser.Username {
		if err := u.ensureUsernameIsNew(db); err != nil {
			return err
		}
	}

	return nil
}

func (u *User) ValidateEmail(db *gorm.DB) error {
	if u.Email != "" && !emailReg.MatchString(u.Email) {
		return externalerrors.BadRequest("invalid email")
	}

	// New user
	if u.ID == 0 {
		return u.ensureEmailIsUniqueIfPresent(db)
	}

	// Updating user
	var originalUser User
	result := db.Model(&User{}).
		Where("id = ?", u.ID).
		First(&originalUser)
	if err := result.Error; err != nil {
		return err
	}

	if u.Email != originalUser.Email {
		if err := u.ensureEmailIsUniqueIfPresent(db); err != nil {
			return err
		}
	}

	return nil
}

func (u *User) Validate(db *gorm.DB) error {
	return u.ensurePersonIsNewOrExists(db)
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
	if err := u.ValidateUnencryptedPassword(); err != nil {
		return err
	}

	if err := u.HashPassword(); err != nil {
		return err
	}

	if err := u.ValidateUsername(db); err != nil {
		return err
	}

	if err := u.ValidateEmail(db); err != nil {
		return err
	}

	return u.Validate(db)
}

func (u *User) AfterCreate(db *gorm.DB) error {
	// Default role is "Basic" upon creation
	return u.UpdateRoles(db, []int64{4})
}

func (u *User) BeforeUpdate(db *gorm.DB) error {
	if err := u.ValidateUsername(db); err != nil {
		return err
	}

	if err := u.ValidateEmail(db); err != nil {
		return err
	}

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

func (u *User) UpdateRoles(db *gorm.DB, roleIDs []int64) error {
	// Remove all existing roles
	result := db.
		Model(&UserRoles{}).
		Delete("user_id = ?", u.ID)
	if result.Error != nil {
		return result.Error
	}

	// Add new roles
	for _, roleID := range roleIDs {
		err := db.Exec(
			"INSERT INTO user_roles (user_id, role_id) VALUES (?, ?)",
			u.ID, roleID,
		).Error
		if err != nil {
			return err
		}
	}

	return nil
}
