package model

import validation "github.com/go-ozzo/ozzo-validation/v4"

type DietItem struct {
	Base

	ProductId   *int    `json:"product_id"`
	RecipeId    *int    `json:"recipe_id"`
	Amount      float64 `json:"amount"`
	DietId      *int    `json:"-"`
	CreatedById int     `json:"-"`
}

func (m *DietItem) TableName() string {
	return "diet_item"
}

type DietItemList struct {
	Total  int         `json:"total"`
	Limit  int         `json:"limit"`
	Offset int         `json:"offset"`
	Data   []*DietItem `json:"data"`
}

type CreateDietItem struct {
	Id          int     `json:"-"`
	ProductId   *int    `json:"product_id"`
	RecipeId    *int    `json:"recipe_id"`
	Amount      float64 `json:"amount" binding:"required"`
	DietId      *int    `json:"diet_id"`
	CreatedById int     `json:"-"`
}

func (m CreateDietItem) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.ProductId, validation.When(m.RecipeId == nil, validation.Required)),
		validation.Field(&m.RecipeId, validation.When(m.ProductId == nil, validation.Required)),
		validation.Field(&m.Amount, validation.Min(0.0)),
	)
}

type UpdateDietItem struct {
	Amount *float64 `json:"amount"`
	DietId *int     `json:"diet_id"`
}

func (m UpdateDietItem) Validate() error {
	return nil
}

type DietItemFilter struct {
	DietId      int    `json:"diet_id" schema:"diet_id" binding:"required"`
	IdList      *[]int `json:"-" schema:"-"`
	CreatedById *int   `json:"-" schema:"-"`
}

func (m DietItemFilter) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.DietId, validation.Required),
	)
}
