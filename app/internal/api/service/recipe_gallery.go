package service

import (
	"food/internal/api/model"
	"food/internal/api/repository"
	fileservice "food/internal/file_service"
	"food/pkg/exceptions"
	"mime/multipart"
	"time"
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

func (s *RecipeGalleryService) Create(currentUserId int, data *model.CreateRecipeGallery, file multipart.File, fileHeader *multipart.FileHeader) (*model.RecipeGallery, error) {
	if err := data.Validate(); err != nil {
		return nil, &exceptions.ValidationError{Err: err}
	}
	dbRecipe, err := s.recipeRepo.GetById(data.RecipeId)
	if err != nil {
		return nil, err
	}
	if dbRecipe.CreatedById != currentUserId {
		return nil, &exceptions.UserPermissionError{}
	}
	now := time.Now().UTC()
	data.CreatedAt = now
	data.UpdatedAt = now
	data.CreatedById = currentUserId
	filePrefix := s.getFilePrefix(dbRecipe.Slug)
	data.Photo, err = s.fileService.UploadFile(filePrefix, file, fileHeader)
	if err != nil {
		return nil, err
	}
	id, err := s.repo.Create(data)
	if err != nil {
		return nil, err
	}
	return s.GetById(id)
}

func (s *RecipeGalleryService) GetById(id int) (*model.RecipeGallery, error) {
	return s.repo.GetById(id)
}

func (s *RecipeGalleryService) Update(id, currentUserId int, data *model.UpdateRecipeGallery) (*model.RecipeGallery, error) {
	if err := data.Validate(); err != nil {
		return nil, &exceptions.ValidationError{Err: err}
	}
	if _, err := s.getAndCheckPermissions(id, currentUserId); err != nil {
		return nil, err
	}
	now := time.Now().UTC()
	data.UpdatedAt = now
	if err := s.repo.Update(id, data); err != nil {
		return nil, err
	}
	return s.repo.GetById(id)
}

func (s *RecipeGalleryService) Delete(id, currentUserId int) (*model.RecipeGallery, error) {
	dbModel, err := s.getAndCheckPermissions(id, currentUserId)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Delete(id); err != nil {
		return nil, err
	}
	return dbModel, nil
}

func (s *RecipeGalleryService) getFilePrefix(recipeSlug string) string {
	// filePrefix := fmt.Sprintf("%s/%s", recipeGalleryCategory, recipeSlug)
	// return filePrefix
	return recipeGalleryCategory
}

func (s *RecipeGalleryService) getAndCheckPermissions(id, currentUserId int) (*model.RecipeGallery, error) {
	dbModel, err := s.GetById(id)
	if err != nil {
		return nil, err
	}
	if err := s.checkPermissions(dbModel, currentUserId); err != nil {
		return nil, err
	}
	return dbModel, nil
}

func (s *RecipeGalleryService) checkPermissions(dbModel *model.RecipeGallery, currentUserId int) error {
	if dbModel.CreatedById != currentUserId {
		return &exceptions.UserPermissionError{}
	}
	return nil
}
