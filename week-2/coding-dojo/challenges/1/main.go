package main

import (
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
)

type IResizer interface {
	Resize(height, width uint, name string)
}

type resizer struct{}

func newNfntResize() IResizer {
	return resizer{}
}

func (resizer) Resize(height, width uint, name string) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	m := resize.Resize(width, height, img, resize.Lanczos3)

	out, err := os.Create("resized-" + name)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}

func main() {
	resizer := newNfntResize()

	resizer.Resize(100, 100, "img.png")
}
