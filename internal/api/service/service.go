package service

import (
	"food/internal/api/model"
	"food/internal/api/repository"
	"food/pkg/config"
)

type Auth interface {
	CreateUser(data *model.Login) (*model.LoginResponse, error)
	Login(data *model.LoginConfirm) (*model.LoginConfirmResponse, error)
	ParseToken(accessToken string) (int, error)
}

type User interface{}

type Service struct {
	Auth
	User
}

func NewService(repos repository.Repository, config *config.Config) *Service {
	return &Service{
		Auth: NewAuthService(config.AccessTokenTTL, config.TokenSecret, repos.Auth(), repos.User()),
		User: NewUserService(repos.User()),
	}
}
