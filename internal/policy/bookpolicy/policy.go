package bookpolicy

import (
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/abilities"
	"lms-backend/internal/policy/commonpolicy"
)

func CreatePolicy() policy.Policy {
	return commonpolicy.Any(
		commonpolicy.HasAnyAbility(abilities.CanManageAll.Name, abilities.CanCreateBook.Name),
	)
}
