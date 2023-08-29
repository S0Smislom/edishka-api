package service

import (
	"food/internal/api/model"
	"food/internal/api/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetById(item_id int) (*model.User, error) {
	return s.repo.GetById(item_id)
}
