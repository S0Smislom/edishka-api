package objectstorage

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitMinio(
	endpoint string,
	accessKey string,
	secretKey string,
	useSSL bool,
) (*minio.Client, error) {
	minioClinet, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}
	return minioClinet, nil
}
