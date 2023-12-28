package userview

import (
	"lms-backend/internal/model"
	"lms-backend/internal/view/bookmarkview"
	"lms-backend/internal/view/fineview"
	"lms-backend/internal/view/loanview"
	"lms-backend/internal/view/personview"
	"lms-backend/internal/view/reservationview"
	"lms-backend/internal/view/sharedview"
	"lms-backend/util/sliceutil"
)

type LoginView struct {
	User         sharedview.UserView            `json:"user"`
	PersonView   *personview.View               `json:"person_attributes"`
	Abilities    []string                       `json:"abilities"`
	Bookmarks    []bookmarkview.DetailedView    `json:"bookmarks"`
	Loans        []loanview.WithBookView        `json:"loans"`
	Reservations []reservationview.WithBookView `json:"reservations"`
	Fines        []fineview.WithBookView        `json:"fines"`
}

func ToLoginView(
	user *model.User,
	abilities []model.Ability,
	bookmarks []model.Bookmark,
	loans []model.Loan,
	reservations []model.Reservation,
	fines []model.Fine,
) *LoginView {
	b := []bookmarkview.DetailedView{}
	for _, bookmark := range bookmarks {
		//nolint:gosec
		b = append(b, *bookmarkview.ToDetailedView(&bookmark))
	}

	l := []loanview.WithBookView{}
	for _, loan := range loans {
		//nolint:gosec
		l = append(l, *loanview.ToBookView(&loan))
	}

	r := []reservationview.WithBookView{}
	for _, reservation := range reservations {
		//nolint:gosec
		r = append(r, *reservationview.ToBookView(&reservation))
	}

	f := []fineview.WithBookView{}
	for _, fine := range fines {
		//nolint:gosec
		f = append(f, *fineview.ToBookView(&fine))
	}
	return &LoginView{
		User:         *sharedview.ToUserView(user),
		PersonView:   personview.ToView(user.Person),
		Abilities:    sliceutil.Map(abilities, func(a model.Ability) string { return a.Name }),
		Bookmarks:    b,
		Loans:        l,
		Reservations: r,
		Fines:        f,
	}
}
