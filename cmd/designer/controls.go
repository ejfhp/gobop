package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeToolsPanel() fyne.CanvasObject {
	verticalPanel := container.NewVBox(makeControlsPanel())
	return verticalPanel
}

func makeControlsPanel() fyne.CanvasObject {
	addressItem := makeControlItem("Address")
	addressQRItem := makeControlItem("QR Address")
	keyItem := makeControlItem("Key")
	keyQRItem := makeControlItem("QR Key")

	positionTabs := container.NewVBox(addressItem, addressQRItem, keyItem, keyQRItem)
	return positionTabs
}

func makeControlItem(name string) fyne.CanvasObject {
	itemSize := &widget.FormItem{
		Text:     "Size",
		Widget:   widget.NewEntry(),
		HintText: "Set size of element",
	}
	itemXPos := &widget.FormItem{
		Text:     "Position X",
		Widget:   widget.NewSlider(0, 1000),
		HintText: "Position X",
	}
	itemYPos := &widget.FormItem{
		Text:     "Position Y",
		Widget:   widget.NewSlider(0, 1000),
		HintText: "Position Y",
	}

	form := widget.NewForm(itemSize, itemXPos, itemYPos)
	itemCard := widget.NewCard(name, "", form)

	return itemCard
}
