package minio

import (
	"fmt"
	fileservice "food/internal/file_service"
	fileconverter "food/pkg/file_converter"
	fileprocessor "food/pkg/file_processor"
	fileprovider "food/pkg/file_provider"
	"food/pkg/utils"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/gosimple/slug"
)

type FileService struct {
	fileProvider  fileprovider.FileProvider
	fileConverter *fileconverter.ImageConverter
	fileProcessor *fileprocessor.ImageProcessor
}

func NewFileServcie(fileProvider fileprovider.FileProvider) fileservice.FileService {
	return &FileService{fileProvider: fileProvider, fileConverter: &fileconverter.ImageConverter{}, fileProcessor: fileprocessor.NewImageProcessor(fileProvider)}
}

func (s *FileService) UploadFile(
	category string, file multipart.File, fileHeader *multipart.FileHeader,
) (string, error) {
	// Convert file to JPEG
	file, fileHeader, err := s.fileConverter.ConvertToJpg(file, fileHeader)
	if err != nil {
		return "", err
	}
	// Get slugified fileName with extension
	objectName := s.slugifyFileName(fileHeader.Filename)
	// Generate filePath
	filePath := fmt.Sprintf("%s/%s/%s", category, time.Now().Format("2006/01/02"), objectName)
	// Update fileName if it already exists
	filePath, err = s.checkFile(filePath)
	if err != nil {
		return "", err
	}
	// Upload file to bucket
	newFilePath, err := s.fileProvider.PutObject(filePath, file, fileHeader)
	if err != nil {
		return "", err
	}
	// Process image
	if utils.ContainsString([]string{".webp", ".png", ".jpg", ".jpeg"}, filepath.Ext(fileHeader.Filename)) {
		s.fileProcessor.ProcessFile(filePath, file, fileHeader)
	}
	return newFilePath, nil
}

func (s *FileService) DeleteFile(filePath string) error {
	return s.fileProvider.RemoveObject(filePath)
}

func (s *FileService) slugifyFileName(fileName string) string {
	fileName, fileExtension := s.spliteFileName(fileName)
	return fmt.Sprintf("%s.%s", slug.Make(fileName), fileExtension)
}

func (s *FileService) spliteFileName(fileName string) (string, string) {
	splitedFilePath := strings.Split(fileName, ".")
	// Get file extension
	fileExtension := splitedFilePath[len(splitedFilePath)-1]
	// Get file name
	fileName = strings.Join(splitedFilePath[0:len(splitedFilePath)-1], "-")
	return fileName, fileExtension
}

func (s *FileService) checkFile(
	filePath string,
) (string, error) {
	count := 0
	fileName, fileExtension := s.spliteFileName(filePath)
	for i := 0; i < 1000; i++ {
		postfix := ""
		if count > 0 {
			postfix = fmt.Sprintf("-%d", count)
		}
		tempName := fmt.Sprintf("%s%s.%s", fileName, postfix, fileExtension)
		exists, err := s.fileProvider.CheckFile(tempName)
		if err != nil {
			return "", nil
		}
		if !exists {
			return tempName, nil
		}
		count++
	}
	return "", nil
}
