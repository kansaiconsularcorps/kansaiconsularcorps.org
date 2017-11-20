package home

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/components"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/model"
)

// Get ...
func Get(ctx *aero.Context) string {
	language := "en"
	stories := model.FilterStories(func(story *model.Story) bool {
		return story.Language == language
	})

	sort.Slice(stories, func(i, j int) bool {
		return stories[i].Created > stories[j].Created
	})

	return ctx.HTML(components.Home(stories))
}
