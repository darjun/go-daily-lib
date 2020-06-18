package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("TabContainer")

	nameLabel := widget.NewLabel("Name: dajun")
	sexLabel := widget.NewLabel("Sex: male")
	ageLabel := widget.NewLabel("Age: 18")
	addressLabel := widget.NewLabel("Province: shanghai")
	addressLabel.Hide()
	profile := widget.NewVBox(nameLabel, sexLabel, ageLabel, addressLabel)

	musicRadio := widget.NewRadio([]string{"on", "off"}, func(string) {})
	showAddressCheck := widget.NewCheck("show address?", func(value bool) {
		if !value {
			addressLabel.Hide()
		} else {
			addressLabel.Show()
		}
	})
	memberTypeSelect := widget.NewSelect([]string{"junior", "senior", "admin"}, func(string) {})

	setting := widget.NewForm(
		&widget.FormItem{"music", musicRadio},
		&widget.FormItem{"check", showAddressCheck},
		&widget.FormItem{"member type", memberTypeSelect},
	)

	tabs := widget.NewTabContainer(
		widget.NewTabItem("Profile", profile),
		widget.NewTabItem("Setting", setting),
	)

	myWindow.SetContent(tabs)
	myWindow.Resize(fyne.NewSize(200, 200))
	myWindow.ShowAndRun()
}
