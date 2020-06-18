package main

import (
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("ProgressBar")

	bar1 := widget.NewProgressBar()
	bar1.Min = 0
	bar1.Max = 100
	bar2 := widget.NewProgressBarInfinite()

	go func() {
		for i := 0; i <= 100; i++ {
			time.Sleep(time.Millisecond * 500)
			bar1.SetValue(float64(i))
		}
	}()

	content := widget.NewVBox(bar1, bar2)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(150, 150))
	myWindow.ShowAndRun()
}
