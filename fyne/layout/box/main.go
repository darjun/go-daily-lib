package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Box Layout")

	hcontainer1 := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		canvas.NewText("left", color.White),
		canvas.NewText("right", color.White))

	// 左对齐
	hcontainer2 := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		canvas.NewText("left", color.White),
		canvas.NewText("right", color.White))

	// 右对齐
	hcontainer3 := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		canvas.NewText("left", color.White),
		canvas.NewText("right", color.White),
		layout.NewSpacer())

	// 中间对齐
	hcontainer4 := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		canvas.NewText("left", color.White),
		canvas.NewText("right", color.White),
		layout.NewSpacer())

	// 两边对齐
	hcontainer5 := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		canvas.NewText("left", color.White),
		layout.NewSpacer(),
		canvas.NewText("right", color.White))

	myWindow.SetContent(fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		hcontainer1, hcontainer2, hcontainer3, hcontainer4, hcontainer5))
	myWindow.Resize(fyne.NewSize(200, 200))
	myWindow.ShowAndRun()
}
