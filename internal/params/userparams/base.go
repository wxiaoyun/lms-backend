package userparams

import (
	"regexp"
	"technical-test/internal/model"

	"github.com/gofiber/fiber/v2"
)

var (
	emailReg = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

type BaseUserParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (b *BaseUserParams) ToModel() *model.User {
	return &model.User{
		Email:    b.Email,
		Password: b.Password,
	}
}

func (b *BaseUserParams) Validate() error {
	if !emailReg.MatchString(b.Email) {
		return fiber.NewError(fiber.StatusBadRequest, "invalid email")
	}

	return nil
}
