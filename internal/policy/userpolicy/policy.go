package userpolicy

import (
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/abilities"
	"lms-backend/internal/policy/commonpolicy"
)

func ReadPolicy(userID int64) policy.Policy {
	return commonpolicy.Any(
		commonpolicy.HasAnyAbility(abilities.CanManageAll.Name, abilities.CanReadUser.Name),
		AllowIfIsSelf(userID),
	)
}

func UpdatePolicy(userID int64) policy.Policy {
	return commonpolicy.Any(
		commonpolicy.HasAnyAbility(abilities.CanManageAll.Name, abilities.CanUpdateUser.Name),
		AllowIfIsSelf(userID),
	)
}

func DeletePolicy(userID int64) policy.Policy {
	return commonpolicy.Any(
		commonpolicy.HasAnyAbility(abilities.CanManageAll.Name, abilities.CanDeleteUser.Name),
		AllowIfIsSelf(userID),
	)
}

func UpdateRolePolicy(userID, roleID int64) policy.Policy {
	return commonpolicy.Any(
		commonpolicy.HasAnyAbility(abilities.CanManageAll.Name),

		commonpolicy.All(
			commonpolicy.HasAnyAbility(abilities.CanUpdateUserRole.Name),
			AllowIfIsNotSelf(userID),
			AllowIfPromoteBelowOwnRank(roleID),
		),
	)
}
