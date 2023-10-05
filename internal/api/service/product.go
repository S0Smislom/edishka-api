package service

import (
	"fmt"
	"food/internal/api/model"
	"food/internal/api/repository"
	fileservice "food/internal/file_service"
	"food/pkg/exceptions"
	"mime/multipart"
)

const (
	// TODO Вынести в отдельный package
	productFileCategory = "product"
)

type ProductService struct {
	repo        repository.Product
	fileService fileservice.FileService
}

func NewProductService(repo repository.Product, fileService fileservice.FileService) *ProductService {
	return &ProductService{repo: repo, fileService: fileService}
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
func (s *ProductService) Update(id int, currentUserId int, data *model.UpdateProduct) (*model.Product, error) {
	_, err := s.getAndCheckPermissions(id, currentUserId)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Update(id, data); err != nil {
		return nil, err
	}
	return s.repo.GetById(id)
}
func (s *ProductService) Delete(id int, currentUserId int) (*model.Product, error) {
	dbModel, err := s.getAndCheckPermissions(id, currentUserId)
	if err != nil {
		return nil, err
	}
	s.repo.Delete(id)
	return dbModel, nil
}
func (s *ProductService) UploadPhoto(id int, currentUserId int, file multipart.File, fileHeader *multipart.FileHeader) (*model.Product, error) {
	dbModel, err := s.getAndCheckPermissions(id, currentUserId)
	if err != nil {
		return nil, err
	}
	filePrefix := s.getFilePrefix(dbModel.Slug)
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
func (s *ProductService) DeletePhoto(id int, currentUserId int) (*model.Product, error) {
	_, err := s.getAndCheckPermissions(id, currentUserId)
	if err != nil {
		return nil, err
	}
	if err := s.repo.UpdatePhoto(id, nil); err != nil {
		return nil, err
	}
	return s.repo.GetById(id)
}

func (s *ProductService) checkPermissions(dbModel *model.Product, currentUserId int) error {
	if dbModel.CreatedById != currentUserId {
		return &exceptions.UserPermissionError{Msg: "Forbidden"}
	}
	return nil
}

func (s *ProductService) getAndCheckPermissions(id int, currentUserId int) (*model.Product, error) {
	dbModel, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	if err := s.checkPermissions(dbModel, currentUserId); err != nil {
		return nil, err
	}
	return dbModel, nil
}
func (s *ProductService) getFilePrefix(slug string) string {
	filePrefix := fmt.Sprintf("/%s/%s", productFileCategory, slug)
	return filePrefix
}
