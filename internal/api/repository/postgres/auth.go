package postgres

import (
	"database/sql"
	"food/internal/api/model"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(data *model.Login) (int, error) {
	// user_id := &model.LoginResponse{}
	var userId int
	row := r.db.QueryRow(
		"INSERT INTO \"user\" (phone, code) values ($1, $2) RETURNING id",
		data.Phone,
		data.Code,
	)
	if err := row.Scan(&userId); err != nil {
		return userId, err
	}
	return userId, nil
}
