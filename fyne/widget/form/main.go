package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Form")

	nameEntry := widget.NewEntry()
	passEntry := widget.NewPasswordEntry()

	form := widget.NewForm(
		&widget.FormItem{"Name", nameEntry},
		&widget.FormItem{"Pass", passEntry},
	)
	form.OnSubmit = func() {
		fmt.Println("name:", nameEntry.Text, "pass:", passEntry.Text, "login in")
	}
	form.OnCancel = func() {
		fmt.Println("login canceled")
	}

	myWindow.SetContent(form)
	myWindow.Resize(fyne.NewSize(150, 150))
	myWindow.ShowAndRun()
}
