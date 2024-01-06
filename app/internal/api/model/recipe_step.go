package model

import "time"

type RecipeStep struct {
	Base
	Timestamp
	Creator

	Title       string  `json:"title"`
	Description *string `json:"description"`
	Photo       *string `json:"photo"`

	Ordering int `json:"ordering"`
	RecipeId int `json:"recipe_id"`
}

func (r *RecipeStep) TableName() string {
	return "recipe_step"
}

type CreateRecipeStep struct {
	Title       string  `json:"title" binding:"required"`
	Description *string `json:"description"`

	Ordering int `json:"ordering" binding:"required"`
	RecipeId int `json:"recipeId" binding:"required"`

	CreatedById int       `json:"-"`
	Id          int       `json:"-"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

func (m CreateRecipeStep) Validate() error {
	return nil
}

type UpdateRecipeStep struct {
	Title       *string   `json:"title"`
	Description *string   `json:"description"`
	Ordering    *int      `json:"ordering"`
	UpdatedById *int      `json:"-"`
	UpdatedAt   time.Time `json:"-"`
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
	RecipeID int     `json:"recipe_id" schema:"recipe_id" binding:"required"`
	Title    *string `json:"title" schema:"title"`
}
