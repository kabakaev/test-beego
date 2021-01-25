package controllers

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/astaxie/beego/cache"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

var (
	responses cache.Cache
)

type MainController struct {
	beego.Controller
}

func init() {
	responses, _ = cache.NewCache("memory", `{"interval":0}`) // See https://pkg.go.dev/gitea.com/tango/cache#Options
}

func (c *MainController) Get() {
	name := os.Getenv("NAME")
	if len(name) == 0 {
		name = "User"
	}
	c.Data["Name"] = name

	h1size := c.GetString("h1size")
	switch c.GetString("action") {
	case "good":
		if responses.IsExist(h1size) {
			if err := responses.Incr(h1size); err != nil {
				logs.Warn("Cannot increment %s: %w", h1size, err)
			} else {
				logs.Info("Incremented " + h1size)
			}
		} else {
			if err := responses.Put(h1size, 1, 0); err != nil {
				logs.Warn("Cannot initialize %s: %w", h1size, err)
			}
		}
	case "bad":
		if responses.IsExist(h1size) {
			if err := responses.Decr(h1size); err != nil {
				logs.Warn("Cannot decrement %s: %w", h1size, err)
			} else {
				logs.Info("Decremented " + h1size)
			}
		} else {
			if err := responses.Put(h1size, -1, 0); err != nil {
				logs.Warn("Cannot initialize %s: %w", h1size, err)
			}
		}
	}

	headerFontMin := 32 / 16
	headerFontMax := 128 / 16
	// Random integer in steps of 16.
	randomInt := 16 * (rand.Intn(headerFontMax-headerFontMin) + headerFontMin)
	logs.Info("Random between %i: %s", headerFontMax-headerFontMin, randomInt)
	c.Data["HeaderFontSize"] = fmt.Sprint(randomInt)
	c.TplName = "index.tpl"
}
