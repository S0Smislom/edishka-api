package service

import (
	"fmt"
	"food/internal/api_admin/model"
	"food/internal/api_admin/repository"
	fileservice "food/internal/file_service"
	"mime/multipart"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	recipeGalleryCategory = "recipe-gallery"
)

type RecipeGalleryService struct {
	repo        repository.RecipeGallery
	recipeRepo  repository.Recipe
	fileService fileservice.FileService
}

func NewRecipeGalleryService(repo repository.RecipeGallery, recipeRepo repository.Recipe, fileService fileservice.FileService) *RecipeGalleryService {
	return &RecipeGalleryService{
		repo:        repo,
		recipeRepo:  recipeRepo,
		fileService: fileService,
	}
}

func (s *RecipeGalleryService) Create(data *model.CreateRecipeGallery, file multipart.File, fileHeader *multipart.FileHeader) (*model.RecipeGallery, error) {
	if err := data.Validate(); err != nil {
		if e, ok := err.(validation.InternalError); ok {
			// an internal error happened
			return nil, e.InternalError()
		}
		return nil, err
	}
	fmt.Println(data)
	dbRecipe, err := s.recipeRepo.GetById(data.RecipeId)
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	filePrefix := s.getFilePrefix(dbRecipe.Slug)
	data.Photo, err = s.fileService.UploadFile(filePrefix, file, fileHeader)
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	id, err := s.repo.Create(data)
	if err != nil {
		return nil, err
	}
	return s.GetById(id)
}

func (s *RecipeGalleryService) GetById(id int) (*model.RecipeGallery, error) {
	return s.repo.GetById(id)
}

func (s *RecipeGalleryService) Update(id int, data *model.UpdateRecipeGallery) (*model.RecipeGallery, error) {
	if err := s.repo.Update(id, data); err != nil {
		return nil, err
	}
	return s.repo.GetById(id)
}

func (s *RecipeGalleryService) Delete(id int) (*model.RecipeGallery, error) {
	dbModel, err := s.GetById(id)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Delete(id); err != nil {
		return nil, err
	}
	return dbModel, nil
}

func (s *RecipeGalleryService) getFilePrefix(recipeSlug string) string {
	filePrefix := fmt.Sprintf("/%s/%s", recipeGalleryCategory, recipeSlug)
	return filePrefix
}
