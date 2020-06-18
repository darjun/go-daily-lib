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

	menu := widget.NewMenu(
		"File",
		widget.NewMenuItem("New", func() {
			fmt.Println("New")
		}),
		widget.NewMenuItem("Open", func() {
			fmt.Println("Open")
		}),
		widget.NewMenuItemSeparator(),
		widget.NewMenuItem("Save", func() {
			fmt.Println("Save")
		}),
	)

	logo := canvas.NewImageFromResource(theme.FyneLogo())

	myWindow.SetContent(logo)
	myWindow.ShowAndRun()
}
