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
)

type PanelWalletPanel struct {
	image          image.Image
	canvas         *canvas.Image
	imageContainer *fyne.Container
}

func newPaperWalletPanel() *PanelWalletPanel {
	ppw := PanelWalletPanel{
		image: getLogo(),
	}
	cvas := canvas.NewImageFromImage(ppw.image)
	cvas.FillMode = canvas.ImageFillContain
	cvas.SetMinSize(fyne.Size{Width: 100, Height: 100})
	ppw.canvas = cvas
	return &ppw
}

func (pp *PanelWalletPanel) setGraphic(image image.Image) {
	pp.image = image
	pp.canvas.Image = pp.image
	pp.imageContainer.Refresh()

}

func (pp *PanelWalletPanel) makePanel() (fyne.CanvasObject, error) {
	pp.imageContainer = container.NewMax(pp.canvas)

	butOpenGraphich := widget.NewButton("open graphic", func() {
		fmt.Printf("open graphic\n")
		dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
			fmt.Printf("file selected\n")
			if err != nil {
				fmt.Printf("error selecting file\n")
				pp.setGraphic(getLogo())
				return
			}
			im, err := png.Decode(uc)
			if err != nil {
				fmt.Printf("error decoding file\n")
				pp.setGraphic(getLogo())
				return
			}
			fmt.Printf("setting graphic\n")
			pp.setGraphic(im)
		}, mainWindow).Show()

	})
	hbox := container.NewHBox(butOpenGraphich)
	return container.NewVBox(hbox, pp.imageContainer), nil
}
