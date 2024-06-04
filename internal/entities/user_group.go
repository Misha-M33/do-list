package entities

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type UserGroup struct {
	Id       int       `json:"id"`
	User_id  uuid.UUID `json:"user_id"`
	Group_id int       `json:"group_id"`
}

type UsersGroups []UserGroup

func (g *UserGroup) ValidateCreateUserGroup() error {
	return validation.ValidateStruct(g,
		validation.Field(&g.User_id, validation.Required, is.UUID),
		validation.Field(&g.Group_id, validation.Required),
	)
}

func (g *UserGroup) ValidateDeleteUserGroup() error {
	return validation.ValidateStruct(g,
		validation.Field(&g.User_id, validation.Required, is.UUID),
	)
}
