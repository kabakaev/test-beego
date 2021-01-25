package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type AboutController struct {
	beego.Controller
}

func (c *AboutController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "about.tpl"
}
