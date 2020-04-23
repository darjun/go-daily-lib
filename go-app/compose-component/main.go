package main

import "github.com/maxence-charriere/go-app/v6/pkg/app"

type Greeting struct {
	app.Compo
}

func (g *Greeting) Render() app.UI {
	return app.P().Body(
		app.Text("Hello, "),
		&Name{name: "dj"},
	)
}

type Name struct {
	app.Compo
	name string
}

func (n *Name) Render() app.UI {
	return app.Text(n.name)
}

func main() {
	app.Route("/", &Greeting{})
	app.Run()
}
