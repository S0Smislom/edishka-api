package model

import validation "github.com/go-ozzo/ozzo-validation/v4"

type User struct {
	ID        int     `json:"id"`
	Phone     string  `json:"phone"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Birthday  *string `json:"birthday"`
	Photo     *string `json:"photo"`
	Code      *string `json:"-"`
}

func (m *User) TableName() string {
	return "\"user\""
}

type UpdateUser struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Birthday  *string `json:"birthday"`
}

func (m *UpdateUser) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.FirstName, validation.When(m.FirstName != nil, validation.Required)),
		validation.Field(&m.LastName, validation.When(m.LastName != nil, validation.Required)),
		validation.Field(&m.Birthday, validation.When(m.Birthday != nil, validation.Date("2006-01-02"))),
	)
}
