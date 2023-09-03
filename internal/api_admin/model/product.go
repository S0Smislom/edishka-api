package model

import "time"

type Product struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Title       string  `json:"title"`
	Slug        string  `json:"slug"`
	Description *string `json:"description"`
	Photo       *string `json:"photo"`

	Calories      int     `json:"calories"`
	Squirrels     float64 `json:"squirrels"`
	Fats          float64 `json:"fats"`
	Carbohydrates float64 `json:"carbohydrates"`
}

type ProductList struct {
	Total  int        `json:"total"`
	Limit  int        `json:"limit"`
	Offset int        `json:"offset"`
	Data   []*Product `json:"data"`
}

type CreateProduct struct {
	Title       string  `json:"title"`
	Slug        string  `json:"slug"`
	Description *string `json:"description"`
	// Photo       *string `json:"photo"`

	Calories      int     `json:"calories"`
	Squirrels     float64 `json:"squirrels"`
	Fats          float64 `json:"fats"`
	Carbohydrates float64 `json:"carbohydrates"`
}

func (m CreateProduct) Validate() error {
	return nil
}

type UpdateProduct struct {
	Title       *string `json:"title"`
	Slug        *string `json:"slug"`
	Description *string `json:"description"`
	// Photo       *string `json:"photo"`

	Calories      *int     `json:"calories"`
	Squirrels     *float64 `json:"squirrels"`
	Fats          *float64 `json:"fats"`
	Carbohydrates *float64 `json:"carbohydrates"`
}

func (m UpdateProduct) Validate() error {
	return nil
}

type ProductFilter struct {
	Title            *string  `json:"title"`
	Slug             *string  `json:"slug"`
	CaloriesGTE      *int     `json:"caloriesGTE"`
	CaloriesLTE      *int     `json:"caloriesLTE"`
	SquirrelsGTE     *float64 `json:"squirrelsGTE"`
	SquirrelsLTE     *float64 `json:"squirrelsLTE"`
	FatsGTE          *float64 `json:"fatsGTE"`
	FatsLTE          *float64 `json:"fatsLTE"`
	CarbohydratesGTE *float64 `json:"carbohydratesGTE"`
	CarbohydratesLTE *float64 `json:"carbohydratesLTE"`
}
