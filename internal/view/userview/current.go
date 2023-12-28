package userview

import (
	"lms-backend/internal/model"
)

type CurrentUserView struct {
	IsLoggedIn bool `json:"is_logged_in"`
	LoginView
}

func ToGuestView() *CurrentUserView {
	return &CurrentUserView{
		IsLoggedIn: false,
	}
}

func ToCurrentUserView(
	user *model.User,
	abilities []model.Ability,
	bookmarks []model.Bookmark,
	loans []model.Loan,
	reservations []model.Reservation,
	fines []model.Fine,
) *CurrentUserView {
	return &CurrentUserView{
		IsLoggedIn: true,
		LoginView:  *ToLoginView(user, abilities, bookmarks, loans, reservations, fines),
	}
}
