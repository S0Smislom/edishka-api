package fileprocessor

import (
	"fmt"
	fileprovider "food/pkg/file_provider"
	"food/pkg/utils"
	"image"
	"image/jpeg"
	"log"
	"mime/multipart"
	"strings"

	"github.com/nfnt/resize"
)

type ImageProcessor struct {
	fileProvider fileprovider.FileProvider
}

func NewImageProcessor(fileProvider fileprovider.FileProvider) *ImageProcessor {
	return &ImageProcessor{fileProvider: fileProvider}
}

func (s *ImageProcessor) ProcessFile(filePath string, file multipart.File, fileHeader *multipart.FileHeader) error {
	go func() {
		geometry := [][]int{
			{100, 100},
			{200, 200},
			{60, 60},
		}
		img, err := jpeg.Decode(file)
		if err != nil {
			fmt.Println(err)
		}
		for _, v := range geometry {
			go func(v []int) {
				scaledImg := s.scale(img, v)
				cropedImg := s.crop(scaledImg, v)
				newFile, newFileHeader, err := utils.EncodeJpeg(cropedImg, fileHeader)
				if err != nil {
					log.Println(err)
					return
				}
				newFilePath := s.convertFilePath(filePath, v)
				s.fileProvider.PutObject(newFilePath, newFile, newFileHeader)
			}(v)

		}
	}()

	return nil
}

func (s *ImageProcessor) convertFilePath(filePath string, geometry []int) string {
	splited := strings.Split(filePath, "/")
	newFilePath := fmt.Sprintf("%s/%dx%d/%s", strings.Join(splited[0:len(splited)-1], "/"), geometry[0], geometry[1], splited[len(splited)-1])
	return newFilePath
}

func (s *ImageProcessor) scale(img image.Image, geometry []int) image.Image {
	bounds := img.Bounds()
	x, y := bounds.Dx(), bounds.Dy()
	factor := s.calculate_scale_factor(x, y, geometry)
	if factor < 1 {
		width := float64(x) * factor
		height := float64(y) * factor
		img = resize.Resize(uint(width), uint(height), img, resize.Lanczos2)
	}
	return img
}

func (s *ImageProcessor) crop(img image.Image, geometry []int) image.Image {
	bounds := img.Bounds()
	x, y := bounds.Dx(), bounds.Dy()
	width, height := geometry[0], geometry[1]

	startX := (x - width) / 2
	startY := (y - height) / 2

	cropRect := image.Rect(startX, startY, startX+geometry[0], startY+geometry[1])
	croppedImg := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(cropRect)
	return croppedImg
}

func (s *ImageProcessor) calculate_scale_factor(x_img, y_img int, geometry []int) float64 {
	factors := []float64{
		float64(geometry[0]) / float64(x_img),
		float64(geometry[1]) / float64(y_img),
	}
	min := factors[0]
	max := factors[1]
	for i := 1; i < len(factors); i++ {
		if factors[i] < min {
			min = factors[i]
		}
		if factors[i] > max {
			max = factors[i]
		}
	}
	return max
}
