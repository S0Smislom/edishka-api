package service

import (
	"food/internal/api_admin/model"
	"food/internal/api_admin/repository"
)

type StepProductService struct {
	repo        repository.StepProduct
	productRepo repository.Product
}

func NewStepProductService(repo repository.StepProduct, productRepo repository.Product) *StepProductService {
	return &StepProductService{repo: repo, productRepo: productRepo}
}

func (s *StepProductService) Create(data *model.CreateStepProduct) (*model.StepProduct, error) {
	if err := data.Validate(); err != nil {
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

func (s *StepProductService) Update(id int, data *model.UpdateStepProduct) (*model.StepProduct, error) {
	if err := s.repo.Update(id, data); err != nil {
		return nil, err
	}
	return s.repo.GetById(id)
}

func (s *StepProductService) Delete(id int) (*model.StepProduct, error) {
	dbModel, err := s.GetById(id)
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
