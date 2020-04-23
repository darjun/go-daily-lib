package main

import "github.com/maxence-charriere/go-app/v6/pkg/app"

type Greeting struct {
	app.Compo
	name string
}

func (g *Greeting) Render() app.UI {
	return app.Div().Body(
		app.Main().Body(
			app.H1().Body(
				app.Text("Hello, "),
				app.If(g.name != "",
					app.Text(g.name),
				).Else(
					app.Text("World"),
				),
			),
		),
		app.Input().
			Value(g.name).
			Placeholder("What is your name?").
			AutoFocus(true).
			OnChange(g.OnInputChange),
	)
}

func (g *Greeting) OnInputChange(src app.Value, e app.Event) {
	g.name = src.Get("value").String()
	g.Update()
}

func main() {
	app.Route("/", &Greeting{})
	app.Route("/app", &app.Compo{})
	app.Run()
}
