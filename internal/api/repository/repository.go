package repository

import "food/internal/api/model"

type Auth interface {
	CreateUser(data *model.Login) (int, error)
}

type User interface {
	GetById(itemId int) (*model.User, error)
	Update(id int, data *model.UpdateUser) error
	GetByPhone(phone string) (int, error)
	UpdateCode(id int, code string) error
	UpdatePhoto(id int, photo *string) error
}

type Repository interface {
	Auth() Auth
	User() User
}
