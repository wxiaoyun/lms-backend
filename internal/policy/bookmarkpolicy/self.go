package bookmarkpolicy

import (
	"lms-backend/internal/policy"
	"lms-backend/internal/session"

	"github.com/gofiber/fiber/v2"
)

type QuerySelf struct{}

func AllowIfQuerySelf() *QuerySelf {
	return &QuerySelf{}
}

func (*QuerySelf) Validate(c *fiber.Ctx) (policy.Decision, error) {
	userID, err := session.GetLoginSession(c)
	if err != nil {
		return policy.Deny, err
	}

	queryUserID := c.QueryInt("filter[user_id]", 0)
	if int(userID) != queryUserID {
		return policy.Deny, nil
	}

	return policy.Allow, nil
}

func (*QuerySelf) Reason() string {
	return "You cannot query bookmark that is not yours."
}
