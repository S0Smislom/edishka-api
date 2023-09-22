package service

import (
	"food/internal/api_admin/repository"
	fileservice "food/internal/file_service"
	"food/pkg/config"
)

type Service struct {
	*AuthService
	*UserService
	*ProductService
	*RecipeService
	*RecipeStepService
	*RecipeGalleryService
	*StepProductService
}

func NewService(config *config.Config, repo repository.Repository, fileService fileservice.FileService) *Service {
	return &Service{
		AuthService:          NewAuthService(config.AdminAccessTokenTTL, config.AdminTokenSecret, repo.User()),
		UserService:          NewUserService(repo.User()),
		ProductService:       NewProductService(repo.Product(), fileService),
		RecipeService:        NewRecipeService(repo.Recipe(), repo.RecipeStep(), repo.StepProduct(), repo.Product(), fileService),
		RecipeStepService:    NewRecipeStepService(repo.RecipeStep(), fileService),
		RecipeGalleryService: NewRecipeGalleryService(repo.RecipeGallery(), repo.Recipe(), fileService),
		StepProductService:   NewStepProductService(repo.StepProduct(), repo.Product()),
	}
}
