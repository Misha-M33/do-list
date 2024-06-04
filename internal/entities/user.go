package entities

import (
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"github.com/google/uuid"
)

type User struct {
	First_name string    `json:"first_name"`
	Last_name  string    `json:"last_name"`
	Full_name  string    `json:"full_name"`
	Nickname   string    `json:"nickname"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Is_deleted bool      `json:"is_deleted"`
	Is_block   bool      `json:"is_block"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Id         uuid.UUID `json:"id"`
}

type Users []User

func (u *User) ValidateCreateUser() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Id, validation.Required, is.UUID),
		validation.Field(&u.First_name,
			validation.Required, validation.Match(regexp.MustCompile("^[А-ЯЁA-Z]{1}[а-яёa-z]{1,21}$")).Error("must be a string first A-Z next a-z")),
		validation.Field(&u.Last_name,
			validation.Match(regexp.MustCompile("^[А-ЯЁA-Z]{1}[а-яёa-z]{1,21}$")).Error("must be a string first A-Z next a-z")),
		validation.Field(&u.Full_name,
			validation.Match(regexp.MustCompile("^[А-ЯЁA-Z]{1}[а-яёa-z]{1,21}$")).Error("must be a string first A-Z next a-z")),
		validation.Field(&u.Nickname,
			validation.Match(regexp.MustCompile("[А-ЯЁа-яёA-Za-z0-9]{1,21}")).Error("must be a string with А-ЯЁа-яё, A-Za-z, 0-9")),
		validation.Field(&u.Email, validation.Required,
			is.Email.Error("must be a string type email"), is.EmailFormat.Error("to fix format")),
		validation.Field(&u.Password, validation.Required,
			validation.Length(4, 12)),
	)
}

func (u *User) ValidateRegistration() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required,
			is.Email.Error("must be a string type email"), is.EmailFormat.Error("to fix format")),
		validation.Field(&u.Password, validation.Required,
			validation.Length(4, 12)),
	)
}

func (u *User) ValidateUpdate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Id, validation.Required, is.UUID),
		validation.Field(&u.First_name, validation.Required,
			validation.Match(regexp.MustCompile("^[А-ЯЁA-Z]{1}[а-яёa-z]{1,21}$")).Error("must be a string first A-Z next a-z")),
		validation.Field(&u.Last_name,
			validation.Match(regexp.MustCompile("^[А-ЯЁA-Z]{1}[а-яёa-z]{1,21}$")).Error("must be a string first A-Z next a-z")),
		validation.Field(&u.Full_name,
			validation.Match(regexp.MustCompile("^[А-ЯЁA-Z]{1}[а-яёa-z]{1,21}$")).Error("must be a string first A-Z next a-z")),
		validation.Field(&u.Nickname,
			validation.Match(regexp.MustCompile("[А-ЯЁа-яёA-Za-z0-9]{1,21}")).Error("must be a string with А-ЯЁа-яё, A-Za-z, 0-9")),
	)
}

func (u *User) ValidatePassword() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Id, validation.Required, is.UUID),
		validation.Field(&u.Password,
			validation.Required, validation.Length(4, 12)),
	)
}

func (u *User) ValidateEmail() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required,
			is.Email.Error("must be a string type email"), is.EmailFormat.Error("to fix format")),
	)
}
