package auth

import (
	"auth-practice/internal/api"
	"auth-practice/internal/dataaccess/user"
	"auth-practice/internal/database"
	"auth-practice/internal/params/userparams"
	"auth-practice/internal/session"
	"auth-practice/internal/view/userview"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// @Summary sign in a user
// @Description signs in a user and creates a session
// @Tags user
// @Accept */*
// @Produce plain
// @Success 200 "OK"
// @Router /api/v1/login [post]
func HandleSignIn(c *fiber.Ctx) error {
	var params userparams.BaseUserParams
	err := c.BodyParser(&params)
	if err != nil {
		return err
	}

	userModel := params.ToModel()
	db := database.GetDB()
	err = user.VerifyLogin(db, userModel)
	if err != nil {
		return err
	}

	sess, err := session.Store.Get(c)
	if err != nil {
		return err
	}

	sess.Set(session.CookieKey, session.LoginSession{
		UserID:         userModel.ID,
		Email:          userModel.Email,
		IsMasquerading: false,
	})
	err = sess.Save()
	if err != nil {
		return err
	}
	c.UserContext()

	view := userview.ToView(userModel)

	return c.Status(fiber.StatusOK).JSON(api.Response{
		Data: view,
		Messages: []string{fmt.Sprintf(
			"User %s is logged in successfully", userModel.Email,
		)},
	})
}
