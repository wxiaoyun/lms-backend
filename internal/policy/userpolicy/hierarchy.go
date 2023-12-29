package userpolicy

import (
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/database"
	"lms-backend/internal/policy"
	"lms-backend/internal/session"

	"github.com/gofiber/fiber/v2"
)

type SubjectBelowOwnRank struct {
	userID int64
}

func AllowIfSubjectBelowOwnRank(userID int64) *SubjectBelowOwnRank {
	return &SubjectBelowOwnRank{userID}
}

func (p *SubjectBelowOwnRank) Validate(c *fiber.Ctx) (policy.Decision, error) {
	currentID, err := session.GetLoginSession(c)
	if err != nil {
		return policy.Deny, err
	}

	db := database.GetDB()

	currentRoles, err := user.GetRoles(db, currentID)
	if err != nil {
		return policy.Deny, err
	}

	if len(currentRoles) == 0 {
		return policy.Deny, nil
	}

	subjectRoles, err := user.GetRoles(db, p.userID)
	if err != nil {
		return policy.Deny, err
	}

	if len(subjectRoles) == 0 {
		return policy.Deny, nil
	}

	// Roles are ordered by rank in descending order. The first role is the highest rank.
	// Higher rank means lower ID number.
	if int64(currentRoles[0].ID) >= int64(subjectRoles[0].ID) {
		return policy.Deny, nil
	}

	return policy.Allow, nil
}

func (*SubjectBelowOwnRank) Reason() string {
	return "You are not allowed to perform actions on users above your own rank."
}
