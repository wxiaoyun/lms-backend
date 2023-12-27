package bookmarkpolicy

import (
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/abilities"
	"lms-backend/internal/policy/commonpolicy"
)

func ListPolicy() policy.Policy {
	return commonpolicy.Any(
		commonpolicy.HasAnyAbility(
			abilities.CanManageAll.Name,
			abilities.CanReadBookMark.Name,
		),
		AllowIfQuerySelf(),
	)
}

func CreatePolicy() policy.Policy {
	return commonpolicy.Any(
		commonpolicy.HasAnyAbility(
			abilities.CanManageAll.Name,
			abilities.CanCreateBookMark.Name,
		),
	)
}

func DeletePolicy(bookmarkID int64) policy.Policy {
	return commonpolicy.Any(
		commonpolicy.HasAnyAbility(
			abilities.CanManageAll.Name,
			abilities.CanDeleteBookMark.Name,
		),
		AllowIfBookmarkBelongsToUser(bookmarkID),
	)
}
