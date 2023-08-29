package postgres

import (
	"database/sql"
	"errors"
	"food/internal/api/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetById(item_id int) (*model.User, error) {
	u := &model.User{}
	if err := r.db.QueryRow(
		"select id, phone, first_name, last_name, birthday, code from \"user\" where id=$1",
		item_id,
	).Scan(
		&u.ID,
		&u.Phone,
		&u.FirstName,
		&u.LastName,
		&u.Birthday,
		&u.Code,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("User not found")
		}
		return nil, err
	}
	return u, nil
}
