package repository

import "food/internal/api/model"

type Auth interface {
	CreateUser(data *model.Login) (int, error)
}

type User interface {
	GetById(itemId int) (*model.User, error)
	Update(id int, data *model.UpdateUser) error
	GetByPhone(phone string) (int, error)
	UpdateCode(id int, code string) error
	UpdatePhoto(id int, photo *string) error
}

type Product interface {
	GetById(id int) (*model.Product, error)
	GetList(limit, offset int, filters *model.ProductFilter) ([]*model.Product, error)
	Create(data *model.CreateProduct) (int, error)
	Count(filters *model.ProductFilter) (int, error)
	Delete(id int) error
	Update(id int, data *model.UpdateProduct) error
	UpdatePhoto(id int, photo *string) error
}

type Recipe interface {
	Create(data *model.CreateRecipe) (int, error)
	GetById(id int) (*model.Recipe, error)
	GetList(limit, offset int, filters *model.RecipeFilter) ([]*model.Recipe, error)
	Count(filters *model.RecipeFilter) (int, error)
	Update(id int, data *model.UpdateRecipe) error
	Delete(id int) error
	GetOne(filters *model.RecipeFilter) (*model.Recipe, error)
}

type RecipeStep interface {
	Create(data *model.CreateRecipeStep) (int, error)
	GetById(id int) (*model.RecipeStep, error)
	GetList(limit, offset int, filters *model.RecipeStepFilter) ([]*model.RecipeStep, error)
	Count(filters *model.RecipeStepFilter) (int, error)
	Update(id int, data *model.UpdateRecipeStep) error
	Delete(id int) error
	UpdatePhoto(id int, photo *string) error
}

type RecipeGallery interface {
	Create(data *model.CreateRecipeGallery) (int, error)
	GetById(id int) (*model.RecipeGallery, error)
	GetList(limit, offset int, filters *model.RecipeGalleryFilter) ([]*model.RecipeGallery, error)
	Count(filters *model.RecipeGalleryFilter) (int, error)
	Update(id int, data *model.UpdateRecipeGallery) error
	Delete(id int) error
}

type RecipeProduct interface {
	Create(data *model.CreateRecipeProduct) (int, error)
	GetById(id int) (*model.RecipeProduct, error)
	GetList(limit, offset int, filters *model.RecipeProductFilter) ([]*model.RecipeProduct, error)
	Count(filters *model.RecipeProductFilter) (int, error)
	Update(id int, data *model.UpdateRecipeProduct) error
	Delete(id int) error
}

type ShoppingItem interface {
	Create(data *model.CreateShoppingItem) (int, error)
	GetById(id int) (*model.ShoppingItem, error)
	GetList(limit, offset int, filters *model.ShoppingItemFilter) ([]*model.ShoppingItem, error)
	Count(filters *model.ShoppingItemFilter) (int, error)
	Update(id int, data *model.UpdateShoppingItem) error
	Delete(id int) error
}

type Repository interface {
	Auth() Auth
	User() User
	Product() Product
	Recipe() Recipe
	RecipeStep() RecipeStep
	RecipeGallery() RecipeGallery
	RecipeProduct() RecipeProduct
	ShoppingItem() ShoppingItem
}
