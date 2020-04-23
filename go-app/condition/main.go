package main

import (
	"strconv"

	"github.com/maxence-charriere/go-app/v6/pkg/app"
)

type ScoreUI struct {
	app.Compo
	score int
}

func (c *ScoreUI) Render() app.UI {
	return app.Div().Body(
		app.If(c.score >= 90,
			app.H1().
				Style("color", "green").
				Body(
					app.Text("Good!"),
				),
		).ElseIf(c.score >= 60,
			app.H1().
				Style("color", "orange").
				Body(
					app.Text("Pass!"),
				),
		).Else(
			app.H1().
				Style("color", "red").
				Body(
					app.Text("fail!"),
				),
		),
		app.Input().
			Value(c.score).
			Placeholder("Input your score?").
			AutoFocus(true).
			OnChange(c.OnInputChange),
	)
}

func (c *ScoreUI) OnInputChange(src app.Value, e app.Event) {
	score, _ := strconv.ParseUint(src.Get("value").String(), 10, 32)
	c.score = int(score)
	c.Update()
}

func main() {
	app.Route("/", &ScoreUI{})
	app.Run()
}