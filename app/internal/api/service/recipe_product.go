package service

import (
	"food/internal/api/model"
	"food/internal/api/repository"
	"food/pkg/exceptions"
	"time"
)

type RecipeProductService struct {
	repo        repository.RecipeProduct
	productRepo repository.Product
	recipeRepo  repository.Recipe
}

func NewRecipeProductService(repo repository.RecipeProduct, productRepo repository.Product, recipeRepo repository.Recipe) *RecipeProductService {
	return &RecipeProductService{repo: repo, productRepo: productRepo, recipeRepo: recipeRepo}
}

func (s *RecipeProductService) Create(currentUserId int, data *model.CreateRecipeProduct) (*model.RecipeProduct, error) {
	if err := data.Validate(); err != nil {
		return nil, &exceptions.ValidationError{Err: err}
	}
	// Check permissions
	dbRecipe, err := s.recipeRepo.GetById(data.RecipeId)
	if err != nil {
		return nil, err
	}
	if dbRecipe.CreatedById != currentUserId {
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

func (s *RecipeProductService) GetById(id int) (*model.RecipeProduct, error) {
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

func (s *RecipeProductService) GetList(limit, offset int, filters *model.RecipeProductFilter) (*model.RecipeProductList, error) {
	total, err := s.repo.Count(filters)
	if err != nil {
		return nil, err
	}
	result := &model.RecipeProductList{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
	if total == 0 {
		result.Data = []*model.RecipeProduct{}
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

func (s *RecipeProductService) Update(id, currentUserId int, data *model.UpdateRecipeProduct) (*model.RecipeProduct, error) {
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
	return s.GetById(id)
}

func (s *RecipeProductService) Delete(id, currentUserId int) (*model.RecipeProduct, error) {
	dbModel, err := s.getAndCheckPermissions(id, currentUserId)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Delete(id); err != nil {
		return nil, err
	}
	return dbModel, nil
}

func (s *RecipeProductService) fetchProducts(dbModels []*model.RecipeProduct) {
	for _, dbModel := range dbModels {
		dbProduct, _ := s.productRepo.GetById(dbModel.ProductId)
		dbModel.Product = dbProduct
	}
}

func (s *RecipeProductService) getAndCheckPermissions(id, currentUserId int) (*model.RecipeProduct, error) {
	dbModel, err := s.GetById(id)
	if err != nil {
		return nil, err
	}
	if err := s.checkPermissions(dbModel, currentUserId); err != nil {
		return nil, err
	}
	return dbModel, nil
}

func (s *RecipeProductService) checkPermissions(dbModel *model.RecipeProduct, currentUserId int) error {
	if dbModel.CreatedById != currentUserId {
		return &exceptions.UserPermissionError{}
	}
	return nil
}
