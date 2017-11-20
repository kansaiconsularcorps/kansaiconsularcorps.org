package fullpage

import (
	"github.com/aerogo/aero"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/components"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/model"
)

// Render renders the layout.
func Render(ctx *aero.Context, content string) string {
	pageChannel := model.DB.All("Page")
	pages := []*model.Page{}

	for pageObj := range pageChannel {
		pages = append(pages, pageObj.(*model.Page))
	}

	return components.Layout(ctx.App, pages, content)
}
