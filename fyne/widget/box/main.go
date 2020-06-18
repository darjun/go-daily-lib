package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWin := myApp.NewWindow("Box")

	content := widget.NewVBox(
		widget.NewLabel("The top row of VBox"),
		widget.NewHBox(
			widget.NewLabel("Label 1"),
			widget.NewLabel("Label 2"),
		),
	)
	content.Append(widget.NewButton("Append", func() {
		content.Append(widget.NewLabel("Appended"))
	}))
	content.Append(widget.NewButton("Prepend", func() {
		content.Prepend(widget.NewLabel("Prepended"))
	}))

	myWin.SetContent(content)
	myWin.Resize(fyne.NewSize(150, 150))
	myWin.ShowAndRun()
}
