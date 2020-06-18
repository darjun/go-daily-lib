package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWin := myApp.NewWindow("Label")

	l1 := widget.NewLabel("Name")
	l2 := widget.NewLabel("da\njun")

	container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), l1, l2)
	myWin.SetContent(container)
	myWin.Resize(fyne.NewSize(150, 150))
	myWin.ShowAndRun()
}
