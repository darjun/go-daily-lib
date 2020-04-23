package main

import "github.com/maxence-charriere/go-app/v6/pkg/app"

type RangeUI struct {
	app.Compo
	name string
}

func (*RangeUI) Render() app.UI {
	langs := []string{"Go", "JavaScript", "Python", "C"}
	return app.Ul().Body(
		app.Range(langs).Slice(func(i int) app.UI {
			return app.Li().Body(
				app.Text(langs[i]),
			)
		}),
	)
}

func main() {
	app.Route("/", &RangeUI{})
	app.Run()
}
