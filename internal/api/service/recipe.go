package service

import (
	"food/internal/api/model"
	"food/internal/api/repository"
	fileservice "food/internal/file_service"
	"food/pkg/exceptions"
)

const (
	recipeFileCategory = "recipe"
)

type RecipeService struct {
	repo            repository.Recipe
	recipeStepRepo  repository.RecipeStep
	productRepo     repository.Product
	stepProductRepo repository.StepProduct
	fileService     fileservice.FileService
}

func NewRecipeService(
	repo repository.Recipe,
	recipeStepRepo repository.RecipeStep,
	stepProductRepo repository.StepProduct,
	productRepo repository.Product,
	fileService fileservice.FileService,
) *RecipeService {
	return &RecipeService{
		repo:            repo,
		recipeStepRepo:  recipeStepRepo,
		stepProductRepo: stepProductRepo,
		productRepo:     productRepo,
		fileService:     fileService,
	}
}

func (s *RecipeService) Create(data *model.CreateRecipe) (*model.Recipe, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}
	id, err := s.repo.Create(data)
	if err != nil {
		return nil, err
	}
	return s.GetById(id)
}

func (s *RecipeService) GetById(id int) (*model.Recipe, error) {
	published := true
	return s.repo.GetOne(&model.RecipeFilter{Id: &id, Published: &published})
}

func (s *RecipeService) GetByIdPrivate(id, currentUserId int) (*model.Recipe, error) {
	return s.repo.GetOne(&model.RecipeFilter{Id: &id, CreatedById: &currentUserId})
}

func (s *RecipeService) GetList(limit, offset int, filters *model.RecipeFilter) (*model.RecipeList, error) {
	published := true
	filters.Published = &published
	return s.getList(limit, offset, filters)
}

func (s *RecipeService) GetListPrivate(limit, offset, currentUserId int, filters *model.RecipeFilter) (*model.RecipeList, error) {
	filters.CreatedById = &currentUserId
	return s.getList(limit, offset, filters)
}

func (s *RecipeService) Update(id, currentUserId int, data *model.UpdateRecipe) (*model.Recipe, error) {
	if _, err := s.getAndCheckPermissions(id, currentUserId); err != nil {
		return nil, err
	}
	if err := s.repo.Update(id, data); err != nil {
		return nil, err
	}
	return s.repo.GetById(id)
}

func (s *RecipeService) Delete(id, currentUserId int) (*model.Recipe, error) {
	dbModel, err := s.getAndCheckPermissions(id, currentUserId)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Delete(id); err != nil {
		return nil, err
	}
	return dbModel, nil
}

func (s *RecipeService) checkPermissions(dbModel *model.Recipe, currentUserId int) error {
	if dbModel.CreatedById != currentUserId {
		return &exceptions.UserPermissionError{Msg: "Forbidden"}
	}
	return nil
}

func (s *RecipeService) getAndCheckPermissions(id, currentUserId int) (*model.Recipe, error) {
	dbModel, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	if err := s.checkPermissions(dbModel, currentUserId); err != nil {
		return nil, err
	}
	return dbModel, nil
}

func (s *RecipeService) getList(limit, offset int, filters *model.RecipeFilter) (*model.RecipeList, error) {
	total, err := s.repo.Count(filters)
	if err != nil {
		return nil, err
	}
	result := &model.RecipeList{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
	if total == 0 {
		result.Data = []*model.Recipe{}
		return result, nil
	}
	data, err := s.repo.GetList(limit, offset, filters)
	if err != nil {
		return nil, err
	}
	result.Data = data
	return result, nil
}
