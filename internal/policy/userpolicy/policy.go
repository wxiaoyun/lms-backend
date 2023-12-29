package userpolicy

import (
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/abilities"
	"lms-backend/internal/policy/commonpolicy"
)

func ListPolicy() policy.Policy {
	return commonpolicy.Any(
		commonpolicy.HasAnyAbility(abilities.CanManageAll.Name, abilities.CanReadUser.Name),
	)
}

func ReadPolicy(userID int64) policy.Policy {
	return commonpolicy.Any(
		commonpolicy.HasAnyAbility(abilities.CanManageAll.Name, abilities.CanReadUser.Name),
		AllowIfIsSelf(userID),
	)
}

func UpdatePolicy(userID int64) policy.Policy {
	return commonpolicy.Any(
		commonpolicy.HasAnyAbility(abilities.CanManageAll.Name),
		AllowIfIsSelf(userID),
		commonpolicy.All(
			commonpolicy.HasAnyAbility(abilities.CanUpdateUser.Name),
			AllowIfSubjectBelowOwnRank(userID),
		),
	)
}

func DeletePolicy(userID int64) policy.Policy {
	return commonpolicy.Any(
		commonpolicy.HasAnyAbility(abilities.CanManageAll.Name),
		commonpolicy.All(
			commonpolicy.HasAnyAbility(abilities.CanDeleteUser.Name),
			AllowIfSubjectBelowOwnRank(userID),
		),
	)
}

func UpdateRolePolicy(userID, roleID int64) policy.Policy {
	return commonpolicy.Any(
		commonpolicy.HasAnyAbility(abilities.CanManageAll.Name),
		commonpolicy.All(
			commonpolicy.HasAnyAbility(abilities.CanUpdateUserRole.Name),
			AllowIfIsNotSelf(userID),
			AllowIfSubjectBelowOwnRank(userID),
			AllowIfPromoteBelowOwnRank(userID, roleID),
		),
	)
}
