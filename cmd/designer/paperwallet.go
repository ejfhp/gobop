package main

import (
	"fmt"
	"image"
	"image/png"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/ejfhp/gobop/lib/bitcoin"
	"github.com/ejfhp/gobop/lib/graphic"
)

type PanelWalletPanel struct {
	parentWindow   fyne.Window
	background     image.Image
	image          image.Image
	canvas         *canvas.Image
	imageContainer *fyne.Container
	bitcoinKey     string
	bitcoinAddress string
}

func newPaperWalletPanel(parentWindow fyne.Window) *PanelWalletPanel {
	ppw := PanelWalletPanel{
		parentWindow: parentWindow,
		image:        getLogo(),
	}
	cvas := canvas.NewImageFromImage(ppw.image)
	cvas.FillMode = canvas.ImageFillContain
	cvas.SetMinSize(fyne.Size{Width: 100, Height: 100})
	ppw.canvas = cvas
	return &ppw
}

func (pp *PanelWalletPanel) setBitcoinKey(key string) error {
	pp.bitcoinKey = key
	add, err := bitcoin.AddressOf(key)
	if err != nil {
		return fmt.Errorf("cannot generate address from the given key")
	}
	pp.bitcoinAddress = add
	return nil
}

func (pp *PanelWalletPanel) setBackground(image image.Image) {
	pp.image = image
	pp.background = image
	pp.canvas.Image = pp.image
	pp.canvas.Resize(fyne.Size{Width: float32(image.Bounds().Dx()), Height: float32(image.Bounds().Dy())})
	pp.imageContainer.Refresh()
}

func (pp *PanelWalletPanel) addImage(image image.Image, scale, x, y, rotationDeg float64) {
	pp.image = graphic.AddImage(pp.background, image, scale, rotationDeg, x, y)
	pp.canvas.Image = pp.image
	pp.imageContainer.Refresh()
}

func (pp *PanelWalletPanel) makePanel() (fyne.CanvasObject, error) {
	pp.imageContainer = container.NewMax(pp.canvas)

	butRefreshGraphic := widget.NewButton("refresh", func() {
		qrK, err := graphic.QRCode(pp.bitcoinKey)
		if err != nil {
			dialog.NewError(fmt.Errorf("cannot create key qrcode: %w"), pp.parentWindow)
		}
		qrA, err := graphic.QRCode(pp.bitcoinAddress)
		if err != nil {
			dialog.NewError(fmt.Errorf("cannot create address qrcode: %w"), pp.parentWindow)
		}
		pp.addImage(qrK, 1, 200, 200, 0)
		pp.addImage(qrA, 1, 600, 600, 0)

	})
	butOpenGraphic := widget.NewButton("open graphic", func() {
		fmt.Printf("open graphic\n")
		dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
			fmt.Printf("file selected\n")
			if err != nil {
				fmt.Printf("error selecting file\n")
				pp.setBackground(getLogo())
				return
			}
			im, err := png.Decode(uc)
			if err != nil {
				fmt.Printf("error decoding file\n")
				pp.setBackground(getLogo())
				return
			}
			fmt.Printf("setting graphic\n")
			pp.setBackground(im)
		}, pp.parentWindow).Show()

	})
	hbox := container.NewHBox(butRefreshGraphic, butOpenGraphic)
	return container.NewVBox(hbox, pp.imageContainer), nil
}
