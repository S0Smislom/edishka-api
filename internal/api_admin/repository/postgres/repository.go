package postgres

import (
	"database/sql"
	"food/internal/api_admin/repository"
)

type Repository struct {
	db *sql.DB
	*UserRepository
	*ProductRepository
	*RecipeRepository
	*RecipeStepRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
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
