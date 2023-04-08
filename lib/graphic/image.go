package graphic

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"path/filepath"

	"github.com/fogleman/gg"
	"github.com/goki/freetype/truetype"
	"golang.org/x/image/font"
)

const (
	toRadians = 3.1415 / 180
)

func AddText(canvasImage image.Image, text string, textFontName string, textColor color.Color, textSize float64, rotationdeg, x, y float64) (image.Image, error) {
	c := gg.NewContextForImage(canvasImage)
	c.SetColor(textColor)

	if ff, err := LoadFontFace(textFontName, textSize); err == nil {
		c.SetFontFace(ff)
	} else {
		return nil, fmt.Errorf("cannot load font file '%s': %w", textFontName, err)
	}
	xd := x*math.Cos(rotationdeg*toRadians) - y*math.Sin(rotationdeg*toRadians)
	yd := x*math.Sin(rotationdeg*toRadians) + y*math.Cos(rotationdeg*toRadians)
	c.Rotate(gg.Radians(rotationdeg))
	c.DrawString(text, xd, yd)
	return c.Image(), nil
}

func AddImage(canvasImage image.Image, srcImage image.Image, scale, rotationdeg, x, y float64) image.Image {
	c := gg.NewContextForImage(canvasImage)
	xd := x*math.Cos(rotationdeg*toRadians) - y*math.Sin(rotationdeg*toRadians)
	yd := x*math.Sin(rotationdeg*toRadians) + y*math.Cos(rotationdeg*toRadians)
	c.Rotate(gg.Radians(rotationdeg))
	c.Scale(scale, scale)
	c.DrawImage(srcImage, int(xd), int(yd))
	// c.DrawImageAnchored(srcImage, int(xd), int(yd), 0.5, 0.5)
	return c.Image()
}

func LoadFontFace(path string, points float64) (font.Face, error) {
	fontBytes, err := fonts.ReadFile(filepath.Join("fonts", path))
	if err != nil {
		return nil, err
	}
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}
	face := truetype.NewFace(f, &truetype.Options{
		Size: points,
		// Hinting: font.HintingFull,
	})
	return face, nil
}
