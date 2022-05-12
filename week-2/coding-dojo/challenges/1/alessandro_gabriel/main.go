package main

import (
	"image"
)

type ResizeImg interface {
	Resize(width, height uint, img image.Image) image.Image
}

func (r *resizeImg) Resize(width, height uint, img image.Image) image.Image {
	return img
}

func NewResizeImg() ResizeImg {
}

func resize(res ResizeImg, width, height uint, img image.Image) {
	res.Resize(width, height, img)
}

func main() {

	var image image.Image

	resizeTest := resizeImg.Resize(10, 20, image)

}
