package main

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	std := plotter.Values{35510, 1960, 99}
	easyjson := plotter.Values{8499, 160, 4}
	jsoniter := plotter.Values{5623, 160, 3}

	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	p.Title.Text = "jsoniter vs easyjson vs std"
	p.Y.Label.Text = ""

	w := vg.Points(20)
	stdBar, err := plotter.NewBarChart(std, w)
	if err != nil {
		log.Fatal(err)
	}
	stdBar.LineStyle.Width = vg.Length(0)
	stdBar.Color = plotutil.Color(0)
	stdBar.Offset = -w

	easyjsonBar, err := plotter.NewBarChart(easyjson, w)
	if err != nil {
		log.Fatal(err)
	}
	easyjsonBar.LineStyle.Width = vg.Length(0)
	easyjsonBar.Color = plotutil.Color(1)

	jsoniterBar, err := plotter.NewBarChart(jsoniter, w)
	if err != nil {
		log.Fatal(err)
	}
	jsoniterBar.LineStyle.Width = vg.Length(0)
	jsoniterBar.Color = plotutil.Color(2)
	jsoniterBar.Offset = w

	p.Add(stdBar, easyjsonBar, jsoniterBar)
	p.Legend.Add("std", stdBar)
	p.Legend.Add("easyjson", easyjsonBar)
	p.Legend.Add("jsoniter", jsoniterBar)
	p.Legend.Top = true
	p.NominalX("ns/op", "allocation bytes", "allocation times")

	if err = p.Save(5*vg.Inch, 5*vg.Inch, "barchart.png"); err != nil {
		log.Fatal(err)
	}
}
