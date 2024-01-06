package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type User struct {
	Base
	Timestamp
	Phone     string  `json:"phone"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Birthday  *string `json:"birthday"`

	IsSuperuser bool `json:"is_superuser"`
	IsStaff     bool `json:"is_staff"`

	Password *string `json:"-"`
}

type CreateUser struct {
	Phone       string `json:"phone" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Password2   string `json:"password2" binding:"required"`
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
