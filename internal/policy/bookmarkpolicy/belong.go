package bookmarkpolicy

import (
	"fmt"
	"lms-backend/internal/database"
	"lms-backend/internal/model"
	"lms-backend/internal/policy"
	"lms-backend/internal/session"

	"github.com/gofiber/fiber/v2"
)

type BookMarkBelongsToUser struct {
	BookmarkID int64
}

func AllowIfBookmarkBelongsToUser(bookmarkID int64) *BookMarkBelongsToUser {
	return &BookMarkBelongsToUser{
		BookmarkID: bookmarkID,
	}
}

func (p *BookMarkBelongsToUser) Validate(c *fiber.Ctx) (policy.Decision, error) {
	userID, err := session.GetLoginSession(c)
	if err != nil {
		return policy.Deny, err
	}

	db := database.GetDB()

	var exists int64
	result := db.Model(&model.Bookmark{}).
		Where("id = ? AND user_id = ?", p.BookmarkID, userID).
		Count(&exists)
	if result.Error != nil {
		return policy.Deny, result.Error
	}

	if exists == 0 {
		return policy.Deny, nil
	}

	return policy.Allow, nil
}

func (p *BookMarkBelongsToUser) Reason() string {
	return fmt.Sprintf("Bookmark with ID %d does not belong to you.", p.BookmarkID)
}
