package service

import (
	"fmt"
	"food/internal/api_admin/model"
	"food/internal/api_admin/repository"
	fileservice "food/internal/file_service"
	"mime/multipart"
)

const (
	fileCategory = "product"
)

type ProductService struct {
	repo        repository.Product
	fileService fileservice.FileService
}

func NewProductService(repo repository.Product, fileService fileservice.FileService) *ProductService {
	return &ProductService{repo: repo, fileService: fileService}
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

func (s *ProductService) UploadPhoto(id int, file multipart.File, fileHeader *multipart.FileHeader) (*model.Product, error) {
	filePrefix := s.getFilePrefix(id)
	filePath, err := s.fileService.UploadFile(filePrefix, file, fileHeader)
	if err != nil {
		return nil, err
	}
	if err := s.repo.UpdatePhoto(id, &filePath); err != nil {
		return nil, err
	}
	return s.repo.GetById(id)
}

func (s *ProductService) DeletePhoto(id int) (*model.Product, error) {
	if err := s.repo.UpdatePhoto(id, nil); err != nil {
		return nil, err
	}
	return s.repo.GetById(id)
}

func (s *ProductService) getFilePrefix(id int) string {
	filePrefix := fmt.Sprintf("/%s/%d", fileCategory, id)
	return filePrefix
}
