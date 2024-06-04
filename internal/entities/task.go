package entities

import (
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type Task struct {
	Id            int       `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Responsible   uuid.UUID `json:"responsible"`
	Priority      int       `json:"priority"`
	Is_done       bool      `json:"is_done"`
	Creator       uuid.UUID `json:"creator"`
	Group_id      int       `json:"group_id"`
	Deadline_date string    `json:"deadline_date"`
	Created_at    time.Time `json:"created_at"`
}

type Tasks []Task

func (t *Task) ValidateUpdateTaskById() error {
	return validation.ValidateStruct(t,
		validation.Field(&t.Id, validation.Required),
		validation.Field(&t.Title, validation.Required, validation.Match(regexp.MustCompile("[А-ЯЁа-яёA-Za-z0-9]{1,21}")).Error("must be a string with А-ЯЁа-яё, A-Za-z, 0-9")),
		validation.Field(&t.Description, validation.Required, validation.Match(regexp.MustCompile("[А-ЯЁа-яёA-Za-z0-9]{1,21}")).Error("must be a string with А-ЯЁа-яё, A-Za-z, 0-9")),
		validation.Field(&t.Responsible, is.UUID),
		validation.Field(&t.Priority, validation.Required),
		validation.Field(&t.Is_done, validation.Required),
		validation.Field(&t.Creator, validation.Required, is.UUID),
		validation.Field(&t.Group_id, validation.Required),
		validation.Field(&t.Deadline_date, validation.Required),
	)
}

func (t *Task) ValidateDeleteTaskById() error {
	return validation.ValidateStruct(t,
		validation.Field(&t.Id, validation.Required),
	)
}

func (t *Task) ValidateCreateTask() error {
	return validation.ValidateStruct(t,
		validation.Field(&t.Title, validation.Required, validation.Match(regexp.MustCompile("[А-ЯЁа-яёA-Za-z0-9]{1,21}")).Error("must be a string with А-ЯЁа-яё, A-Za-z, 0-9")),
		validation.Field(&t.Description, validation.Required, validation.Match(regexp.MustCompile("[А-ЯЁа-яёA-Za-z0-9]{1,21}")).Error("must be a string with А-ЯЁа-яё, A-Za-z, 0-9")),
		validation.Field(&t.Responsible, is.UUID),
		validation.Field(&t.Priority, validation.Required),
		validation.Field(&t.Creator, validation.Required, is.UUID),
		validation.Field(&t.Group_id, validation.Required),
		validation.Field(&t.Deadline_date, validation.Required),
	)
}

func (t *Task) ValidateGetTaskByGroupId() error {
	return validation.ValidateStruct(t,
		validation.Field(&t.Group_id, validation.Required),
	)
}
