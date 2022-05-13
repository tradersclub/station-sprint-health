package main

import (
	"cha1/img_manipulator"
	"fmt"
)

func main() {
	imgManipulator := img_manipulator.New()

	err := imgManipulator.Resize(100, 100, "img.png")
	if err != nil {
		fmt.Println(err)
	}
}
