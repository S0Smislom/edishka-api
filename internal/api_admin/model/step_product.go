package model

import "time"

type StepProduct struct {
	Id           int      `json:"id"`
	RecipeStepId int      `json:"recipe_step_id" binding:"required"`
	ProductId    int      `json:"-"`
	Product      *Product `json:"product" binding:"required"`
	Amount       float64  `json:"amount" binding:"required"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateStepProduct struct {
	RecipeStepId int     `json:"recipe_step_id" binding:"required"`
	ProductId    int     `json:"product_id" binding:"required"`
	Amount       float64 `json:"amount" binding:"required"`
}

func (m *CreateStepProduct) Validate() error {
	return nil
}

type UpdateStepProduct struct {
	Amount *float64 `json:"amount"`
}

func (m *UpdateStepProduct) Validate() error {
	return nil
}

type StepProductFilter struct {
	RecipeStepId *int `schema:"recipe_step_id"`
	ProductId    *int `schema:"product_id"`
}

type StepProductList struct {
	Total  int            `json:"total"`
	Limit  int            `json:"limit"`
	Offset int            `json:"offset"`
	Data   []*StepProduct `json:"data"`
}
