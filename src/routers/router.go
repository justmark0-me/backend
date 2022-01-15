package routers

import (
	"github.com/beego/beego/v2/server/web"
	"main/src/controllers"
)

func init() {
	web.Router("/", &controllers.MainController{})
}
