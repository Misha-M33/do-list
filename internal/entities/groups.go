package entities

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type Group struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Owner_id uuid.UUID `json:"owner_id"`
}

type Groups []Group

func (g *Group) ValidateCreateGroup() error {
	return validation.ValidateStruct(g,
		validation.Field(&g.Name, validation.Required,
			validation.Match(regexp.MustCompile("[А-ЯЁа-яёA-Za-z0-9]{1,21}")).Error("must be a string with А-ЯЁа-яё, A-Za-z, 0-9")),
	)
}

func (g *Group) ValidateDeleteGroup() error {
	return validation.ValidateStruct(g,
		validation.Field(&g.Name, validation.Required,
			validation.Match(regexp.MustCompile("[А-ЯЁа-яёA-Za-z0-9]{1,21}")).Error("must be a string with А-ЯЁа-яё, A-Za-z, 0-9")),
	)
}

func (g *Group) ValidateUpdateGroup() error {
	return validation.ValidateStruct(g,
		validation.Field(&g.Id, validation.Required),
		validation.Field(&g.Name, validation.Required,
			validation.Match(regexp.MustCompile("[А-ЯЁа-яёA-Za-z0-9]{1,21}")).Error("must be a string with А-ЯЁа-яё, A-Za-z, 0-9")),
		validation.Field(&g.Owner_id, validation.Required, is.UUID),
	)
}
