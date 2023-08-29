package repository

import "food/internal/api/model"

type Auth interface {
	CreateUser(data *model.Login) (*model.LoginResponse, error)
}

type User interface {
	GetById(itemId int) (*model.User, error)
}

type Repository interface {
	Auth() Auth
	User() User
}
