package main

import (
	"github.com/aerogo/aero"
	"github.com/aerogo/layout"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/components/css"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/components/js"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/layout"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/pages/about"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/pages/contact"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/pages/gallery"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/pages/home"
)

func main() {
	app := aero.New()
	configure(app).Run()
}

func configure(app *aero.Application) *aero.Application {
	// Layout
	l := layout.New(app)
	l.Render = fullpage.Render

	// Routing
	l.Page("/", home.Get)
	l.Page("/about", about.Get)
	l.Page("/gallery", gallery.Get)
	l.Page("/contact", contact.Get)

	// Styles
	app.SetStyle(css.Bundle())

	// Scripts
	scripts := js.Bundle()

	app.Get("/scripts", func(ctx *aero.Context) string {
		return ctx.JavaScript(scripts)
	})

	return app
}
