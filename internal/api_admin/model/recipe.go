package model

import "time"

type DifficultyLevel string

const (
	Easy   DifficultyLevel = "easy"
	Normal DifficultyLevel = "normal"
	Hard   DifficultyLevel = "hard"
)

type Recipe struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Title       string  `json:"title"`
	Slug        string  `json:"slug"`
	Description *string `json:"description"`
	Photo       *string `json:"photo"`

	CookingTime   int  `json:"cooking_time"`
	PreparingTime *int `json:"preparing_time"`
	// TODO Вынести в отдельную таблицу
	Kitchen         string          `json:"kitchen"`
	DifficultyLevel DifficultyLevel `json:"difficulty_level"`
	Published       bool            `json:"published"`
}

type CreateRecipe struct {
	Title       string  `json:"title"`
	Slug        string  `json:"slug"`
	Description *string `json:"description"`

	CookingTime   int  `json:"cooking_time"`
	PreparingTime *int `json:"preparing_time"`
	// TODO Вынести в отдельную таблицу
	Kitchen         string          `json:"kitchen"`
	DifficultyLevel DifficultyLevel `json:"difficulty_level"`
}

func (m CreateRecipe) Validate() error {
	return nil
}

type UpdateRecipe struct {
	Title       *string `json:"title"`
	Slug        *string `json:"slug"`
	Description *string `json:"description"`

	CookingTime   *int `json:"cooking_time"`
	PreparingTime *int `json:"preparing_time"`
	// TODO Вынести в отдельную таблицу
	Kitchen         *string          `json:"kitchen"`
	DifficultyLevel *DifficultyLevel `json:"difficulty_level"`
	Published       *bool            `json:"published"`
}

func (m UpdateRecipe) Validate() error {
	return nil
}

type RecipeList struct {
	Total  int       `json:"total"`
	Limit  int       `json:"limit"`
	Offset int       `json:"offset"`
	Data   []*Recipe `json:"data"`
}

type RecipeFilter struct {
	Title           *string          `json:"title"`
	Slug            *string          `json:"slug"`
	CookingTimeGTE  *int             `json:"cookingTimeGTE"`
	CookingTimeLTE  *int             `json:"cookingTimeLTE"`
	Kitchen         *string          `json:"kitchen"`
	DifficultyLevel *DifficultyLevel `json:"difficulty_level"`
}
