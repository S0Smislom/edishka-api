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
	recipeStepFileCategory = "recipe-step"
)

type RecipeStepService struct {
	repo        repository.RecipeStep
	recipeRepo  repository.Recipe
	fileService fileservice.FileService
}

func NewRecipeStepService(repo repository.RecipeStep, recipeRepo repository.Recipe, fileService fileservice.FileService) *RecipeStepService {
	return &RecipeStepService{repo: repo, recipeRepo: recipeRepo, fileService: fileService}
}

func (s *RecipeStepService) Create(currentUserId int, data *model.CreateRecipeStep) (*model.RecipeStep, error) {
	// Validate data to create
	if err := data.Validate(); err != nil {
		return nil, &exceptions.ValidationError{Err: err}
	}
	// Check if recipe exists
	dbRecipe, err := s.recipeRepo.GetById(data.RecipeId)
	if err != nil {
		return nil, err
	}
	// Check user permissions
	if dbRecipe.CreatedById != currentUserId {
		return nil, &exceptions.UserPermissionError{}
	}
	// Create resipe step
	now := time.Now().UTC()
	data.CreatedAt = now
	data.UpdatedAt = now
	data.CreatedById = currentUserId
	id, err := s.repo.Create(data)
	if err != nil {
		return nil, err
	}
	// Return Recipe step
	return s.GetById(id)
}

func (s *RecipeStepService) GetById(id int) (*model.RecipeStep, error) {
	return s.repo.GetById(id)
}

func (s *RecipeStepService) GetList(limit, offset int, filters *model.RecipeStepFilter) (*model.RecipeStepList, error) {
	total, err := s.repo.Count(filters)
	if err != nil {
		return nil, err
	}
	result := &model.RecipeStepList{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
	if total == 0 {
		result.Data = []*model.RecipeStep{}
		return result, nil
	}
	data, err := s.repo.GetList(limit, offset, filters)
	if err != nil {
		return nil, err
	}
	result.Data = data
	return result, nil
}

func (s *RecipeStepService) Update(id, currentUserId int, data *model.UpdateRecipeStep) (*model.RecipeStep, error) {
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

func (s *RecipeStepService) Delete(id, currentUserId int) (*model.RecipeStep, error) {
	dbModel, err := s.getAndCheckPermissions(id, currentUserId)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Delete(id); err != nil {
		return nil, err
	}
	return dbModel, nil
}

func (s *RecipeStepService) UploadPhoto(id, currentUserId int, file multipart.File, fileHeader *multipart.FileHeader) (*model.RecipeStep, error) {
	dbModel, err := s.getAndCheckPermissions(id, currentUserId)
	if err != nil {
		return nil, err
	}
	filePrefix := s.getFilePrefix(dbModel.Id)
	filePath, err := s.fileService.UploadFile(filePrefix, file, fileHeader)
	if err != nil {
		return nil, err
	}
	if err := s.repo.UpdatePhoto(id, &filePath); err != nil {
		return nil, err
	}
	dbModel.Photo = &filePath
	return dbModel, nil
}

func (s *RecipeStepService) DeletePhoto(id, currentUserId int) (*model.RecipeStep, error) {
	if _, err := s.getAndCheckPermissions(id, currentUserId); err != nil {
		return nil, err
	}
	if err := s.repo.UpdatePhoto(id, nil); err != nil {
		return nil, err
	}
	return s.repo.GetById(id)
}

func (s *RecipeStepService) getFilePrefix(id int) string {
	// filePrefix := fmt.Sprintf("%s/%d", recipeStepFileCategory, id)
	// return filePrefix
	return recipeStepFileCategory
}

func (s *RecipeStepService) checkPermissions(dbModel *model.RecipeStep, currentUserId int) error {
	// dbRecipe, err := s.recipeRepo.GetById(dbModel.RecipeId)
	// if err != nil {
	// 	return err
	// }
	// if dbRecipe.CreatedById != currentUserId {
	// 	return &exceptions.UserPermissionError{}
	// }
	if dbModel.CreatedById != currentUserId {
		return &exceptions.UserPermissionError{}
	}
	return nil
}

func (s *RecipeStepService) getAndCheckPermissions(id, currentUserId int) (*model.RecipeStep, error) {
	dbModel, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	if err := s.checkPermissions(dbModel, currentUserId); err != nil {
		return nil, err
	}
	return dbModel, nil
}
