package model

type DifficultyLevel string

const (
	Easy   DifficultyLevel = "easy"
	Normal DifficultyLevel = "normal"
	Hard   DifficultyLevel = "hard"
)

type Recipe struct {
	Base
	Timestamp
	Creator

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

	// Calculated fields
	Calories      float64 `json:"calories"`
	Squirrels     float64 `json:"squirrels"`
	Fats          float64 `json:"fats"`
	Carbohydrates float64 `json:"carbohydrates"`

	Products []*RecipeProduct        `json:"products"`
	Gallery  []*RecipeGalleryMinimal `json:"gallery"`
}

type RecipeProduct struct {
	Product
	Amount int `json:"amount"`
}

type RecipeGalleryMinimal struct {
	Base
	Timestamp
	Ordering  int    `json:"ordering"`
	Published bool   `json:"published"`
	Photo     string `json:"photo"`
}

type CreateRecipe struct {
	Title       string  `json:"title" binding:"required"`
	Slug        string  `json:"slug" binding:"required"`
	Description *string `json:"description"`

	CookingTime   int  `json:"cooking_time" binding:"required"`
	PreparingTime *int `json:"preparing_time"`
	// TODO Вынести в отдельную таблицу
	Kitchen         string          `json:"kitchen" binding:"required"`
	DifficultyLevel DifficultyLevel `json:"difficulty_level" binding:"required"`

	CreatedById int `json:"-"`
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

	UpdatedById *int `json:"-"`
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
	Title           *string          `schema:"title" json:"title"`
	Slug            *string          `schema:"slug" json:"slug"`
	CookingTimeGTE  *int             `schema:"cooking_time__gte" json:"cooking_time__gte"`
	CookingTimeLTE  *int             `schema:"cooking_time__lte" json:"cooking_time__lte"`
	Kitchen         *string          `schema:"kitchen" json:"kitchen"`
	DifficultyLevel *DifficultyLevel `schema:"difficulty_level" json:"difficulty_level"`
	CreatedById     *int             `schema:"-" json:"-"`
	Published       *bool            `schema:"-" json:"-"`
	Id              *int             `schema:"-" json:"-"`
}
