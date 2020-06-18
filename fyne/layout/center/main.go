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
	myWindow := myApp.NewWindow("Center Layout")

	image := canvas.NewImageFromResource(theme.FyneLogo())
	image.FillMode = canvas.ImageFillOriginal
	text := canvas.NewText("Fyne Logo", color.Black)

	container := fyne.NewContainerWithLayout(
		layout.NewCenterLayout(),
		image, text,
	)
	myWindow.SetContent(container)
	myWindow.ShowAndRun()
}
