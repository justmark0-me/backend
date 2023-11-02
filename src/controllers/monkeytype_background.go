package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

// MonkeytypeBackgroundController is a struct that returns background for monkeytype
type MonkeytypeBackgroundController struct {
	web.Controller
}

// Get method for a background
func (c *MonkeytypeBackgroundController) Get() {
	c.Ctx.Output.ContentType("image/jpeg")
	picturePath := "/app/files/monkeytype_background.jpg"
	c.Ctx.Output.Download(picturePath)
}
