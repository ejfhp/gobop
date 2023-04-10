package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

var mainWindow fyne.Window

func main() {
	a := app.New()
	a.Settings().SetTheme(&bopTheme{})
	mainWindow = a.NewWindow("Hello Person")
	paperWalletPanel := newPaperWalletPanel()
	paperPanel, err := paperWalletPanel.makePanel()
	if err != nil {
		fmt.Printf("cannot build default image panel: %v\n", err)
		os.Exit(1)
	}
	mainContent := container.New(layout.NewFormLayout(), makeToolsPanel(), paperPanel)
	mainWindow.SetContent(mainContent)
	// w.SetContent(container.NewHBox(makeImagePanel()))
	// w.SetContent(makeImagePanel())
	mainWindow.ShowAndRun()
}
