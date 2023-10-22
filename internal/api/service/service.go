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
	Refresh(userId int) (*model.LoginConfirmResponse, error)
	ParseToken(accessToken string) (*model.TokenClaims, error)
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

type Recipe interface {
	GetById(id int) (*model.Recipe, error)
	GetList(limit, offset int, filters *model.RecipeFilter) (*model.RecipeList, error)
	Create(data *model.CreateRecipe) (*model.Recipe, error)
	Update(id int, currentUserId int, data *model.UpdateRecipe) (*model.Recipe, error)
	Delete(id int, currentUserId int) (*model.Recipe, error)
	GetListPrivate(limit, offset, currentUserId int, filters *model.RecipeFilter) (*model.RecipeList, error)
	GetByIdPrivate(id, currentUserId int) (*model.Recipe, error)
}

type RecipeStep interface {
	GetById(id int) (*model.RecipeStep, error)
	Create(currentUserId int, data *model.CreateRecipeStep) (*model.RecipeStep, error)
	GetList(limit, offset int, filters *model.RecipeStepFilter) (*model.RecipeStepList, error)
	Update(id, currentUserId int, data *model.UpdateRecipeStep) (*model.RecipeStep, error)
	Delete(id, currentUserId int) (*model.RecipeStep, error)
	UploadPhoto(id, currentUserId int, file multipart.File, fileHeader *multipart.FileHeader) (*model.RecipeStep, error)
	DeletePhoto(id, currentUserId int) (*model.RecipeStep, error)
}

type StepProduct interface {
	GetById(id int) (*model.StepProduct, error)
	Create(currentUserId int, data *model.CreateStepProduct) (*model.StepProduct, error)
	GetList(limit, offset int, filters *model.StepProductFilter) (*model.StepProductList, error)
	Update(id, currentUserId int, data *model.UpdateStepProduct) (*model.StepProduct, error)
	Delete(id, currentUserId int) (*model.StepProduct, error)
}

type RecipeGallery interface {
	Create(currentUserId int, data *model.CreateRecipeGallery, file multipart.File, fileHeader *multipart.FileHeader) (*model.RecipeGallery, error)
	GetById(id int) (*model.RecipeGallery, error)
	Update(id, currentUserId int, data *model.UpdateRecipeGallery) (*model.RecipeGallery, error)
	Delete(id, currentUserId int) (*model.RecipeGallery, error)
}

type Service struct {
	Auth
	User
	Product
	Recipe
	RecipeStep
	StepProduct
	RecipeGallery
}

func NewService(repos repository.Repository, fileService fileservice.FileService, config *config.Config) *Service {
	return &Service{
		Auth:          NewAuthService(config.AccessTokenTTL, config.RefreshTokenTTL, config.TokenSecret, repos.Auth(), repos.User()),
		User:          NewUserService(repos.User(), fileService),
		Product:       NewProductService(repos.Product(), fileService),
		Recipe:        NewRecipeService(repos.Recipe(), repos.RecipeStep(), repos.StepProduct(), repos.Product(), fileService),
		RecipeStep:    NewRecipeStepService(repos.RecipeStep(), repos.Recipe(), fileService),
		StepProduct:   NewStepProductService(repos.StepProduct(), repos.Product(), repos.RecipeStep()),
		RecipeGallery: NewRecipeGalleryService(repos.RecipeGallery(), repos.Recipe(), fileService),
	}
}
