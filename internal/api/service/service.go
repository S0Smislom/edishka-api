package service

import (
	"food/internal/api/model"
	"food/internal/api/repository"
	fileservice "food/internal/file_service"
	"food/pkg/config"
	"mime/multipart"
)

type Auth interface {
	CreateUser(data *model.Login) (*model.LoginResponse, error)
	Login(data *model.LoginConfirm) (*model.LoginConfirmResponse, error)
	ParseToken(accessToken string) (int, error)
}

type User interface {
	GetById(id int) (*model.User, error)
	Update(id int, data *model.UpdateUser) (*model.User, error)
	UploadPhoto(id int, file multipart.File, fileHeader *multipart.FileHeader) (*model.User, error)
	DeletePhoto(id int) (*model.User, error)
}

type Product interface {
	GetById(id int) (*model.Product, error)
	GetList(limit, offset int, filters *model.ProductFilter) (*model.ProductList, error)
	Create(data *model.CreateProduct) (*model.Product, error)
	Update(id int, currentUserId int, data *model.UpdateProduct) (*model.Product, error)
	Delete(id int, currentUserId int) (*model.Product, error)
	UploadPhoto(id int, currentUserId int, file multipart.File, fileHeader *multipart.FileHeader) (*model.Product, error)
	DeletePhoto(id int, currentUserId int) (*model.Product, error)
}

type Service struct {
	Auth
	User
	Product
}

func NewService(repos repository.Repository, fileService fileservice.FileService, config *config.Config) *Service {
	return &Service{
		Auth:    NewAuthService(config.AccessTokenTTL, config.TokenSecret, repos.Auth(), repos.User()),
		User:    NewUserService(repos.User(), fileService),
		Product: NewProductService(repos.Product(), fileService),
	}
}
