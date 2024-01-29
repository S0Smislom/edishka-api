package fileconverter

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"

	"golang.org/x/image/webp"

	"io"
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

	return s.encodeJpeg(img, fileHeader)
}

func (s *ImageConverter) pngToJpeg(file multipart.File, fileHeader *multipart.FileHeader) (multipart.File, *multipart.FileHeader, error) {
	img, err := png.Decode(file)
	if err != nil {
		return nil, nil, err
	}
	return s.encodeJpeg(img, fileHeader)
}

func (s *ImageConverter) encodeJpeg(img image.Image, fileHeader *multipart.FileHeader) (multipart.File, *multipart.FileHeader, error) {
	var byteBuffer bytes.Buffer

	if err := jpeg.Encode(&byteBuffer, img, &jpeg.Options{Quality: 90}); err != nil {
		return nil, nil, err
	}

	newFileHeader := &multipart.FileHeader{
		Filename: fileHeader.Filename[:len(fileHeader.Filename)-len(filepath.Ext(fileHeader.Filename))] + ".jpeg",
		Header: map[string][]string{
			"Content-Type": {"image/jpeg"},
		},
		Size: int64(byteBuffer.Len()),
	}
	r := io.NewSectionReader(bytes.NewReader(byteBuffer.Bytes()), 0, int64(byteBuffer.Len()))
	newFile := struct {
		*io.SectionReader
		io.Closer
	}{
		r,
		nil,
	}
	return newFile, newFileHeader, nil
}
