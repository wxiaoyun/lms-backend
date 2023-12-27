package bookmarkparams

import (
	"lms-backend/internal/model"
	"lms-backend/pkg/error/externalerrors"
)

type BaseParams struct {
	UserID int64 `json:"user_id"`
	BookID int64 `json:"book_id"`
}

func (b *BaseParams) Validate() error {
	if b.UserID <= 0 {
		return externalerrors.BadRequest("user id is required")
	}

	if b.BookID <= 0 {
		return externalerrors.BadRequest("book id is required")
	}

	return nil
}

func (b *BaseParams) ToModel() *model.Bookmark {
	return &model.Bookmark{
		UserID: uint(b.UserID),
		BookID: uint(b.BookID),
	}
}
