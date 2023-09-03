package service

import (
	"food/internal/api_admin/model"
	"food/internal/api_admin/repository"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(data *model.CreateProduct) (*model.Product, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	id, err := s.repo.Create(data)
	if err != nil {
		return nil, err
	}

	return s.GetById(id)
}

func (s *ProductService) GetById(id int) (*model.Product, error) {
	return s.repo.GetById(id)
}

func (s *ProductService) GetList(limit, offset int, filters *model.ProductFilter) (*model.ProductList, error) {
	total, err := s.repo.Count(filters)
	if err != nil {
		return nil, err
	}
	result := &model.ProductList{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
	if total == 0 {
		result.Data = []*model.Product{}
		return result, nil
	}
	data, err := s.repo.GetList(limit, offset, filters)
	if err != nil {
		return nil, err
	}
	result.Data = data
	return result, nil
}

func (s *ProductService) Update(id int, data *model.UpdateProduct) (*model.Product, error) {
	if err := s.repo.Update(id, data); err != nil {
		return nil, err
	}
	return s.repo.GetById(id)
}

func (s *ProductService) Delete(id int) (*model.Product, error) {
	dbModel, err := s.GetById(id)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Delete(id); err != nil {
		return nil, err
	}
	return dbModel, nil
}
