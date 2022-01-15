package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
)

// IpRequest ip request of any user that visited this site
type IpRequest struct {
	Id          int
	Ip          string
	Os          string
	Country     string
	City        string
	Geolocation string
}

func init() {
	// register model
	orm.RegisterModel(new(IpRequest))
}
