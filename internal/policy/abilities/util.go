package abilities

import (
	"lms-backend/internal/model"
)

func GetAllAbilities() []model.Ability {
	return []model.Ability{
		CanManageAll,

		CanReadAuditLog,
		CanCreateAuditLog,

		CanUpdateUser,
		CanUpdateUserRole,

		CanCreatePerson,
		CanUpdatePerson,

		CanReadBook,
		CanCreateBook,
		CanUpdateBook,
		CanDeleteBook,

		CanManageBookRecords,

		CanLoanBook,
		CanReturnBook,
		CanRenewBook,
		CanReadLoan,
		CanDeleteLoan,

		CanReadFine,
		CanSettleFine,
		CanDeleteFine,

		CanReadReservation,
		CanCreateReservation,
		CanCancelReservation,
		CanCheckoutReservation,
		CanDeleteReservation,
	}
}
