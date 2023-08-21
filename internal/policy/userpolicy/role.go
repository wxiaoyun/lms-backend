package userpolicy

import (
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/database"
	"lms-backend/internal/policy"
	"lms-backend/internal/session"

	"github.com/gofiber/fiber/v2"
)

type PromoteBelowOwnRank struct {
	RoleID int64
}

func AllowIfPromoteBelowOwnRank(roleID int64) *PromoteBelowOwnRank {
	return &PromoteBelowOwnRank{roleID}
}

func (p *PromoteBelowOwnRank) Validate(c *fiber.Ctx) (policy.Decision, error) {
	userID, err := session.GetLoginSession(c)
	if err != nil {
		return policy.Deny, err
	}

	db := database.GetDB()

	roles, err := user.GetRoles(db, userID)
	if err != nil {
		return policy.Deny, err
	}

	if len(roles) == 0 {
		return policy.Deny, nil
	}

	// Roles are ordered by rank in descending order. The first role is the highest rank.
	// Higher rank means lower ID number.
	if int64(roles[0].ID) <= p.RoleID {
		return policy.Deny, nil
	}

	return policy.Allow, nil
}

func (*PromoteBelowOwnRank) Reason() string {
	return "You are not allowed to promote users beyond your own rank."
}
