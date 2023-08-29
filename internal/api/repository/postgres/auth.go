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

func (r *AuthRepository) CreateUser(data *model.Login) (*model.LoginResponse, error) {
	user_id := &model.LoginResponse{}
	row := r.db.QueryRow(
		"INSERT INTO \"user\" (phone, first_name, last_name, code) values ($1, $2, $3, $4) RETURNING id",
		data.Phone,
		data.FirstName,
		data.LastName,
		data.Code,
	)
	if err := row.Scan(&user_id.ID); err != nil {
		return nil, err
	}
	return user_id, nil
}
