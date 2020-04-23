package main

import "github.com/maxence-charriere/go-app/v6/pkg/app"

type ShowSelect struct {
	app.Compo
	option string
  }
  
  func (s *ShowSelect) Render() app.UI {
	return app.Div().Body(
	  app.Main().Body(
		app.H1().Body(
		  app.If(s.option == "",
			app.Text("Please select!"),
		  ).Else(
			app.Text("You've selected "+s.option),
		  ),
		),
	  ),
	  app.Select().Body(
		app.Option().Body(
		  app.Text("apple"),
		),
		app.Option().Body(
		  app.Text("orange"),
		),
		app.Option().Body(
		  app.Text("banana"),
		),
	  ).
		OnChange(s.OnSelectChange),
	)
  }
  
  func (s *ShowSelect) OnSelectChange(src app.Value, e app.Event) {
	s.option = src.Get("value").String()
	s.Update()
  }
  
  func main() {
	app.Route("/", &ShowSelect{})
	app.Run()
  }