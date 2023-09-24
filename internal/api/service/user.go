package service

import (
	"fmt"
	"food/internal/api/model"
	"food/internal/api/repository"
	fileservice "food/internal/file_service"
	"mime/multipart"
)

const (
	userPhotoCategory = "user"
)

type UserService struct {
	repo        repository.User
	fileService fileservice.FileService
}

func NewUserService(repo repository.User, fileService fileservice.FileService) *UserService {
	return &UserService{
		repo:        repo,
		fileService: fileService,
	}
}

func (s *UserService) GetById(item_id int) (*model.User, error) {
	return s.repo.GetById(item_id)
}

func (s *UserService) Update(id int, data *model.UpdateUser) (*model.User, error) {
	if err := s.repo.Update(id, data); err != nil {
		return nil, err
	}
	return s.GetById(id)
}

func (s *UserService) UploadPhoto(id int, file multipart.File, fileHeader *multipart.FileHeader) (*model.User, error) {
	dbModel, err := s.GetById(id)
	if err != nil {
		return nil, err
	}
	filePrefix := s.getFilePrefix(dbModel.ID)
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

func (s *UserService) DeletePhoto(id int) (*model.User, error) {
	if err := s.repo.UpdatePhoto(id, nil); err != nil {
		return nil, err
	}
	return s.repo.GetById(id)
}

func (s *UserService) getFilePrefix(id int) string {
	filePrefix := fmt.Sprintf("/%s/%d", userPhotoCategory, id)
	return filePrefix
}
