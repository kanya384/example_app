package helpers

import (
	"bytes"
	"image"
	"image/jpeg"
)

func ConvertImageToBytes(image image.Image) (res []byte, err error) {
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, image, nil)
	if err != nil {
		return
	}
	res = buf.Bytes()
	return
}
