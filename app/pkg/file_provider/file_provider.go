package fileprovider

import "mime/multipart"

type FileProvider interface {
	PutObject(filePath string, file multipart.File, fileHeader *multipart.FileHeader) (string, error)
	CheckFile(filePath string) (bool, error)
	RemoveObject(filePath string) error
}
