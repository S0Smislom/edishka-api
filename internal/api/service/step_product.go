package service

import (
	"food/internal/api/model"
	"food/internal/api/repository"
	"food/pkg/exceptions"
)

type StepProductService struct {
	repo           repository.StepProduct
	productRepo    repository.Product
	recipeStepRepo repository.RecipeStep
}

func NewStepProductService(repo repository.StepProduct, productRepo repository.Product, recipeStepRepo repository.RecipeStep) *StepProductService {
	return &StepProductService{repo: repo, productRepo: productRepo, recipeStepRepo: recipeStepRepo}
}

func (s *StepProductService) Create(currentUserId int, data *model.CreateStepProduct) (*model.StepProduct, error) {
	if err := data.Validate(); err != nil {
		return nil, &exceptions.ValidationError{Err: err}
	}
	// Check permissions
	dbRecipeStep, err := s.recipeStepRepo.GetById(data.RecipeStepId)
	if err != nil {
		return nil, err
	}
	if dbRecipeStep.CreatedById != currentUserId {
		return nil, &exceptions.UserPermissionError{}
	}
	data.CreatedById = currentUserId
	// Check if product exists
	_, err = s.productRepo.GetById(data.ProductId)
	if err != nil {
		return nil, err
	}

	id, err := s.repo.Create(data)
	if err != nil {
		return nil, err
	}

	return s.GetById(id)
}

func (s *StepProductService) GetById(id int) (*model.StepProduct, error) {
	dbModel, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	dbProduct, err := s.productRepo.GetById(dbModel.ProductId)
	if err != nil {
		return nil, err
	}
	dbModel.Product = dbProduct
	return dbModel, nil
}

func (s *StepProductService) GetList(limit, offset int, filters *model.StepProductFilter) (*model.StepProductList, error) {
	total, err := s.repo.Count(filters)
	if err != nil {
		return nil, err
	}
	result := &model.StepProductList{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
	if total == 0 {
		result.Data = []*model.StepProduct{}
		return result, nil
	}
	data, err := s.repo.GetList(limit, offset, filters)
	if err != nil {
		return nil, err
	}
	s.fetchProducts(data)
	result.Data = data
	return result, nil
}

func (s *StepProductService) Update(id, currentUserId int, data *model.UpdateStepProduct) (*model.StepProduct, error) {
	if err := data.Validate(); err != nil {
		return nil, &exceptions.ValidationError{Err: err}
	}
	if _, err := s.getAndCheckPermissions(id, currentUserId); err != nil {
		return nil, err
	}
	if err := s.repo.Update(id, data); err != nil {
		return nil, err
	}
	return s.repo.GetById(id)
}

func (s *StepProductService) Delete(id, currentUserId int) (*model.StepProduct, error) {
	dbModel, err := s.getAndCheckPermissions(id, currentUserId)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Delete(id); err != nil {
		return nil, err
	}
	return dbModel, nil
}

func (s *StepProductService) fetchProducts(dbModels []*model.StepProduct) {
	for _, dbModel := range dbModels {
		dbProduct, _ := s.productRepo.GetById(dbModel.ProductId)
		dbModel.Product = dbProduct
	}
}

func (s *StepProductService) getAndCheckPermissions(id, currentUserId int) (*model.StepProduct, error) {
	dbModel, err := s.GetById(id)
	if err != nil {
		return nil, err
	}
	if err := s.checkPermissions(dbModel, currentUserId); err != nil {
		return nil, err
	}
	return dbModel, nil
}

func (s *StepProductService) checkPermissions(dbModel *model.StepProduct, currentUserId int) error {
	if dbModel.CreatedById != currentUserId {
		return &exceptions.UserPermissionError{}
	}
	return nil
}
