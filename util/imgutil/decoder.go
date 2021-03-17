package imgutil

import (
	"fmt"
	"golang.org/x/image/bmp"
	"golang.org/x/image/webp"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

func GetDecoder(mimeType string) func(io.Reader) (image.Image, error) {
	switch mimeType {
	case "image/jpeg":
		return jpeg.Decode
	case "image/png":
		return png.Decode
	case "image/bmp":
		return bmp.Decode
	case "image/gif":
		return gif.Decode
	case "image/webp":
		return webp.Decode
	}

	return func(_ io.Reader) (image.Image, error) {
		return nil, fmt.Errorf("%s, mime type not supported", mimeType)
	}
}