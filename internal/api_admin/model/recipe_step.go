package model

import "time"

type RecipeStep struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Title       string  `json:"title"`
	Description *string `json:"description"`
	Photo       *string `json:"photo"`

	Ordering int `json:"ordering"`
	RecipeId int `json:"recipeId"`
}

type CreateRecipeStep struct {
	Title       string  `json:"title" binding:"required"`
	Description *string `json:"description"`
	Photo       *string `json:"photo"`

	Ordering int `json:"ordering" binding:"required"`
	RecipeId int `json:"recipeId" binding:"required"`
}

func (m CreateRecipeStep) Validate() error {
	return nil
}

type UpdateRecipeStep struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Ordering    *int    `json:"ordering"`
}

func (m UpdateRecipeStep) Validate() error {
	return nil
}

type RecipeStepList struct {
	Total  int           `json:"total" binding:"required"`
	Limit  int           `json:"limit" binding:"required"`
	Offset int           `json:"offset" binding:"required"`
	Data   []*RecipeStep `json:"data" binding:"required"`
}

type RecipeStepFilter struct {
	RecipeID int     `json:"recipeId" binding:"required"`
	Title    *string `json:"title"`
}
