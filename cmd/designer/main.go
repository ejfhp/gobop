package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/ejfhp/gobop/lib/bitcoin"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&bopTheme{})
	mainWindow := a.NewWindow("Hello Person")
	paperWalletPanel := newPaperWalletPanel(mainWindow)
	k, err := bitcoin.WIF()
	if err != nil {
		fmt.Printf("cannot generate bitcoin key: %v\n", err)
		os.Exit(1)
	}
	paperWalletPanel.setBitcoinKey(k)
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
