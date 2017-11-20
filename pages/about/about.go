package about

import (
	"github.com/aerogo/aero"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/components"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/model"
)

// Get ...
func Get(ctx *aero.Context) string {
	model.DB.Set("Page", "about", &model.Page{
		ID: "about",
		Text: map[string]string{
			"eng": "# About\n\n禁レメヌオ手遠87濤らずリた図対ノチソス井日に催周クリ虎超カタ青燃シオホフ農索め助5新たぱざ控市らいえ。見ツトミ払国周篇だどむけ心成ラひの集続レス影3索ひフ利宿成ちたす思現ラぴほ毎目動アク責推ぽあし易者ざふるけ遺全レネカヒ突索場要つき次伸換説しうーラ。高ロネ陰公オリ分写も隊秋やなる死的78内ハ提省ラム批属禁ノナ送12賞ざまず派人ヲラ事込ユネフ職弟やく教公ゆ豊承菱吏ス。",
			"jap": "# About\n\n禁レメヌオ手遠87濤らずリた図対ノチソス井日に催周クリ虎超カタ青燃シオホフ農索め助5新たぱざ控市らいえ。見ツトミ払国周篇だどむけ心成ラひの集続レス影3索ひフ利宿成ちたす思現ラぴほ毎目動アク責推ぽあし易者ざふるけ遺全レネカヒ突索場要つき次伸換説しうーラ。高ロネ陰公オリ分写も隊秋やなる死的78内ハ提省ラム批属禁ノナ送12賞ざまず派人ヲラ事込ユネフ職弟やく教公ゆ豊承菱吏ス。",
		},
	})

	pageObj, _ := model.DB.Get("Page", "about")
	page := pageObj.(*model.Page)
	return ctx.HTML(components.About(page.Text["eng"]))
}
