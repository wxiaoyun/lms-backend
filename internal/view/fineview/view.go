package fineview

import (
	"lms-backend/internal/model"
	"lms-backend/internal/view/sharedview"
)

type BaseView struct {
	sharedview.FineView
}

func ToBaseView(fine *model.Fine) *BaseView {
	return &BaseView{
		FineView: *sharedview.ToFineView(fine),
	}
}

func ToViews(fines []model.Fine) []BaseView {
	views := make([]BaseView, 0, len(fines))
	for _, fine := range fines {
		//nolint
		views = append(views, *ToBaseView(&fine))
	}
	return views
}
