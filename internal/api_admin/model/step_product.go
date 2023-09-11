package model

type StepProduct struct {
	Base
	Timestamp
	RecipeStepId int      `json:"recipe_step_id" binding:"required"`
	ProductId    int      `json:"-"`
	Product      *Product `json:"product" binding:"required"`
	Amount       float64  `json:"amount" binding:"required"`
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
	RecipeStepId *int `json:"recipe_step_id" schema:"recipe_step_id"`
	ProductId    *int `json:"product_id" schema:"product_id"`
}

type StepProductList struct {
	Total  int            `json:"total"`
	Limit  int            `json:"limit"`
	Offset int            `json:"offset"`
	Data   []*StepProduct `json:"data"`
}
