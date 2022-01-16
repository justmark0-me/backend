package routers

import (
	"main/src/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.IndexController{})
	web.Router("/v1/new_request/", &controllers.NewRequestController{})
}
