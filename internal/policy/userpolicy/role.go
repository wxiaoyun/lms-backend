package userpolicy

import (
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/database"
	"lms-backend/internal/policy"
	"lms-backend/internal/session"

	"github.com/gofiber/fiber/v2"
)

type PromoteBelowOwnRank struct {
	userID int64
	RoleID int64
}

func AllowIfPromoteBelowOwnRank(userID, roleID int64) *PromoteBelowOwnRank {
	return &PromoteBelowOwnRank{userID, roleID}
}

func (p *PromoteBelowOwnRank) Validate(c *fiber.Ctx) (policy.Decision, error) {
	userID, err := session.GetLoginSession(c)
	if err != nil {
		return policy.Deny, err
	}

	db := database.GetDB()

	promoterRoles, err := user.GetRoles(db, userID)
	if err != nil {
		return policy.Deny, err
	}

	if len(promoterRoles) == 0 {
		return policy.Deny, nil
	}

	// Roles are ordered by rank in descending order. The first role is the highest rank.
	// Higher rank means lower ID number.
	if int64(promoterRoles[0].ID) >= p.RoleID {
		return policy.Deny, nil
	}

	return policy.Allow, nil
}

func (*PromoteBelowOwnRank) Reason() string {
	return "You are not allowed to promote users beyond your own rank."
}
