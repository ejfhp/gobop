package graphic_test

import (
	"image/png"
	"os"
	"testing"

	"github.com/ejfhp/gobop/lib/graphic"
)

func TestQrcode_QRCode(t *testing.T) {
	img, err := graphic.QRCode("L48cWSssxbFnRuuJCVes9NEYP1W987kfpSgWG2RKSaZtcs6iCHpT")
	if err != nil {
		t.Fatalf("cannot generate qrcode image from text: %v", err)
	}
	if img.Bounds().Dx() < 10 {
		t.Fatalf("image too small: %d", img.Bounds().Dx())

	}
	tempOut, err := os.Create("/tmp/testqrcode.png")
	if err != nil {
		t.Fatalf("cannot create temp image file: %v", err)
	}
	err = png.Encode(tempOut, img)
	if err != nil {
		t.Fatalf("cannot create image file: %v", err)
	}
}
