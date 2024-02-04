package utils

import (
	"bytes"
	"image"
	"image/jpeg"
	"io"
	"mime/multipart"
	"path/filepath"
)

func EncodeJpeg(img image.Image, fileHeader *multipart.FileHeader) (multipart.File, *multipart.FileHeader, error) {
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
