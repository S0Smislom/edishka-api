package model

import "time"

type RecipeProduct struct {
	Base
	Timestamp
	Creator

	RecipeId  int      `json:"recipe_id" binding:"required"`
	ProductId int      `json:"-"`
	Product   *Product `json:"product" binding:"required"`
	Amount    float64  `json:"amount" binding:"required"`
}

func (m *RecipeProduct) TableName() string {
	return "recipe_product"
}

type CreateRecipeProduct struct {
	RecipeId  int     `json:"recipe_id" binding:"required"`
	ProductId int     `json:"product_id" binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`

	CreatedById int       `json:"-"`
	Id          int       `json:"-"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

func (m *CreateRecipeProduct) Validate() error {
	return nil
}

type UpdateRecipeProduct struct {
	Amount *float64 `json:"amount"`

	UpdatedById *int      `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

func (m *UpdateRecipeProduct) Validate() error {
	return nil
}

type RecipeProductFilter struct {
	RecipeId  *int `json:"recipe_id" schema:"recipe_id"`
	ProductId *int `json:"product_id" schema:"product_id"`
}

type RecipeProductList struct {
	Total  int              `json:"total"`
	Limit  int              `json:"limit"`
	Offset int              `json:"offset"`
	Data   []*RecipeProduct `json:"data"`
}
