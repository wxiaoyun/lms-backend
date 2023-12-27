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

func HandleSignIn(c *fiber.Ctx) error {
	// if handler, ok := c.Locals(csrf.ConfigDefault.HandlerContextKey).(*csrf.CSRFHandler); ok {
	// 	if err := handler.DeleteToken(c); err != nil {
	// 		return err
	// 	}
	// }

	var params userparams.SignInParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	if err := params.Validate(); err != nil {
		return err
	}

	userModel := params.ToModel()
	db := database.GetDB()
	userModel, err := user.Login(db, userModel)
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

	err = sess.Regenerate()
	if err != nil {
		return err
	}

	sess.Set(session.CookieKey, userModel.ID)
	err = sess.Save()
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(api.Response{
		Data: userview.ToView(userModel, abilites...),
		Messages: api.Messages(
			api.SilentMessage(fmt.Sprintf(
				"%s is logged in successfully", userModel.Username,
			))),
	})
}
