package main

import (
	"image"
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	img1 := canvas.NewImageFromResource(theme.FyneLogo())
	img1.FillMode = canvas.ImageFillOriginal

	img2 := canvas.NewImageFromFile("./luffy.jpg")
	img2.FillMode = canvas.ImageFillOriginal

	image := image.NewNRGBA(image.Rectangle{image.Point{0, 0}, image.Point{100, 100}})
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			image.Set(i, j, color.NRGBA{uint8(i % 256), uint8(j % 256), 0, 255})
		}
	}
	img3 := canvas.NewImageFromImage(image)
	img3.FillMode = canvas.ImageFillOriginal

	container := fyne.NewContainerWithLayout(
		layout.NewGridWrapLayout(fyne.NewSize(150, 150)),
		img1, img2, img3)
	w.SetContent(container)
	w.ShowAndRun()
}
