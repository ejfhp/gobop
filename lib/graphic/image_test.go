package graphic_test

import (
	"image/color"
	"image/png"
	"os"
	"testing"

	"github.com/ejfhp/gobop/lib/graphic"
)

func TestImage_AddText(t *testing.T) {
	imgFile, err := os.Open("testdata/white.png")
	if err != nil {
		t.Fatalf("cannot read test png")
	}
	img, err := png.Decode(imgFile)
	if err != nil {
		t.Fatalf("cannot decode test png")
	}
	txtImg, err := graphic.AddText(img, "prova", "luxisr.ttf", color.Black, 30, 90, 100, 100)
	if err != nil {
		t.Fatalf("cannot att text to test png: %v", err)
	}

	tempOut, err := os.Create("/tmp/writtenimage.png")
	if err != nil {
		t.Fatalf("cannot create temp image file: %v", err)
	}
	err = png.Encode(tempOut, txtImg)
	if err != nil {
		t.Fatalf("cannot create image file: %v", err)
	}

}

func TestImage_AddImage(t *testing.T) {
	imgFile, err := os.Open("testdata/white.png")
	if err != nil {
		t.Fatalf("cannot read test png")
	}
	qrFile, err := os.Open("testdata/qrcode.png")
	if err != nil {
		t.Fatalf("cannot read test qrcode png")
	}
	img, err := png.Decode(imgFile)
	if err != nil {
		t.Fatalf("cannot decode test png")
	}
	qrimg, err := png.Decode(qrFile)
	if err != nil {
		t.Fatalf("cannot decode test png")
	}
	imgImg := graphic.AddImage(img, qrimg, 1, 5, 100, 100)

	tempOut, err := os.Create("/tmp/imageinimage.png")
	if err != nil {
		t.Fatalf("cannot create temp image file: %v", err)
	}
	err = png.Encode(tempOut, imgImg)
	if err != nil {
		t.Fatalf("cannot create image file: %v", err)
	}

}
