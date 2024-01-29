package fileservice

import "mime/multipart"

const (
	MediaPath = "media"
)

type FileService interface {
	UploadFile(category string, file multipart.File, fileHeader *multipart.FileHeader) (string, error)
	DeleteFile(filePath string) error
}
