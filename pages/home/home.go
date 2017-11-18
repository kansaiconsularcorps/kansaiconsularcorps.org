package home

import (
	"strconv"

	"github.com/aerogo/aero"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/components"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/model"
)

// Get ...
func Get(ctx *aero.Context) string {
	stories := []*model.Story{}

	for i := 1; i <= 10; i++ {
		stories = append(stories, &model.Story{
			Title:   "Story " + strconv.Itoa(i),
			Summary: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam tortor est, iaculis nec interdum eu, gravida et tortor. Morbi eleifend est sit amet dolor tristique vestibulum. Quisque leo urna, placerat vel dapibus in, elementum vitae lacus. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. In placerat lorem sed nisi posuere varius.",
			Image:   "http://d1dqs20vqfxmh2.cloudfront.net/wp-content/uploads/2017/06/DSC_0333_low-res-306x192.jpg",
		})
	}
	return ctx.HTML(components.Home(stories))
}
