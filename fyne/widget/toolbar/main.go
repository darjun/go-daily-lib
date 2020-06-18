package main

import (
	"fmt"
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Toolbar")

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			fmt.Println("New document")
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {
			fmt.Println("Cut")
		}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {
			fmt.Println("Copy")
		}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {
			fmt.Println("Paste")
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)

	content := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(toolbar, nil, nil, nil),
		toolbar, widget.NewLabel(`Lorem ipsum dolor, 
		sit amet consectetur adipisicing elit.
		Quidem consectetur ipsam nesciunt,
		quasi sint expedita minus aut,
		porro iusto magnam ducimus voluptates cum vitae.
		Vero adipisci earum iure consequatur quidem.`),
	)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
