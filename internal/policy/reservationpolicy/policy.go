package reservationpolicy

import (
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/abilities"
	"lms-backend/internal/policy/commonpolicy"
)

func ReservePolicy() policy.Policy {
	return commonpolicy.Any(
		commonpolicy.HasAnyAbility(
			abilities.CanManageAll.Name,
			abilities.CanManageBookRecords.Name,
			abilities.CanCreateReservation.Name,
		),
	)
}

func CheckoutPolicy() policy.Policy {
	return commonpolicy.Any(
		commonpolicy.HasAnyAbility(
			abilities.CanManageAll.Name,
			abilities.CanManageBookRecords.Name,
			abilities.CanCheckoutReservation.Name,
		),
	)
}

func CancelPolicy() policy.Policy {
	return commonpolicy.Any(
		commonpolicy.HasAnyAbility(
			abilities.CanManageAll.Name,
			abilities.CanManageBookRecords.Name,
			abilities.CanCancelReservation.Name,
		),
	)
}

func ManageReservationPolicy() policy.Policy {
	return commonpolicy.Any(
		commonpolicy.HasAnyAbility(abilities.CanManageAll.Name, abilities.CanManageBookRecords.Name),
	)
}
