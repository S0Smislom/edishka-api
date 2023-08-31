package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type User struct {
	ID        int     `json:"id"`
	Phone     string  `json:"phone"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Birthday  *string `json:"birthday"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	IsSuperuser bool `json:"is_superuser"`
	IsStaff     bool `json:"is_staff"`

	Password *string `json:"-"`
}

type CreateUser struct {
	Phone       string `json:"phone"`
	Password    string `json:"password"`
	Password2   string `json:"password2"`
	IsSuperuser *bool  `json:"is_superuser"`
	IsStaff     *bool  `json:"is_staff"`
}

func (l CreateUser) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Phone, validation.Required),
		validation.Field(&l.Password, validation.Required),
		validation.Field(&l.Password2, validation.Required),
	)
}
