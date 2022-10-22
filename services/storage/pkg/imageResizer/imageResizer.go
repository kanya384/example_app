package imageResizer

import (
	"bytes"
	"fmt"
	"image"

	"github.com/disintegration/imaging"
)

type imageResizer struct{}

func NewImageResizer() *imageResizer {
	return &imageResizer{}
}

func (i *imageResizer) ResizeImage(imgBytes []byte, width int, height int) (img image.Image, err error) {
	image, format, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		return
	}
	fmt.Println(format)
	if height != 0 {
		return imaging.Fill(image, width, height, imaging.Center, imaging.Lanczos), nil
	} else {
		return imaging.Resize(image, width, 0, imaging.Lanczos), nil
	}
}
