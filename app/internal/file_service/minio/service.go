package minio

import (
	"context"
	"fmt"
	fileservice "food/internal/file_service"
	"mime/multipart"
	"strings"
	"time"

	"github.com/gosimple/slug"
	"github.com/minio/minio-go/v7"
)

const (
	bucketName = "public"
)

type FileService struct {
	minioClient *minio.Client
}

func NewFileServcie(minioClient *minio.Client) fileservice.FileService {
	return &FileService{minioClient: minioClient}
}

func (s *FileService) UploadFile(
	category string, file multipart.File, fileHeader *multipart.FileHeader,
) (string, error) {
	ctx := context.Background()
	// Get slugified fileName with extension
	objectName := s.slugifyFileName(fileHeader.Filename)
	// Generate filePath
	filePath := fmt.Sprintf("%s/%s/%s", category, time.Now().Format("2006/01/02"), objectName)
	// Update fileName if it already exists
	filePath, err := s.checkFile(bucketName, filePath)
	if err != nil {
		return "", err
	}
	fmt.Println(fileHeader.Header)
	// Upload file to bucket
	_, err = s.minioClient.PutObject(ctx, bucketName, filePath, file, fileHeader.Size, minio.PutObjectOptions{ContentType: fileHeader.Header.Get("Content-Type")})
	if err != nil {
		return "", err
	}
	// Generate full filePath with bucketName
	// return fmt.Sprintf("/%s%s", bucketName, filePath), nil
	return filePath, nil
}

func (s *FileService) DeleteFile(filePath string) error {
	if err := s.minioClient.RemoveObject(context.Background(), bucketName, filePath, minio.RemoveObjectOptions{}); err != nil {
		return err
	}
	return nil
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
	bucketName string,
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
		_, err := s.minioClient.StatObject(context.Background(), bucketName, tempName, minio.StatObjectOptions{})
		if err != nil {
			if err.Error() == "The specified key does not exist." {
				return tempName, nil
			}
			return "", err
		}

		count++
	}
	return "", nil
}
