package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWin := myApp.NewWindow("Button")

	btn1 := widget.NewButton("text button", func() {
		fmt.Println("text button clicked")
	})

	btn2 := widget.NewButtonWithIcon("icon", theme.HomeIcon(), func() {
		fmt.Println("icon button clicked")
	})

	container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), btn1, btn2)
	myWin.SetContent(container)
	myWin.Resize(fyne.NewSize(150, 50))
	myWin.ShowAndRun()
}
