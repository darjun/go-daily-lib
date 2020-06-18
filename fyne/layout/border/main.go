package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Border Layout")

	left := canvas.NewText("left", color.White)
	right := canvas.NewText("right", color.White)
	top := canvas.NewText("top", color.White)
	bottom := canvas.NewText("bottom", color.White)
	content := widget.NewLabel(`Lorem ipsum dolor, 
	sit amet consectetur adipisicing elit.
	Quidem consectetur ipsam nesciunt,
	quasi sint expedita minus aut,
	porro iusto magnam ducimus voluptates cum vitae.
	Vero adipisci earum iure consequatur quidem.`)

	container := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(top, bottom, left, right),
		top, bottom, left, right, content,
	)
	myWindow.SetContent(container)
	myWindow.ShowAndRun()
}
