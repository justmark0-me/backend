package routers

import (
	"main/src/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.IndexController{})
	web.Router("/v1/new_request/", &controllers.NewRequestController{})
	web.Router("/v1/admin/add_ip_addresses/", &controllers.NewBulkAddIPAddresses{}) // since I do not have authentication for now I made this strange url
}
