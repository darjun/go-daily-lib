package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow("Hello")
	w.SetContent(widget.NewLabel("Hello Fyne!"))
	w.Resize(fyne.NewSize(200, 200))

	w.ShowAndRun()
}
