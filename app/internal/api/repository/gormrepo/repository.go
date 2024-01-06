package gormrepo

import (
	"food/internal/api/repository"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
	*AuthRepository
	*UserRepository
	*ProductRepository
	*RecipeRepository
	*RecipeStepRepository
	*RecipeProductRepository
	*RecipeGalleryRepository
	*ShoppingItemRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Auth() repository.Auth {
	if r.AuthRepository == nil {
		r.AuthRepository = NewAuthRepository(r.db)
	}
	return r.AuthRepository
}

func (r *Repository) User() repository.User {
	if r.UserRepository == nil {
		r.UserRepository = NewUserRepository(r.db)
	}
	return r.UserRepository
}

func (r *Repository) Product() repository.Product {
	if r.ProductRepository == nil {
		r.ProductRepository = NewProductRepository(r.db)
	}
	return r.ProductRepository
}

func (r *Repository) Recipe() repository.Recipe {
	if r.RecipeRepository == nil {
		r.RecipeRepository = NewRecipeRepository(r.db)
	}
	return r.RecipeRepository
}

func (r *Repository) RecipeStep() repository.RecipeStep {
	if r.RecipeStepRepository == nil {
		r.RecipeStepRepository = NewRecipeStepRepository(r.db)
	}
	return r.RecipeStepRepository
}
func (r *Repository) RecipeGallery() repository.RecipeGallery {
	if r.RecipeGalleryRepository == nil {
		r.RecipeGalleryRepository = NewRecipeGalleryRepository(r.db)
	}
	return r.RecipeGalleryRepository
}

func (r *Repository) RecipeProduct() repository.RecipeProduct {
	if r.RecipeProductRepository == nil {
		r.RecipeProductRepository = NewRecipeProductRepository(r.db)
	}
	return r.RecipeProductRepository
}

func (r *Repository) ShoppingItem() repository.ShoppingItem {
	if r.ShoppingItemRepository == nil {
		r.ShoppingItemRepository = NewShoppingItemRepository(r.db)
	}
	return r.ShoppingItemRepository
}
