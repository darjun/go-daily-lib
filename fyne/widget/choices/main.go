package main

import (
	"fmt"

	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWin := myApp.NewWindow("Choices")

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("input name")

	passEntry := widget.NewPasswordEntry()
	passEntry.SetPlaceHolder("input password")

	repeatPassEntry := widget.NewPasswordEntry()
	repeatPassEntry.SetPlaceHolder("repeat password")

	nameBox := widget.NewHBox(widget.NewLabel("Name"), layout.NewSpacer(), nameEntry)
	passwordBox := widget.NewHBox(widget.NewLabel("Password"), layout.NewSpacer(), passEntry)
	repeatPasswordBox := widget.NewHBox(widget.NewLabel("Repeat Password"), layout.NewSpacer(), repeatPassEntry)

	sexRadio := widget.NewRadio([]string{"male", "female", "unknown"}, func(value string) {
		fmt.Println("sex:", value)
	})
	sexBox := widget.NewHBox(widget.NewLabel("Sex"), sexRadio)

	football := widget.NewCheck("football", func(value bool) {
		fmt.Println("football:", value)
	})
	basketball := widget.NewCheck("basketball", func(value bool) {
		fmt.Println("basketball:", value)
	})
	pingpong := widget.NewCheck("pingpong", func(value bool) {
		fmt.Println("pingpong:", value)
	})
	hobbyBox := widget.NewHBox(widget.NewLabel("Hobby"), football, basketball, pingpong)

	provinceSelect := widget.NewSelect([]string{"anhui", "zhejiang", "shanghai"}, func(value string) {
		fmt.Println("province:", value)
	})
	provinceBox := widget.NewHBox(widget.NewLabel("Province"), layout.NewSpacer(), provinceSelect)

	registerBtn := widget.NewButton("Register", func() {
		fmt.Println("name:", nameEntry.Text, "password:", passEntry.Text, "register")
	})

	content := widget.NewVBox(nameBox, passwordBox, repeatPasswordBox, sexBox, hobbyBox, provinceBox, registerBtn)
	myWin.SetContent(content)
	myWin.ShowAndRun()
}
