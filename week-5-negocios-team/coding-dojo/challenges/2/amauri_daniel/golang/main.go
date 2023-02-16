package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/sunshineplan/imgconv"
)

type Image interface {
	Resize(inputPath string, outputPath string, width int, height int) error
}

func Resizer(imageResizerAdapter Image, inputPath string, outputPath string, width int, height int) error {
	return imageResizerAdapter.Resize(inputPath, outputPath, width, height)
}

type ImgconvAdapter struct{}

func (i ImgconvAdapter) Resize(inputPath string, outputPath string, width int, height int) error {
	src, err := imgconv.Open(inputPath)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	image := imgconv.Resize(src, &imgconv.ResizeOption{Width: 10})

	return imgconv.Save(outputPath, image, &imgconv.FormatOption{Format: imgconv.JPEG})
}

func main() {
	imagePath := "imagens/cachorro.jpg"
	imagePathDestination := fmt.Sprintf("imagens/resized/new%s.jpg", strconv.FormatInt(time.Now().UnixMilli(), 10))

	imgconvAdapter := ImgconvAdapter{}
	imgconvAdapter.Resize(imagePath, imagePathDestination, 50, 50)
}
