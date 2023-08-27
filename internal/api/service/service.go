package service

import "food/internal/api/repository"

type Service struct {
}

func NewService(repos repository.Repository) *Service {
	return &Service{}
}
