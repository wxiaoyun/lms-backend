package reservationpolicy

import (
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/abilities"
	"lms-backend/internal/policy/commonpolicy"
)

func ReadPolicy() policy.Policy {
	return commonpolicy.Any(
		commonpolicy.HasAnyAbility(
			abilities.CanManageAll.Name,
			abilities.CanManageBookRecords.Name,
			abilities.CanReadReservation.Name,
		),
	)
}

func ReadBookPolicy() policy.Policy {
	return commonpolicy.Any(
		commonpolicy.HasAnyAbility(
			abilities.CanManageAll.Name,
			abilities.CanManageBookRecords.Name,
		),
		AllowIfLoanSelf(),
	)
}

func DeletePolicy() policy.Policy {
	return commonpolicy.Any(
		commonpolicy.HasAnyAbility(
			abilities.CanManageAll.Name,
			abilities.CanManageBookRecords.Name,
			abilities.CanDeleteReservation.Name,
		),
	)
}

func ReservePolicy() policy.Policy {
	return commonpolicy.Any(
		commonpolicy.HasAnyAbility(
			abilities.CanCreateReservation.Name,
		),
	)
}

func CreatePolicy() policy.Policy {
	return commonpolicy.Any(
		commonpolicy.HasAnyAbility(
			abilities.CanManageAll.Name,
			abilities.CanManageBookRecords.Name,
			abilities.CanCreateReservation.Name,
		),
	)
}

func CheckoutPolicy(resID int64) policy.Policy {
	return commonpolicy.Any(
		commonpolicy.HasAnyAbility(
			abilities.CanManageAll.Name,
			abilities.CanManageBookRecords.Name,
			abilities.CanCheckoutReservation.Name,
		),
		AllowIfReservationBelongsToUser(resID),
	)
}

func CancelPolicy(resID int64) policy.Policy {
	return commonpolicy.Any(
		commonpolicy.HasAnyAbility(
			abilities.CanManageAll.Name,
			abilities.CanManageBookRecords.Name,
			abilities.CanCancelReservation.Name,
		),
		AllowIfReservationBelongsToUser(resID),
	)
}
