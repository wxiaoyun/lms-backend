package model

import (
	"technical-test/internal/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestUserValidationNegative(t *testing.T) {
	db, err := testutil.ConnectToDB()
	if err != nil {
		t.Errorf("failed to load env and connect to db: %v", err)
		return
	}

	db = db.Begin()

	t.Run("Email is empty", func(t *testing.T) {
		user := &User{Email: "", Password: "Password@123"}
		err := user.Validate(db)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "email is required")
	})

	t.Run("Email is invalid", func(t *testing.T) {
		user := &User{Email: "not_an_email", Password: "Password@123"}
		err := user.Validate(db)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid email")
	})

	t.Run("Email already exists", func(t *testing.T) {
		existingUser := &User{Email: "test@test.com", Password: "Password@123"}
		err := db.Create(existingUser).Error
		assert.NoError(t, err)

		newUser := &User{Email: "test@test.com", Password: "Password@123"}
		err = newUser.Validate(db)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "email already exists")
	})

	t.Run("Password is too short", func(t *testing.T) {
		user := &User{Email: "test2@test.com", Password: "Short1!"}
		err := user.Validate(db)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "password must be at least")
	})

	t.Run("Password is too long", func(t *testing.T) {
		longPassword := "Long1!" + string(make([]byte, MaximumPasswordLength))
		user := &User{Email: "test3@test.com", Password: longPassword}
		err := user.Validate(db)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "password must be at most")
	})

	t.Run("Password does not meet complexity requirements", func(t *testing.T) {
		user := &User{Email: "test4@test.com", Password: "simplepassword"}
		err := user.Validate(db)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "password must contain")
	})

	db.Rollback()
}

func TestUserValidationPositive(t *testing.T) {
	// Create a new GORM connection with an in-memory sqlite database.
	db, err := testutil.ConnectToDB()
	if err != nil {
		t.Errorf("failed to load env and connect to db: %v", err)
		return
	}

	db = db.Begin()

	t.Run("Valid User", func(t *testing.T) {
		user := &User{Email: "test5@test.com", Password: "Password@123"}
		err := user.Validate(db)
		assert.NoError(t, err)
	})

	t.Run("Create Valid User", func(t *testing.T) {
		user := &User{Email: "test6@test.com", Password: "Password@123"}
		passwordBeforeHash := user.Password
		err := user.Create(db)
		assert.NoError(t, err)

		var dbUser User
		err = db.Where("email = ?", user.Email).First(&dbUser).Error
		assert.NoError(t, err)

		// Check if password is hashed
		err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(passwordBeforeHash))
		assert.NoError(t, err)
	})

	db.Rollback()
}
