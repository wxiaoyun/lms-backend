package commonpolicy

import (
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/database"
	"lms-backend/internal/policy"
	"lms-backend/internal/session"

	"github.com/gofiber/fiber/v2"
)

type AnyAbility struct {
	Abilities []string
}

func HasAnyAbility(abilities ...string) *AnyAbility {
	return &AnyAbility{
		Abilities: abilities,
	}
}

func (a *AnyAbility) Validate(c *fiber.Ctx) (policy.Decision, error) {
	userID, err := session.GetLoginSession(c)
	if err != nil {
		return policy.Deny, err
	}

	db := database.GetDB()

	abilites, err := user.GetAbilities(db, userID)
	if err != nil {
		return policy.Deny, err
	}

	abilitesMap := ToAbilitiesMap(abilites)

	// Check if user has any abilities
	for _, ability := range a.Abilities {
		// If user has the ability
		if exist, ok := abilitesMap[ability]; ok && exist {
			return policy.Allow, nil
		}
	}

	return policy.Deny, nil
}
