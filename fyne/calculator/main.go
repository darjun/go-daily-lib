package main

import (
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/Knetic/govaluate"
)

func input(display *widget.Entry, value string) func() {
	return func() {
		display.Text += value
		display.Refresh()
	}
}

func equals(display *widget.Entry) func() {
	return func() {
		lines := strings.Split(display.Text, "\n")
		if len(lines) == 0 {
			return
		}

		line := lines[len(lines)-1]
		line = strings.Trim(line, "+÷×")
		exprLine := strings.Replace(line, "÷", "/", -1)
		exprLine = strings.Replace(exprLine, "×", "*", -1)
		expr, _ := govaluate.NewEvaluableExpression(exprLine)
		result, _ := expr.Evaluate(nil)
		line += "=\n"
		line += fmt.Sprint(result)
		display.Text = line
		display.Refresh()
	}
}

func percent(display *widget.Entry) func() {
	return func() {}
}

func clear(display *widget.Entry) func() {
	return func() {
		display.Text = ""
		display.Refresh()
	}
}

func sign(display *widget.Entry) func() {
	return func() {
		lines := strings.Split(display.Text, "\n")
		if len(lines) == 0 {
			return
		}

		line := lines[len(lines)-1]
		value, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return
		}
		lines[len(lines)-1] = strconv.FormatInt(-value, 10)
		display.Text = strings.Join(lines, "\n")
	}
}

func main() {
	myApp := app.New()
	myWin := myApp.NewWindow("Calculator")

	display := widget.NewEntry()
	display.Text = ""
	// display.SetReadOnly(true)
	display.MultiLine = true

	clearBtn := widget.NewButton("AC", clear(display))
	signBtn := widget.NewButton("+/-", sign(display))
	percentBtn := widget.NewButton("%", percent(display))
	divideBtn := widget.NewButton("÷", input(display, "÷"))
	clearContainer := fyne.NewContainerWithLayout(
		layout.NewGridLayout(4),
		clearBtn,
		signBtn,
		percentBtn,
		divideBtn,
	)

	digits := []string{
		"7", "8", "9", "×",
		"4", "5", "6", "-",
		"1", "2", "3", "+",
	}
	var digitBtns []fyne.CanvasObject
	for _, val := range digits {
		digitBtns = append(digitBtns, widget.NewButton(val, input(display, val)))
	}
	digitContainer := fyne.NewContainerWithLayout(
		layout.NewGridLayout(4),
		digitBtns...)

	zeroBtn := widget.NewButton("0", input(display, "0"))
	dotBtn := widget.NewButton(".", input(display, "."))
	equalBtn := widget.NewButton("=", equals(display))
	zeroContainer := fyne.NewContainerWithLayout(
		layout.NewGridLayout(2),
		zeroBtn,
		fyne.NewContainerWithLayout(
			layout.NewGridLayout(2),
			dotBtn,
			equalBtn,
		),
	)

	copyright := widget.NewLabel("2020 © power by fyne")
	copyright.Alignment = fyne.TextAlignCenter

	container := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		display,
		clearContainer,
		digitContainer,
		zeroContainer,
		copyright,
	)
	myWin.SetContent(container)
	myWin.Resize(fyne.NewSize(360, 300))
	myWin.ShowAndRun()
}
