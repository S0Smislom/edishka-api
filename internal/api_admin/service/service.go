package service

import (
	"food/internal/api_admin/repository"
	"food/pkg/config"
)

type Service struct {
	*AuthService
	*UserService
	*ProductService
	*RecipeService
	*RecipeStepService
	*StepProductService
}

func NewService(config *config.Config, repo repository.Repository) *Service {
	return &Service{
		AuthService:        NewAuthService(config.AdminAccessTokenTTL, config.AdminTokenSecret, repo.User()),
		UserService:        NewUserService(repo.User()),
		ProductService:     NewProductService(repo.Product()),
		RecipeService:      NewRecipeService(repo.Recipe()),
		RecipeStepService:  NewRecipeStepService(repo.RecipeStep()),
		StepProductService: NewStepProductService(repo.StepProduct(), repo.Product()),
	}
}
