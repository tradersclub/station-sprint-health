package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

func nfntAdapter(width, height uint, img image.Image) image.Image {
	return resize.Resize(width, height, img, resize.Lanczos3)
}

func main() {

	file, err := os.Open("376500-middle.png")
	if err != nil {
		fmt.Println("erro pegar imagem")
	}

	img, err := png.Decode(file)
	if err != nil {
		fmt.Println("erro decode imagem")
	}

	resizedImage := nfntAdapter(100, 50, img)

	out, err := os.Create("gosymbo.png")
	if err != nil {
		fmt.Println("erro decode imagem")
	}

	defer out.Close()

	png.Encode(out, resizedImage)

}
