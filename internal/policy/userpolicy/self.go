package userpolicy

import (
	"fmt"
	"lms-backend/internal/policy"
	"lms-backend/internal/session"

	"github.com/gofiber/fiber/v2"
)

type IsSelf struct {
	UserID int64
}

func AllowIfIsSelf(userID int64) *IsSelf {
	return &IsSelf{userID}
}

func (p *IsSelf) Validate(c *fiber.Ctx) (policy.Decision, error) {
	userID, err := session.GetLoginSession(c)
	if err != nil {
		return policy.Deny, err
	}

	if userID != p.UserID {
		return policy.Deny, nil
	}

	return policy.Allow, nil
}

func (p *IsSelf) Reason() string {
	return fmt.Sprintf("User with ID %d is not the logged in user.", p.UserID)
}

type IsNotSelf struct {
	UserID int64
}

func AllowIfIsNotSelf(userID int64) *IsNotSelf {
	return &IsNotSelf{userID}
}

func (p *IsNotSelf) Validate(c *fiber.Ctx) (policy.Decision, error) {
	userID, err := session.GetLoginSession(c)
	if err != nil {
		return policy.Deny, err
	}

	if userID == p.UserID {
		return policy.Deny, nil
	}

	return policy.Allow, nil
}

func (p *IsNotSelf) Reason() string {
	return fmt.Sprintf("User with ID %d is the logged in user.", p.UserID)
}
