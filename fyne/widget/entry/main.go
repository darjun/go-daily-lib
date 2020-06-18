package main

import (
	"fmt"

	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWin := myApp.NewWindow("Entry")

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("input name")
	nameEntry.OnChanged = func(content string) {
		fmt.Println("name:", nameEntry.Text, "entered")
	}

	passEntry := widget.NewPasswordEntry()
	passEntry.SetPlaceHolder("input password")

	nameBox := widget.NewHBox(widget.NewLabel("Name"), layout.NewSpacer(), nameEntry)
	passwordBox := widget.NewHBox(widget.NewLabel("Password"), layout.NewSpacer(), passEntry)

	loginBtn := widget.NewButton("Login", func() {
		fmt.Println("name:", nameEntry.Text, "password:", passEntry.Text, "login in")
	})

	multiEntry := widget.NewEntry()
	multiEntry.SetPlaceHolder("please enter\nyour description")
	multiEntry.MultiLine = true

	content := widget.NewVBox(nameBox, passwordBox, loginBtn, multiEntry)
	myWin.SetContent(content)
	myWin.ShowAndRun()
}
