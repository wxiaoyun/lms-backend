package bookmarkhandler

import (
	"fmt"
	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	"lms-backend/internal/dataaccess/book"
	"lms-backend/internal/dataaccess/bookmark"
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/database"
	"lms-backend/internal/model"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/bookmarkpolicy"
	"lms-backend/internal/session"
	"lms-backend/internal/view/bookmarkview"
	"lms-backend/pkg/error/externalerrors"
	"strconv"

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

	param := c.Params("book_id")
	bookID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid loan id.", param))
	}

	userID, err := session.GetLoginSession(c)
	if err != nil {
		return err
	}

	bm := model.Bookmark{BookID: uint(bookID), UserID: uint(userID)}

	db := database.GetDB()

	username, err := user.GetUserName(db, userID)
	if err != nil {
		return err
	}

	bookTitle, err := book.GetBookTitle(db, bookID)
	if err != nil {
		return err
	}

	tx, rollBackOrCommit := audit.Begin(
		c, fmt.Sprintf("%s bookmarking \"%s\"", username, bookTitle),
	)
	defer func() { rollBackOrCommit(err) }()

	b, err := bookmark.Create(tx, &bm)
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
