package paperwallet_test

import (
	"image/color"
	"image/png"
	"os"
	"testing"

	"github.com/ejfhp/gobop/lib/bitcoin"
	"github.com/ejfhp/gobop/lib/paperwallet"
)

func TestPaperwallet_Image(t *testing.T) {
	styleAdd := &paperwallet.Style{Color: color.White, Size: 50.0, Rotation: 0, X: 100, Y: 130}
	styleAddQR := &paperwallet.Style{Color: color.Black, Size: 1.0, Rotation: 0, X: 100, Y: 230}
	styleKey := &paperwallet.Style{Color: color.White, Size: 50.0, Rotation: 0, X: 100, Y: 1230}
	styleKeyQR := &paperwallet.Style{Color: color.Black, Size: 1.0, Rotation: 0, X: 100, Y: 830}
	config := paperwallet.Config{
		Address:   styleAdd,
		AddressQR: styleAddQR,
		Key:       styleKey,
		KeyQR:     styleKeyQR,
		FontName:  "luxisr.ttf",
	}

	imageFile, err := os.Open("testdata/canvas.png")
	if err != nil {
		t.Fatalf("cannot open canvas file: %v", err)
	}

	walletImage, err := png.Decode(imageFile)
	if err != nil {
		t.Fatalf("cannot decode png of canvas file: %v", err)
	}

	key, err := bitcoin.WIF()
	if err != nil {
		t.Fatalf("cannot create Bitcoin WIF: %v", err)
	}

	address, err := bitcoin.AddressOf(key)
	if err != nil {
		t.Fatalf("cannot create Address of WIF: %v", err)
	}
	wallet := paperwallet.PaperWallet{
		Config:  &config,
		Graphic: walletImage,
		WIF:     key,
		Address: address,
	}

	paperw, err := wallet.Image()
	if err != nil {
		t.Fatalf("cannot create PaperWallet image: %v", err)
	}

	paperwFile, err := os.Create("/tmp/paperwallet.png")
	if err != nil {
		t.Fatalf("cannot create PaperWallet file: %v", err)
	}

	err = png.Encode(paperwFile, paperw)
	if err != nil {
		t.Fatalf("cannot encode PaperWallet file: %v", err)
	}

}
