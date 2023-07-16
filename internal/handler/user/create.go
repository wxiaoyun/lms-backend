package userhandler

import (
	"auth-practice/internal/api"
	"auth-practice/internal/database"
	"auth-practice/internal/params/userparams"
	"auth-practice/internal/view/userview"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func HandleCreateUser(c *fiber.Ctx) error {
	var params userparams.CreateUserParams
	err := c.BodyParser(&params)
	if err != nil {
		return err
	}

	user := params.ToModel()
	db := database.GetDB()
	err = user.Create(db)
	if err != nil {
		return err
	}

	view := userview.ToView(user)

	return c.Status(fiber.StatusCreated).JSON(api.ApiResponse{
		Data: view,
		Messages: []string{fmt.Sprintf(
			"User %s created successfully", user.Email,
		)},
	})
}
