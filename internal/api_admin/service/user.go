package service

import (
	"food/internal/api_admin/model"
	"food/internal/api_admin/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetById(userId int) (*model.User, error) {
	return s.repo.GetById(userId)
}
