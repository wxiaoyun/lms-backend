package userhandler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"auth-practice/internal/api"
	"auth-practice/internal/database"
	"auth-practice/internal/params/userparams"
	"auth-practice/internal/view/userview"
)

// @Summary Create a user
// @Description create an instance of user in the database
// @Tags user
// @Accept */*
// @Produce plain
// @Success 200 "OK"
// @Router /api/v1/signup [post]
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

	return c.Status(fiber.StatusCreated).JSON(api.Response{
		Data: view,
		Messages: []string{fmt.Sprintf(
			"User %s created successfully", user.Email,
		)},
	})
}
