package paperwallet

import "image/color"

type Style struct {
	Color    color.Color
	Size     float64
	Rotation float64
	X        float64
	Y        float64
}

type Config struct {
	Address   *Style
	AddressQR *Style
	Key       *Style
	KeyQR     *Style
	FontName  string
}
