package main

import (
	"log"
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	rand.Seed(int64(0))

	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	p.Title.Text = "Get Started"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err = plotutil.AddLinePoints(p,
		"First", randomPoints(15),
		"Second", randomPoints(15),
		"Third", randomPoints(15))
	if err != nil {
		log.Fatal(err)
	}

	if err = p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		log.Fatal(err)
	}
}

func randomPoints(n int) plotter.XYs {
	points := make(plotter.XYs, n)
	for i := range points {
		if i == 0 {
			points[i].X = rand.Float64()
		} else {
			points[i].X = points[i-1].X + rand.Float64()
		}
		points[i].Y = points[i].X + 10 * rand.Float64()
	}

	return points
}