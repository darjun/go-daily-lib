package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("Canvas")

	gradient := canvas.NewRadialGradient(color.White, color.Transparent)
	w.SetContent(gradient)
	w.Resize(fyne.NewSize(200, 200))
	w.ShowAndRun()
}
