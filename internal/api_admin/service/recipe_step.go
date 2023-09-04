package service

import (
	"food/internal/api_admin/model"
	"food/internal/api_admin/repository"
)

type RecipeStepService struct {
	repo repository.RecipeStep
}

func NewRecipeStepService(repo repository.RecipeStep) *RecipeStepService {
	return &RecipeStepService{repo: repo}
}

func (s *RecipeStepService) Create(data *model.CreateRecipeStep) (*model.RecipeStep, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	id, err := s.repo.Create(data)
	if err != nil {
		return nil, err
	}

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

func (s *RecipeStepService) Update(id int, data *model.UpdateRecipeStep) (*model.RecipeStep, error) {
	if err := s.repo.Update(id, data); err != nil {
		return nil, err
	}
	return s.repo.GetById(id)
}

func (s *RecipeStepService) Delete(id int) (*model.RecipeStep, error) {
	dbModel, err := s.GetById(id)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Delete(id); err != nil {
		return nil, err
	}
	return dbModel, nil
}
