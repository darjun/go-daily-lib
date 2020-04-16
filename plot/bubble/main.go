package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
	rand.Seed(0)
	n := 10
	bubbleData := randomTriples(n)

	minZ, maxZ := math.Inf(1), math.Inf(-1)
	for _, xyz := range bubbleData {
		if xyz.Z > maxZ {
			maxZ = xyz.Z
		}
		if xyz.Z < minZ {
			minZ = xyz.Z
		}
	}

	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.Title.Text = "Bubbles"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	bs, err := plotter.NewScatter(bubbleData)
	if err != nil {
		log.Fatal(err)
	}
	bs.GlyphStyleFunc = func(i int) draw.GlyphStyle {
		c := color.RGBA{R: 196, B: 128, A: 255}
		var minRadius, maxRadius = vg.Points(1), vg.Points(20)
		rng := maxRadius - minRadius
		_, _, z := bubbleData.XYZ(i)
		d := (z - minZ) / (maxZ - minZ)
		r := vg.Length(d)*rng + minRadius
		return draw.GlyphStyle{Color: c, Radius: r, Shape: draw.CircleGlyph{}}
	}
	p.Add(bs)

	if err = p.Save(4*vg.Inch, 4*vg.Inch, "bubble.png"); err != nil {
		log.Fatal(err)
	}
}

func randomTriples(n int) plotter.XYZs {
	data := make(plotter.XYZs, n)
	for i := range data {
		if i == 0 {
			data[i].X = rand.Float64()
		} else {
			data[i].X = data[i-1].X + 2*rand.Float64()
		}
		data[i].Y = data[i].X + 10*rand.Float64()
		data[i].Z = data[i].X
	}

	return data
}
