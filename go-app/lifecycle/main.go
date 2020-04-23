package main

import (
	"fmt"
	"net/url"

	"github.com/maxence-charriere/go-app/v6/pkg/app"
)

type Foo struct {
	app.Compo
}

func (*Foo) Render() app.UI {
	return app.P().Body(
		app.Text("Hello World"),
	)
}

func (*Foo) OnMount() {
	fmt.Println("component mounted")
}

func (*Foo) OnNav(u *url.URL) {
	fmt.Println("component navigated:", u)
}

func (*Foo) OnDismount() {
	fmt.Println("component dismounted")
}

func main() {
	app.Route("/", &Foo{})
	app.Run()
}
