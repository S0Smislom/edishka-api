package fileconverter

import (
	"food/pkg/utils"
	"image/png"

	"golang.org/x/image/webp"

	"mime/multipart"
	"path/filepath"
)

type ImageConverter struct{}

func (s *ImageConverter) ConvertToJpg(file multipart.File, fileHeader *multipart.FileHeader) (multipart.File, *multipart.FileHeader, error) {
	fileName := fileHeader.Filename
	fileExt := filepath.Ext(fileName)
	switch {
	case fileExt == ".webp":
		return s.webpToJpeg(file, fileHeader)
	case fileExt == ".png":
		return s.pngToJpeg(file, fileHeader)
	}
	return file, fileHeader, nil
}

func (s *ImageConverter) webpToJpeg(file multipart.File, fileHeader *multipart.FileHeader) (multipart.File, *multipart.FileHeader, error) {
	img, err := webp.Decode(file)
	if err != nil {
		return nil, nil, err
	}

	return utils.EncodeJpeg(img, fileHeader)
}

func (s *ImageConverter) pngToJpeg(file multipart.File, fileHeader *multipart.FileHeader) (multipart.File, *multipart.FileHeader, error) {
	img, err := png.Decode(file)
	if err != nil {
		return nil, nil, err
	}
	return utils.EncodeJpeg(img, fileHeader)
}
