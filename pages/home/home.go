package home

import (
	"github.com/aerogo/aero"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/components"
)

// Get ...
func Get(ctx *aero.Context) string {
	return ctx.HTML(components.Home())
}
