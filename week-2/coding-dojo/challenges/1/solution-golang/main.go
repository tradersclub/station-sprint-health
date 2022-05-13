package main

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
	"github.com/sunshineplan/imgconv"
)

type IResizer interface {
	Resize(height, width int, name string) error
}

type resizerNfnt struct{}

func newNfntResize() IResizer {
	return resizerNfnt{}
}

func (resizerNfnt) Resize(height, width int, name string) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}

	// decode jpeg into image.Image
	img, err := png.Decode(file)
	if err != nil {
		return err
	}
	file.Close()

	m := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	out, err := os.Create("resized-" + name)
	if err != nil {
		return err
	}
	defer out.Close()

	return jpeg.Encode(out, m, nil)
}

type resizerImgConv struct{}

func newImgConv() IResizer {
	return resizerImgConv{}
}

func (resizerImgConv) Resize(height, width int, name string) error {
	src, err := imgconv.Open(name)
	if err != nil {
		return err
	}
	conv := imgconv.ResizeOption{Width: int(width), Height: int(height)}

	mark := imgconv.Resize(src, conv)

	out, err := os.Create("resized-" + name)
	if err != nil {
		return err
	}
	defer out.Close()

	err = imgconv.Write(out, mark, imgconv.FormatOption{Format: imgconv.PNG})
	if err != nil {
		log.Fatalf("failed to write image: %v", err)
	}

	return nil
}

type resizerMock struct{}

func newMock() IResizer {
	return resizerMock{}
}

func (resizerMock) Resize(height, width int, name string) error {
	return nil
}

func newResizer() IResizer {
	return newNfntResize()
}

func main() {
	resizer := newResizer()

	err := resizer.Resize(100, 100, "img.png")
	if err != nil {
		fmt.Println(err)
	}
}
