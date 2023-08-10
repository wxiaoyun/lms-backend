package commonpolicy

import (
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/database"
	"lms-backend/internal/model"
	"lms-backend/internal/policy"
	"lms-backend/internal/session"

	"github.com/gofiber/fiber/v2"
)

type AllAbilities struct {
	Abilities []string
}

func HasAllAbilities(abilities ...string) *AllAbilities {
	return &AllAbilities{
		Abilities: abilities,
	}
}

func (a *AllAbilities) Validate(c *fiber.Ctx) (policy.Decision, error) {
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

	// Check if user has all abilities
	for _, ability := range a.Abilities {
		// If user does not have the ability or the ability does not exist
		if exist, ok := abilitesMap[ability]; !ok || !exist {
			return policy.Deny, nil
		}
	}

	return policy.Allow, nil
}

func ToAbilitiesMap(abilities []model.Ability) map[string]bool {
	abilitiesMap := map[string]bool{}

	for _, ability := range abilities {
		abilitiesMap[ability.Name] = true
	}

	return abilitiesMap
}
