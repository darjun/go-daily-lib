package main

import (
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"fyne.io/fyne/theme"
)

type tappableIcon struct {
	widget.Icon
}

func newTappableIcon(res fyne.Resource) *tappableIcon {
	icon := &tappableIcon{}
	icon.ExtendBaseWidget(icon)
	icon.SetResource(res)

	return icon
}

func (t *tappableIcon) Tapped(e *fyne.PointEvent) {
	log.Println("I have been left tapped at", e)
}

func (t *tappableIcon) TappedSecondary(e *fyne.PointEvent) {
	log.Println("I have been right tapped at", e)
}

func (t *tappableIcon) DoubleTapped(e *fyne.PointEvent) {
	log.Println("I have been double tapped at", e)
}

func main() {
	a := app.New()
	w := a.NewWindow("Tappable")
	w.SetContent(newTappableIcon(theme.FyneLogo()))
	w.Resize(fyne.NewSize(200, 200))
	w.ShowAndRun()
}