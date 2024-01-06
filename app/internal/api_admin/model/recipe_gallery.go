package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type RecipeGallery struct {
	Base
	Timestamp
	Creator
	Ordering  int    `json:"ordering"`
	Published bool   `json:"published"`
	Photo     string `json:"photo"`
	RecipeId  int    `json:"recipe_id"`
}

type RecipeGalleryList struct {
	Total  int              `json:"total"`
	Limit  int              `json:"limit"`
	Offset int              `json:"offset"`
	Data   []*RecipeGallery `json:"data"`
}

type CreateRecipeGallery struct {
	RecipeId  int    `json:"recipe_id" binding:"required" schema:"recipe_id"`
	Published *bool  `json:"published" schema:"published"`
	Ordering  int    `json:"ordering" schema:"ordering"`
	Photo     string `json:"-"`

	CreatedById int `json:"-" schema:"-"`
}

func (m *CreateRecipeGallery) Validate() error {
	if m.Published == nil {
		published := true
		m.Published = &published
	}
	// return nil
	return validation.ValidateStruct(m,
		validation.Field(&m.RecipeId, validation.Required),
	)
}

type UpdateRecipeGallery struct {
	Published *bool `json:"published"`
	Ordering  *int  `json:"ordering"`

	UpdatedById *int `json:"-"`
}

func (m UpdateRecipeGallery) Validate() error {
	return nil
}

type RecipeGalleryFilter struct {
	RecipeId int `json:"recipe_id" binding:"required"`
}
