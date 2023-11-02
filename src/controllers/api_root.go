package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

// IndexController is a struct that handles all requests for index page
type IndexController struct {
	web.Controller
}

// Get method of index endpoint
func (c *IndexController) Get() {
	c.Ctx.WriteString("There is no documentation of API :'(")
}
