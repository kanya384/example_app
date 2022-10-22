package imageResizer

import "image"

type ImageResizer interface {
	ResizeImage(imgBytes []byte, width int, height int) (img image.Image, err error)
}
