package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Diet struct {
	Base
	Timestamp

	Title       string  `json:"title"`
	Slug        string  `json:"slug"`
	Description *string `json:"description"`
	Published   bool    `json:"published"`

	Items []*DietItem `json:"items"`

	UserId int `json:"-"`
}

func (m *Diet) TableName() string {
	return "diet"
}

type DietList struct {
	Total  int     `json:"total"`
	Limit  int     `json:"limit"`
	Offset int     `json:"offset"`
	Data   []*Diet `json:"data"`
}

type CreateDiet struct {
	Title       string  `json:"title" binding:"required"`
	Slug        string  `json:"slug" binding:"required"`
	Description *string `json:"description" binding:"required"`
	Published   bool    `json:"published"`
	ItemIds     *[]int  `json:"diet_item_ids" gorm:"-"`

	UserId int `json:"-"`
	Id     int `json:"-"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (m CreateDiet) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Title, validation.Required),
		validation.Field(&m.Slug, validation.Required),
	)
}

type UpdateDiet struct {
	Title       *string `json:"title"`
	Slug        *string `json:"slug"`
	Description *string `json:"description"`
	Published   *bool   `json:"published"`

	UpdatedAt time.Time `json:"-"`
}

func (m UpdateDiet) Validate() error {
	return nil
}

type DietFilter struct {
	Title  *string `json:"title" schema:"title"`
	UserId int     `json:"-" schema:"-"`
}
