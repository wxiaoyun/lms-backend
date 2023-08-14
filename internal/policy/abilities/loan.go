package abilities

import (
	"lms-backend/internal/model"
)

var (
	CanLoanBook model.Ability = model.Ability{
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
)
