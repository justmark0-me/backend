package controllers

import "github.com/beego/beego/v2/server/web"

type MainController struct {
	web.Controller
}

func (this *MainController) Get() {
	this.TplName = "index.tpl" // version 1.6 use this.TplName = "index.tpl"
	this.Ctx.WriteString("hello world")
}
