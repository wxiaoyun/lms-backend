package bookpolicy

import (
	"fmt"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/abilities"
	"lms-backend/internal/policy/commonpolicy"
)

func CreatePolicy() policy.Policy {
	fmt.Println("line 0")
	return commonpolicy.Any(
		commonpolicy.HasAnyAbility(abilities.CanManageAll.Name, abilities.CanCreateBook.Name),
	)
}
