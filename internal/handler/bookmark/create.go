package bookmarkhandler

import (
	"fmt"
	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	"lms-backend/internal/dataaccess/book"
	"lms-backend/internal/dataaccess/bookmark"
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/database"
	"lms-backend/internal/params/bookmarkparams"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/bookmarkpolicy"
	"lms-backend/internal/view/bookmarkview"

	"github.com/gofiber/fiber/v2"
)

const (
	createBookmarkAction = "create book mark"
)

func HandleCreate(c *fiber.Ctx) error {
	err := policy.Authorize(c, createBookmarkAction, bookmarkpolicy.CreatePolicy())
	if err != nil {
		return err
	}

	var params bookmarkparams.BaseParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	if err := params.Validate(); err != nil {
		return err
	}

	bm := params.ToModel()

	db := database.GetDB()

	username, err := user.GetUserName(db, params.UserID)
	if err != nil {
		return err
	}

	bookTitle, err := book.GetBookTitle(db, params.BookID)
	if err != nil {
		return err
	}

	tx, rollBackOrCommit := audit.Begin(
		c, fmt.Sprintf("%s bookmarking \"%s\"", username, bookTitle),
	)
	defer func() { rollBackOrCommit(err) }()

	b, err := bookmark.Create(tx, bm)
	if err != nil {
		return err
	}

	return c.JSON(api.Response{
		Data: bookmarkview.ToDetailedView(b),
		Messages: api.Messages(
			api.SilentMessage("bookmark created successfully"),
		),
	})
}
