package auth

import (
	"fmt"
	"lms-backend/internal/api"
	"lms-backend/internal/dataaccess/user"
	"lms-backend/internal/database"
	"lms-backend/internal/params/userparams"
	"lms-backend/internal/session"
	"lms-backend/internal/view/userview"

	"github.com/gofiber/fiber/v2"
)

// @Summary sign in a user
// @Description signs in a user and creates a session
// @Tags auth
// @Accept application/json
// @Param user body userparams.BaseUserParams true "User creation request"
// @Produce application/json
// @Success 200 {object} api.SwgResponse[userview.View]
// @Failure 400 {object} api.SwgErrResponse
// @Router /api/v1/auth/login [post]
func HandleSignIn(c *fiber.Ctx) error {
	var params userparams.BaseUserParams
	err := c.BodyParser(&params)
	if err != nil {
		return err
	}

	userModel := params.ToModel()
	db := database.GetDB()
	userModel, err = user.Login(db, userModel)
	if err != nil {
		return err
	}

	abilites, err := user.GetAbilities(db, int64(userModel.ID))
	if err != nil {
		return err
	}

	sess, err := session.Store.Get(c)
	if err != nil {
		return err
	}

	sess.Set(session.CookieKey, userModel.ID)
	err = sess.Save()
	if err != nil {
		return err
	}

	view := userview.ToView(userModel, abilites)

	return c.Status(fiber.StatusOK).JSON(api.Response{
		Data: view,
		Messages: api.Messages(
			api.SilentMessage(fmt.Sprintf(
				"%s is logged in successfully", userModel.Username,
			))),
	})
}
