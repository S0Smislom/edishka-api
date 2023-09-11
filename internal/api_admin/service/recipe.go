package service

import (
	"food/internal/api_admin/model"
	"food/internal/api_admin/repository"
)

type RecipeService struct {
	repo            repository.Recipe
	recipeStepRepo  repository.RecipeStep
	productRepo     repository.Product
	stepProductRepo repository.StepProduct
}

func NewRecipeService(
	repo repository.Recipe,
	recipeStepRepo repository.RecipeStep,
	stepProductRepo repository.StepProduct,
	productRepo repository.Product,
) *RecipeService {
	return &RecipeService{
		repo:            repo,
		recipeStepRepo:  recipeStepRepo,
		stepProductRepo: stepProductRepo,
		productRepo:     productRepo,
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
	return s.repo.GetById(id)
}

func (s *RecipeService) GetList(limit, offset int, filters *model.RecipeFilter) (*model.RecipeList, error) {
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

func (s *RecipeService) Update(id int, data *model.UpdateRecipe) (*model.Recipe, error) {
	if err := s.repo.Update(id, data); err != nil {
		return nil, err
	}
	return s.repo.GetById(id)
}

func (s *RecipeService) Delete(id int) (*model.Recipe, error) {
	dbModel, err := s.GetById(id)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Delete(id); err != nil {
		return nil, err
	}
	return dbModel, nil
}
