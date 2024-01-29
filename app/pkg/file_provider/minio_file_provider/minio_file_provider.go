package miniofileprovider

import (
	"context"
	"mime/multipart"
	"strings"

	"github.com/minio/minio-go/v7"
)

const bucketName = "public"

type minioFileProvider struct {
	minioClient *minio.Client
}

func NewMinioFileProvider(client *minio.Client) *minioFileProvider {
	return &minioFileProvider{
		minioClient: client,
	}
}

func (s *minioFileProvider) PutObject(filePath string, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	ctx := context.Background()
	_, err := s.minioClient.PutObject(ctx, bucketName, filePath, file, fileHeader.Size, minio.PutObjectOptions{ContentType: fileHeader.Header.Get("Content-Type")})
	if err != nil {
		return "", err
	}
	return filePath, nil
}

func (s *minioFileProvider) CheckFile(filePath string) (bool, error) {
	_, err := s.minioClient.StatObject(context.Background(), bucketName, filePath, minio.StatObjectOptions{})
	if err != nil {
		if err.Error() == "The specified key does not exist." {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *minioFileProvider) RemoveObject(filePath string) error {
	return s.minioClient.RemoveObject(context.Background(), bucketName, filePath, minio.RemoveObjectOptions{})
}

func (s *minioFileProvider) GetObject(filePath string) error {
	return nil
}

func (s *minioFileProvider) splitFilePath(filePath string) (string, string) {
	temp := strings.Split(filePath, "/")
	bucketIndex := 0
	if len(temp) > 0 && temp[0] == "" {
		bucketIndex = 1
	}
	bucketName := temp[bucketIndex]
	filePath_ := strings.Join(temp[bucketIndex+1:], "/")

	return bucketName, filePath_
}
