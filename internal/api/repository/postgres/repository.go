package postgres

import (
	"database/sql"
	"food/internal/api/repository"
)

type Repository struct {
	db *sql.DB
	*AuthRepository
	*UserRepository
}

func NewRepository(db *sql.DB) *Repository {
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
