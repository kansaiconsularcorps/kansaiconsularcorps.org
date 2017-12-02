package main

import (
	"github.com/aerogo/aero"
	"github.com/aerogo/layout"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/components/css"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/components/js"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/layout"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/model"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/pages/contact"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/pages/gallery"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/pages/home"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/pages/login"
	"github.com/kansaiconsularcorps/kansaiconsularcorps.org/pages/page"
)

func main() {
	app := aero.New()
	configure(app).Run()
}

func configure(app *aero.Application) *aero.Application {
	// Layout
	l := layout.New(app)
	l.Render = fullpage.Render

	// Routing
	l.Page("/", home.Get)
	l.Page("/about", page.Get)
	l.Page("/gallery", gallery.Get)
	l.Page("/contact", contact.Get)
	l.Page("/login", login.Get)

	// Styles
	app.SetStyle(css.Bundle())

	// Scripts
	scripts := js.Bundle()

	app.Get("/scripts", func(ctx *aero.Context) string {
		return ctx.JavaScript(scripts)
	})

	app.Get("/images/*file", func(ctx *aero.Context) string {
		return ctx.File("images" + ctx.Get("file"))
	})

	model.DB.Prefetch()

	app.OnShutdown(func() {
		model.DB.Close()
	})

	// Initial data
	model.DB.Set("Page", "about", &model.Page{
		ID: "about",
		Title: map[string]string{
			"en": "About",
			"ja": "KCCについて",
		},
		Text: map[string]string{
			"en": "# About\n\n禁レメヌオ手遠87濤らずリた図対ノチソス井日に催周クリ虎超カタ青燃シオホフ農索め助5新たぱざ控市らいえ。見ツトミ払国周篇だどむけ心成ラひの集続レス影3索ひフ利宿成ちたす思現ラぴほ毎目動アク責推ぽあし易者ざふるけ遺全レネカヒ突索場要つき次伸換説しうーラ。高ロネ陰公オリ分写も隊秋やなる死的78内ハ提省ラム批属禁ノナ送12賞ざまず派人ヲラ事込ユネフ職弟やく教公ゆ豊承菱吏ス。",
			"ja": "# About\n\n禁レメヌオ手遠87濤らずリた図対ノチソス井日に催周クリ虎超カタ青燃シオホフ農索め助5新たぱざ控市らいえ。見ツトミ払国周篇だどむけ心成ラひの集続レス影3索ひフ利宿成ちたす思現ラぴほ毎目動アク責推ぽあし易者ざふるけ遺全レネカヒ突索場要つき次伸換説しうーラ。高ロネ陰公オリ分写も隊秋やなる死的78内ハ提省ラム批属禁ノナ送12賞ざまず派人ヲラ事込ユネフ職弟やく教公ゆ豊承菱吏ス。",
		},
	})

	// for i := 1; i <= 6; i++ {
	// 	story := &model.Story{
	// 		ID:       strconv.Itoa(i),
	// 		Title:    "Story " + strconv.Itoa(i),
	// 		Language: "en",
	// 		Summary:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam tortor est, iaculis nec interdum eu, gravida et tortor. Morbi eleifend est sit amet dolor tristique vestibulum. Quisque leo urna, placerat vel dapibus in, elementum vitae lacus. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. In placerat lorem sed nisi posuere varius.",
	// 		Image:    "http://d1dqs20vqfxmh2.cloudfront.net/wp-content/uploads/2017/06/DSC_0333_low-res-306x192.jpg",
	// 		Created:  model.DateTimeUTC(),
	// 	}
	// 	time.Sleep(time.Second)

	// 	model.DB.Set("Story", story.ID, story)
	// }

	// for i := 1; i <= 6; i++ {
	// 	story := &model.Story{
	// 		ID:       strconv.Itoa(i + 6),
	// 		Title:    "ストーリー " + strconv.Itoa(i),
	// 		Language: "ja",
	// 		Summary:  "禁レメヌオ手遠87濤らずリた図対ノチソス井日に催周クリ虎超カタ青燃シオホフ農索め助5新たぱざ控市らいえ。見ツトミ払国周篇だどむけ心成ラひの集続レス影3索ひフ利宿成ちたす思現ラぴほ毎目動アク責推ぽあし易者ざふるけ遺全レネカヒ突索場要つき次伸換説しうーラ。高ロネ陰公オリ分写も隊秋やなる死的78内ハ提省ラム批属禁ノナ送12賞ざまず派人ヲラ事込ユネフ職弟やく教公ゆ豊承菱吏ス。",
	// 		Image:    "http://d1dqs20vqfxmh2.cloudfront.net/wp-content/uploads/2017/06/DSC_0333_low-res-306x192.jpg",
	// 		Created:  model.DateTimeUTC(),
	// 	}
	// 	time.Sleep(time.Second)

	// 	model.DB.Set("Story", story.ID, story)
	// }

	return app
}
