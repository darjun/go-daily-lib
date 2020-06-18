package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	text := canvas.NewText("Come on!", color.White)
	text.Alignment = fyne.TextAlignCenter
	w.SetContent(text)
	w.ShowAndRun()
}
