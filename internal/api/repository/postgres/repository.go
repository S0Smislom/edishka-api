package postgres

import (
	"database/sql"
	"food/internal/api/repository"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) repository.Repository {
	return &Repository{
		db: db,
	}
}
