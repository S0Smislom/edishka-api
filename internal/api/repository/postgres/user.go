package postgres

import (
	"database/sql"
	"errors"
	"food/internal/api/model"
	"strconv"
	"strings"
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
		"select id, phone, first_name, last_name, birthday, code, photo from \"user\" where id=$1",
		item_id,
	).Scan(
		&u.ID,
		&u.Phone,
		&u.FirstName,
		&u.LastName,
		&u.Birthday,
		&u.Code,
		&u.Photo,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("User not found")
		}
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) GetByPhone(phone string) (int, error) {
	var id int
	if err := r.db.QueryRow(
		"select id from \"user\" where phone=$1",
		phone,
	).Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return id, errors.New("User not found")
		}
		return id, err
	}
	return id, nil
}

func (r *UserRepository) UpdateCode(id int, code string) error {
	_, err := r.db.Exec("update \"user\" set code=$1 where id=$2", code, id)
	return err
}

func (r *UserRepository) UpdatePhoto(id int, photo *string) error {
	_, err := r.db.Exec("update \"user\" set photo=$1 where id=$2", photo, id)
	return err
}

// TODO Вынести в pkg
type queryValue []interface{}

func (q queryValue) strLen() string {
	return strconv.Itoa(len(q))
}

func (r *UserRepository) Update(id int, data *model.UpdateUser) error {
	var queryParams []string

	var queryValues queryValue

	if data.Birthday != nil {
		queryValues = append(queryValues, *data.Birthday)
		queryParams = append(queryParams, "birthday=$"+queryValues.strLen())
	}
	if data.FirstName != nil {
		queryValues = append(queryValues, *data.FirstName)
		queryParams = append(queryParams, "first_name=$"+queryValues.strLen())
	}
	if data.LastName != nil {
		queryValues = append(queryValues, *data.LastName)
		queryParams = append(queryParams, "last_name=$"+queryValues.strLen())
	}
	if len(queryValues) == 0 || len(queryParams) == 0 {
		return nil
	}
	queryValues = append(queryValues, id)
	query := "update \"user\" set " + strings.Join(queryParams, ", ") + " where id=$" + queryValues.strLen()
	_, err := r.db.Exec(query, queryValues...)
	return err
}
