package controllers

import (
	"os"

	beego "github.com/beego/beego/v2/server/web"
)

type StatController struct {
	beego.Controller
}

func (c *StatController) Get() {
	name := os.Getenv("NAME")
	if len(name) == 0 {
		name = "User"
	}
	c.Data["Name"] = name
	c.TplName = "stat.tpl"
}
