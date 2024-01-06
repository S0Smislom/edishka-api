package postgres

import (
	"database/sql"
	"food/internal/api_admin/model"
	"food/pkg/exceptions"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByPhone(phone string) (*model.User, error) {
	u := &model.User{}
	if err := r.db.QueryRow(
		"select id, phone, first_name, last_name, birthday, created_at, updated_at, is_superuser, is_staff, password from \"user\" where phone=$1;",
		phone,
	).Scan(
		&u.Id,
		&u.Phone,
		&u.FirstName,
		&u.LastName,
		&u.Birthday,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.IsSuperuser,
		&u.IsStaff,
		&u.Password,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, &exceptions.ObjectNotFoundError{Msg: "User not found"}
		}
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) GetById(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.db.QueryRow(
		"select id, phone, first_name, last_name, birthday, created_at, updated_at, is_superuser, is_staff, password from \"user\" where id=$1;",
		id,
	).Scan(
		&u.Id,
		&u.Phone,
		&u.FirstName,
		&u.LastName,
		&u.Birthday,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.IsSuperuser,
		&u.IsStaff,
		&u.Password,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, &exceptions.ObjectNotFoundError{Msg: "User not found"}
		}
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) Create(data *model.CreateUser) (int, error) {
	var id int
	row := r.db.QueryRow(
		"INSERT INTO \"user\" (phone, password, is_superuser, is_staff) values ($1, $2, $3, $4) RETURNING id",
		data.Phone,
		data.Password,
		data.IsSuperuser,
		data.IsStaff,
	)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
