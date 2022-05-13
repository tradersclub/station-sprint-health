package img_manipulator

type IResizer interface {
	Resize(height, width int, name string) error
}

func New() IResizer {
	return newNfnt()
}
