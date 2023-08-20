package commonpolicy

import (
	"fmt"
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/database"
	"lms-backend/internal/policy"
	"lms-backend/internal/session"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AnyAbility struct {
	Abilities []string
	ReasonStr string
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
	builder := strings.Builder{}
	//nolint
	builder.WriteString("Missing abilities: ")

	// Check if user has any abilities
	for i, ability := range a.Abilities {
		// If user has the ability
		if exist, ok := abilitesMap[ability]; ok && exist {
			return policy.Allow, nil
		}

		if i == len(a.Abilities)-1 {
			//nolint
			builder.WriteString(fmt.Sprintf("%s.", ability))
			continue
		}

		//nolint
		builder.WriteString(fmt.Sprintf("%s, ", ability))
	}

	a.ReasonStr = builder.String()
	return policy.Deny, nil
}

func (a *AnyAbility) Reason() string {
	return "You don't have any of the required abilities. " + a.ReasonStr
}
