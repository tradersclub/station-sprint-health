package img_manipulator

import (
	"image/jpeg"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

type nfntImpl struct{}

func newNfnt() IResizer {
	return nfntImpl{}
}

func (nfntImpl) Resize(height, width int, name string) error {
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
