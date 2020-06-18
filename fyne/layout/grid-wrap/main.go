package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Wrap Layout")

	img1 := canvas.NewImageFromResource(theme.FyneLogo())
	img2 := canvas.NewImageFromResource(theme.FyneLogo())
	img3 := canvas.NewImageFromResource(theme.FyneLogo())
	myWindow.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewGridWrapLayout(fyne.NewSize(150, 150)),
			img1, img2, img3))
	myWindow.ShowAndRun()
}
