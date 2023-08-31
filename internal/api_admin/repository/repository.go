package repository

import "food/internal/api_admin/model"

type User interface {
	Create(data *model.CreateUser) (int, error)
	GetByPhone(phone string) (*model.User, error)
	GetById(id int) (*model.User, error)
}

type Repository interface {
	User() User
}
