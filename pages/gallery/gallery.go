package gallery

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/aerogo/aero"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/components"
)

// Get ...
func Get(ctx *aero.Context) string {
	files, err := ioutil.ReadDir("images/gallery/")

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Couldn't load gallery", err)
	}

	images := []string{}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if strings.HasPrefix(file.Name(), ".") {
			continue
		}

		images = append(images, "/images/gallery/"+file.Name())
	}

	return ctx.HTML(components.Gallery(images))
}
