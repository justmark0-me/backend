package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq" // we need to import postgres driver to use it
)

// RequestInfo ip request of any user that visited this site
type RequestInfo struct {
	ID             int    // Id is not spelled as ID because it would ugly see in postgres
	IP             string // Ip is not spelled as IP because it would ugly see in postgres
	Os             string
	Country        string
	Region         string
	City           string
	Latitude       float64
	Longitude      float64
	UserAgent      string
	AdditionalInfo string
	Timestamp      time.Time
}

func init() {
	// register model
	orm.RegisterModel(new(RequestInfo))
}
