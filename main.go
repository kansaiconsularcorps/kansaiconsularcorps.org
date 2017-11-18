package main

import (
	"github.com/aerogo/aero"
	"github.com/aerogo/layout"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/layout"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/pages/home"
)

func main() {
	app := aero.New()
	configure(app).Run()
}

func configure(app *aero.Application) *aero.Application {
	l := layout.New(app)
	l.Render = fullpage.Render
	l.Page("/", home.Get)

	return app
}
