package home

import (
	"strconv"

	"github.com/aerogo/aero"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/components"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/model"
)

// English ...
func English(ctx *aero.Context) string {
	stories := []*model.Story{}

	for i := 1; i <= 6; i++ {
		stories = append(stories, &model.Story{
			Title:    "Story " + strconv.Itoa(i),
			Language: "en",
			Summary:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam tortor est, iaculis nec interdum eu, gravida et tortor. Morbi eleifend est sit amet dolor tristique vestibulum. Quisque leo urna, placerat vel dapibus in, elementum vitae lacus. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. In placerat lorem sed nisi posuere varius.",
			Image:    "http://d1dqs20vqfxmh2.cloudfront.net/wp-content/uploads/2017/06/DSC_0333_low-res-306x192.jpg",
		})
	}

	return ctx.HTML(components.Home(stories))
}

// Japanese ...
func Japanese(ctx *aero.Context) string {
	stories := []*model.Story{}

	for i := 1; i <= 6; i++ {
		stories = append(stories, &model.Story{
			Title:    "ストーリー " + strconv.Itoa(i),
			Language: "ja",
			Summary:  "禁レメヌオ手遠87濤らずリた図対ノチソス井日に催周クリ虎超カタ青燃シオホフ農索め助5新たぱざ控市らいえ。見ツトミ払国周篇だどむけ心成ラひの集続レス影3索ひフ利宿成ちたす思現ラぴほ毎目動アク責推ぽあし易者ざふるけ遺全レネカヒ突索場要つき次伸換説しうーラ。高ロネ陰公オリ分写も隊秋やなる死的78内ハ提省ラム批属禁ノナ送12賞ざまず派人ヲラ事込ユネフ職弟やく教公ゆ豊承菱吏ス。",
			Image:    "http://d1dqs20vqfxmh2.cloudfront.net/wp-content/uploads/2017/06/DSC_0333_low-res-306x192.jpg",
		})
	}

	return ctx.HTML(components.Home(stories))
}
