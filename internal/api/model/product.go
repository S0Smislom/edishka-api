package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Product struct {
	Base

	Title       string  `json:"title"`
	Slug        string  `json:"slug"`
	Description *string `json:"description"`
	Photo       *string `json:"photo"`

	Calories        int     `json:"calories"`
	Squirrels       float64 `json:"squirrels"`
	Fats            float64 `json:"fats"`
	Carbohydrates   float64 `json:"carbohydrates"`
	SuggestedByUser bool    `json:"-"`
	CreatedById     int     `json:"-"`
}

type ProductList struct {
	Total  int        `json:"total"`
	Limit  int        `json:"limit"`
	Offset int        `json:"offset"`
	Data   []*Product `json:"data"`
}

type CreateProduct struct {
	Title       string  `json:"title" binding:"required"`
	Slug        string  `json:"slug" binding:"required"`
	Description *string `json:"description"`
	// Photo       *string `json:"photo"`

	Calories      int     `json:"calories" binding:"required"`
	Squirrels     float64 `json:"squirrels" binding:"required"`
	Fats          float64 `json:"fats" binding:"required"`
	Carbohydrates float64 `json:"carbohydrates" binding:"required"`

	CreatedById int `json:"-"`
}

func (m CreateProduct) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Title, validation.Required),
		validation.Field(&m.Slug, validation.Required),
		validation.Field(&m.Calories, validation.Min(0)),
		validation.Field(&m.Squirrels, validation.Min(0.0)),
		validation.Field(&m.Fats, validation.Min(0.0)),
		validation.Field(&m.Carbohydrates, validation.Min(0.0)),
	)
}

type UpdateProduct struct {
	Title         *string  `json:"title"`
	Slug          *string  `json:"slug"`
	Description   *string  `json:"description"`
	Calories      *int     `json:"calories"`
	Squirrels     *float64 `json:"squirrels"`
	Fats          *float64 `json:"fats"`
	Carbohydrates *float64 `json:"carbohydrates"`

	UpdatedById *int `json:"-"`
}

func (m UpdateProduct) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Title, validation.When(m.Title != nil, validation.Required)),
		validation.Field(&m.Slug, validation.When(m.Slug != nil, validation.Required)),
		validation.Field(&m.Calories, validation.Min(0)),
		validation.Field(&m.Squirrels, validation.Min(0.0)),
		validation.Field(&m.Fats, validation.Min(0.0)),
		validation.Field(&m.Carbohydrates, validation.Min(0.0)),
	)
}

type ProductFilter struct {
	Title            *string  `json:"title" schema:"title"`
	Slug             *string  `json:"slug" schema:"slug"`
	CaloriesGTE      *int     `json:"calories__gte" schema:"calories__gte"`
	CaloriesLTE      *int     `json:"calories__lte" schema:"calories__lte"`
	SquirrelsGTE     *float64 `json:"squirrels__gte" schema:"squirrels__gte"`
	SquirrelsLTE     *float64 `json:"squirrels__lte" schema:"squirrels__lte"`
	FatsGTE          *float64 `json:"fats__gte" schema:"fats__gte"`
	FatsLTE          *float64 `json:"fats__lte" schema:"fats__lte"`
	CarbohydratesGTE *float64 `json:"carbohydrates__gte" schema:"carbohydrates__gte"`
	CarbohydratesLTE *float64 `json:"carbohydrates__lte" schema:"carbohydrates__lte"`
}
