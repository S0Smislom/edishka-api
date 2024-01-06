package mockfileservice

import (
	fileservice "food/internal/file_service"
	"mime/multipart"
)

type MockFileService struct {
}

func NewMockFileServicez() fileservice.FileService {
	return &MockFileService{}
}

func (s *MockFileService) UploadFile(category string, file multipart.File, fileHeader *multipart.FileHeader,
) (string, error) {
	return "/mock/" + fileHeader.Filename, nil
}

func (s *MockFileService) DeleteFile(filePath string) error {
	return nil
}
