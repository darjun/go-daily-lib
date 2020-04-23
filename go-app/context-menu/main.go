package main

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v6/pkg/app"
)

type ContextMenuUI struct {
	app.Compo
	name string
}

func (c *ContextMenuUI) Render() app.UI {
	return app.Div().Body(
		app.Text("Hello, World"),
	).OnContextMenu(c.OnContextMenu)
}

func (*ContextMenuUI) OnContextMenu(src app.Value, event app.Event) {
	event.PreventDefault()

	app.NewContextMenu(
		app.MenuItem().
			Label("item 1").
			OnClick(func(src app.Value, e app.Event) {
				fmt.Println("item 1 clicked")
			}),
		app.MenuItem().Separator(),
		app.MenuItem().
			Label("item 2").
			OnClick(func(src app.Value, e app.Event) {
				fmt.Println("item 2 clicked")
			}),
	)
}

func main() {
	app.Route("/", &ContextMenuUI{})
	app.Run()
}
