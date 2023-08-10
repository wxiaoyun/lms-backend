package userparams

import (
	"fmt"
	"lms-backend/internal/model"
	"lms-backend/internal/params/peopleparams"
)

type CreateUserParams struct {
	BaseUserParams
	PersonParams peopleparams.BaseParams `json:"person_attributes"`
}

func (c *CreateUserParams) ToModel() *model.User {
	usr := c.BaseUserParams.ToModel()
	usr.Person = c.PersonParams.ToModel()
	fmt.Println(usr)
	return usr
}

func (c *CreateUserParams) Validate() error {
	if err := c.BaseUserParams.Validate(); err != nil {
		return err
	}

	return c.PersonParams.Validate()
}
