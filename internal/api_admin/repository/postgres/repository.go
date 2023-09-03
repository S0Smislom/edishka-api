package postgres

import (
	"database/sql"
	"food/internal/api_admin/repository"
)

type Repository struct {
	db *sql.DB
	*UserRepository
	*ProductRepository
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
