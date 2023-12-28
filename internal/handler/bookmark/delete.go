package bookmarkhandler

import (
	"fmt"
	"lms-backend/internal/api"
	audit "lms-backend/internal/auditlog"
	"lms-backend/internal/dataaccess/bookmark"
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/database"
	"lms-backend/internal/policy"
	"lms-backend/internal/policy/bookmarkpolicy"
	"lms-backend/internal/session"
	"lms-backend/internal/view/bookmarkview"
	"lms-backend/pkg/error/externalerrors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const (
	deleteBookmarkAction = "delete book mark"
)

func HandleDelete(c *fiber.Ctx) error {
	err := policy.Authorize(c, deleteBookmarkAction, bookmarkpolicy.CreatePolicy())
	if err != nil {
		return err
	}

	userID, err := session.GetLoginSession(c)
	if err != nil {
		return err
	}

	id := c.Params("bookmark_id")
	bmID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return externalerrors.BadRequest(fmt.Sprintf("%s is not a valid bookmark id.", id))
	}

	db := database.GetDB()

	username, err := user.GetUserName(db, userID)
	if err != nil {
		return err
	}

	tx, rollBackOrCommit := audit.Begin(
		c, fmt.Sprintf("%s is deleting bookmark id - \"%d\"", username, bmID),
	)
	defer func() { rollBackOrCommit(err) }()

	b, err := bookmark.Delete(tx, bmID)
	if err != nil {
		return err
	}

	return c.JSON(api.Response{
		Data: bookmarkview.ToDetailedView(b),
		Messages: api.Messages(
			api.SilentMessage("bookmark deleted successfully"),
		),
	})
}
