package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Max Layout")

	image := canvas.NewImageFromResource(theme.FyneLogo())
	text := canvas.NewText("Fyne Logo", color.Black)
	text.Alignment = fyne.TextAlignCenter

	container := fyne.NewContainerWithLayout(
		layout.NewMaxLayout(),
		image, text,
	)
	myWindow.SetContent(container)
	myWindow.Resize(fyne.NewSize(200, 200))
	myWindow.ShowAndRun()
}
