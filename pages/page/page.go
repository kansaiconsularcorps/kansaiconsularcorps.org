package page

import (
	"net/http"
	"strings"

	"github.com/aerogo/aero"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/components"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/model"
)

// Get ...
func Get(ctx *aero.Context) string {
	pageID := ctx.URI()
	pageID = strings.TrimPrefix(pageID, "/")

	pageObj, err := model.DB.Get("Page", "about")

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Page not found", err)
	}

	page := pageObj.(*model.Page)

	return ctx.HTML(components.Page(page.Text["en"]))
}
