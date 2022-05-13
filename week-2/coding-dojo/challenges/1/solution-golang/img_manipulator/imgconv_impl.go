package img_manipulator

import (
	"log"
	"os"

	"github.com/sunshineplan/imgconv"
)

type imgConvImpl struct{}

func newImgConv() IResizer {
	return imgConvImpl{}
}

func (imgConvImpl) Resize(height, width int, name string) error {
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
