package repository

import "food/internal/api_admin/model"

type User interface {
	Create(data *model.CreateUser) (int, error)
	GetByPhone(phone string) (*model.User, error)
	GetById(id int) (*model.User, error)
}

type Product interface {
	Create(data *model.CreateProduct) (int, error)
	GetById(id int) (*model.Product, error)
	GetList(limit, offset int, filters *model.ProductFilter) ([]*model.Product, error)
	Count(filters *model.ProductFilter) (int, error)
	Update(id int, data *model.UpdateProduct) error
	Delete(id int) error
}

type Recipe interface {
	Create(data *model.CreateRecipe) (int, error)
	GetById(id int) (*model.Recipe, error)
	GetList(limit, offset int, filters *model.RecipeFilter) ([]*model.Recipe, error)
	Count(filters *model.RecipeFilter) (int, error)
	Update(id int, data *model.UpdateRecipe) error
	Delete(id int) error
}

type RecipeStep interface {
	Create(data *model.CreateRecipeStep) (int, error)
	GetById(id int) (*model.RecipeStep, error)
	GetList(limit, offset int, filters *model.RecipeStepFilter) ([]*model.RecipeStep, error)
	Count(filters *model.RecipeStepFilter) (int, error)
	Update(id int, data *model.UpdateRecipeStep) error
	Delete(id int) error
}

type Repository interface {
	User() User
	Product() Product
	Recipe() Recipe
	RecipeStep() RecipeStep
}
