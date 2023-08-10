package abilities

import (
	"lms-backend/internal/model"
)

var (
	CanCreateBook model.Ability = model.Ability{
		Name:        "canCreateBook",
		Description: "can create book",
	}
	CanReadBook model.Ability = model.Ability{
		Name:        "canReadBook",
		Description: "can read book",
	}
	CanUpdateBook model.Ability = model.Ability{
		Name:        "canUpdateBook",
		Description: "can update book",
	}
	CanDeleteBook model.Ability = model.Ability{
		Name:        "canDeleteBook",
		Description: "can delete book",
	}
	CanBorrowBook model.Ability = model.Ability{
		Name:        "canBorrowBook",
		Description: "can borrow book",
	}
	CanReturnBook model.Ability = model.Ability{
		Name:        "canReturnBook",
		Description: "can return book",
	}
	CanRenewBook model.Ability = model.Ability{
		Name:        "canRenewBook",
		Description: "can renew book",
	}
	CanManageBookRecords model.Ability = model.Ability{
		Name:        "canManageBookRecords",
		Description: "can manage book records",
	}
)
