package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
)

type diagonal struct {
}

func (d *diagonal) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := 0, 0
	for _, o := range objects {
		childSize := o.MinSize()

		w += childSize.Width
		h += childSize.Height
	}

	return fyne.NewSize(w, h)
}

func (d *diagonal) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(0, 0)
	for _, o := range objects {
		size := o.MinSize()
		o.Resize(size)
		o.Move(pos)

		pos = pos.Add(fyne.NewPos(size.Width, size.Height))
	}
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Diagonal Layout")

	img1 := canvas.NewImageFromResource(theme.FyneLogo())
	img1.FillMode = canvas.ImageFillOriginal
	img2 := canvas.NewImageFromResource(theme.FyneLogo())
	img2.FillMode = canvas.ImageFillOriginal
	img3 := canvas.NewImageFromResource(theme.FyneLogo())
	img3.FillMode = canvas.ImageFillOriginal

	container := fyne.NewContainerWithLayout(
		&diagonal{},
		img1, img2, img3,
	)
	myWindow.SetContent(container)
	myWindow.ShowAndRun()
}
