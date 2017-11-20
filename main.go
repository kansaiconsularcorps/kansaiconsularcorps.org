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
	l.Page("/", home.English)
	l.Page("/japanese", home.Japanese)
	l.Page("/about", page.Get)
	l.Page("/gallery", gallery.Get)
	l.Page("/contact", contact.Get)

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

	return app
}
