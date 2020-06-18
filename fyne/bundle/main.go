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
	myWindow := myApp.NewWindow("Bundle Resource")

	img1 := canvas.NewImageFromResource(resourceImage1Png)
	img1.FillMode = canvas.ImageFillOriginal
	img2 := canvas.NewImageFromResource(resourceImage2Jpg)
	img2.FillMode = canvas.ImageFillOriginal
	img3 := canvas.NewImageFromResource(theme.FyneLogo())
	img3.FillMode = canvas.ImageFillOriginal

	container := fyne.NewContainerWithLayout(
		layout.NewGridLayout(1),
		img1, img2, img3,
	)
	myWindow.SetContent(container)
	myWindow.ShowAndRun()
}
