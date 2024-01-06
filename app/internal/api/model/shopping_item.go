package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ShoppingItem struct {
	Base

	Title  string  `json:"title"`
	Amount float64 `json:"amount"`
	UserId int     `json:"-"`
}

func (r *ShoppingItem) TableName() string {
	return "shopping_item"
}

type ShoppingList struct {
	Total  int             `json:"total"`
	Limit  int             `json:"limit"`
	Offset int             `json:"offset"`
	Data   []*ShoppingItem `json:"data"`
}

type CreateShoppingItem struct {
	Id     int     `json:"-"`
	Title  string  `json:"title"`
	Amount float64 `json:"amount"`
	UserId int     `json:"-"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (m CreateShoppingItem) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Title, validation.Required),
		validation.Field(&m.Amount, validation.Min(0.0)),
	)
}

type UpdateShoppingItem struct {
	Title  *string  `json:"title"`
	Amount *float64 `json:"amount"`

	UpdatedAt time.Time `json:"-"`
}

func (m UpdateShoppingItem) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Title, validation.When(m.Title != nil, validation.Required)),
		validation.Field(&m.Amount, validation.Min(0.0)),
	)
}

type ShoppingItemFilter struct {
	Title  *string  `json:"title" schema:"title"`
	Amount *float64 `json:"amount" schema:"amount"`
	UserId int      `json:"-" schema:"-"`
}
