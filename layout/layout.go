package fullpage

import (
	"github.com/aerogo/aero"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/components"
)

// Render renders the layout.
func Render(ctx *aero.Context, content string) string {
	return components.Layout(ctx.App, content)
}
