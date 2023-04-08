package graphic

import (
	"bytes"
	"fmt"
	"image"
	"image/png"

	qrcode "github.com/skip2/go-qrcode"
)

func QRCode(text string) (image.Image, error) {
	imgbytes, err := qrcode.Encode(text, qrcode.Medium, 256)
	if err != nil {
		return nil, fmt.Errorf("cannot encode text '%s' to QRcode: %w", text, err)
	}
	img, err := png.Decode(bytes.NewReader(imgbytes))
	if err != nil {
		return nil, fmt.Errorf("cannot create PNG from encoded QRcode of text '%s': %w", text, err)
	}
	return img, nil
}
