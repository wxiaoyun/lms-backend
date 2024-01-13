package abilities

import (
	"lms-backend/internal/model"
)

var (
	CanReadReservation model.Ability = model.Ability{
		Name:        "canReadReservation",
		Description: "can read reservation",
	}
	CanCreateReservation model.Ability = model.Ability{
		Name:        "canCreateReservation",
		Description: "can create reservation",
	}
	CanCancelReservation model.Ability = model.Ability{
		Name:        "canCancelReservation",
		Description: "can cancel reservation",
	}
	CanDeleteReservation model.Ability = model.Ability{
		Name:        "canDeleteReservation",
		Description: "can delete reservation",
	}
)
