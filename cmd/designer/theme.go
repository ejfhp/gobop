package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var _ fyne.Theme = (*bopTheme)(nil)
var mapSize = map[string]float32{

	"text":           13,
	"headingText":    15,
	"helperText":     11,
	"padding":        6,
	"lineSpacing":    4,
	"iconInline":     20,
	"innerPadding":   8,
	"inputBorder":    1,
	"scrollBar":      16,
	"scrollBarSmall": 3,
}

type bopTheme struct{}

func (m bopTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return color.White
		}
		return color.Black
	}

	return theme.DefaultTheme().Color(name, variant)
}
func (m bopTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	// if name == theme.IconNameHome {
	// 	fyne.NewStaticResource("myHome", homeBytes)
	// }

	return theme.DefaultTheme().Icon(name)
}

func (m bopTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m bopTheme) Size(name fyne.ThemeSizeName) float32 {
	size, ok := mapSize[string(name)]
	if ok {
		if theme.DefaultTheme().Size(name) != size {
			fmt.Printf(" %s size:%f default:%f\n", name, size, theme.DefaultTheme().Size(name))
		}
		return size
	}
	fmt.Printf("Undefined: %s\n", name)
	return theme.DefaultTheme().Size(name)
}
