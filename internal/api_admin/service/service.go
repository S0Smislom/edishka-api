package service

import (
	"food/internal/api_admin/repository"
	"food/pkg/config"
)

type Service struct {
	*AuthService
	*UserService
}

func NewService(config *config.Config, repo repository.Repository) *Service {
	return &Service{
		AuthService: NewAuthService(config.AdminAccessTokenTTL, config.AdminTokenSecret, repo.User()),
		UserService: NewUserService(repo.User()),
	}
}
