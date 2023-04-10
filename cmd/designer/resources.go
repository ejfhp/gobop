package main

import (
	"embed"
	"fmt"
	"image"
	"image/png"
)

//go:embed resources/*
var resources embed.FS

func getLogo() image.Image {
	logoImg, err := resources.Open("resources/logo.png")
	if err != nil {
		fmt.Printf("cannot open logo: %w", err)
		return image.Black
	}
	img, err := png.Decode(logoImg)
	if err != nil {
		fmt.Printf("cannot decode logo: %w", err)
		return image.Black
	}
	return img
}
