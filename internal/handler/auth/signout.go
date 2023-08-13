package auth

import (
	"lms-backend/internal/api"
	"lms-backend/internal/session"

	"github.com/gofiber/fiber/v2"
)

// @Summary SignOut
// @Description signs out the user and revoke their session
// @Tags user
// @Accept */*
// @Produce plain
// @Success 200 "OK"
// @Router /api/v1/auth/logout [get]
func HandleSignOut(c *fiber.Ctx) error {
	sess, err := session.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(api.Response{
			Messages: []api.Message{api.ErrorMessage("User is not logged in")},
		})
	}

	err = sess.Destroy()
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(api.Response{
		Messages: []api.Message{api.SilentMessage("User is logged out successfully")},
	})
}
