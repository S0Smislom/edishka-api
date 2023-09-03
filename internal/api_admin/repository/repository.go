package repository

import "food/internal/api_admin/model"

type User interface {
	Create(data *model.CreateUser) (int, error)
	GetByPhone(phone string) (*model.User, error)
	GetById(id int) (*model.User, error)
}

type Product interface {
	Create(data *model.CreateProduct) (int, error)
	GetById(id int) (*model.Product, error)
	GetList(limit, offset int, filters *model.ProductFilter) ([]*model.Product, error)
	Count(filters *model.ProductFilter) (int, error)
	Update(id int, data *model.UpdateProduct) error
	Delete(id int) error
}

type Repository interface {
	User() User
	Product() Product
}
