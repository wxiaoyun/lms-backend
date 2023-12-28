package userview

import (
	"lms-backend/internal/model"
	"lms-backend/internal/view/bookmarkview"
	"lms-backend/internal/view/fineview"
	"lms-backend/internal/view/loanview"
	"lms-backend/internal/view/reservationview"
)

type CurrentUserView struct {
	IsLoggedIn   bool                           `json:"is_logged_in"`
	User         *View                          `json:"user,omitempty"`
	Bookmarks    []bookmarkview.DetailedView    `json:"bookmarks,omitempty"`
	Loans        []loanview.DetailedView        `json:"loans,omitempty"`
	Reservations []reservationview.DetailedView `json:"reservations,omitempty"`
	Fines        []fineview.DetailedView        `json:"fines,omitempty"`
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
	b := []bookmarkview.DetailedView{}
	for _, bookmark := range bookmarks {
		//nolint:gosec
		b = append(b, *bookmarkview.ToDetailedView(&bookmark))
	}

	l := []loanview.DetailedView{}
	for _, loan := range loans {
		//nolint:gosec
		l = append(l, *loanview.ToDetailedView(&loan))
	}

	r := []reservationview.DetailedView{}
	for _, reservation := range reservations {
		//nolint:gosec
		r = append(r, *reservationview.ToDetailedView(&reservation))
	}

	f := []fineview.DetailedView{}
	for _, fine := range fines {
		//nolint:gosec
		f = append(f, *fineview.ToDetailedView(&fine))
	}

	return &CurrentUserView{
		IsLoggedIn:   true,
		User:         ToView(user, abilities...),
		Bookmarks:    b,
		Loans:        l,
		Reservations: r,
		Fines:        f,
	}
}
