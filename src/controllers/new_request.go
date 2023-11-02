package controllers

import (
	"log"
	"main/src/models"
	"main/src/services"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

// NewRequestController struct that handles requests to register requests to save
type NewRequestController struct {
	web.Controller
}

// Get method of new request endpoint
func (c *NewRequestController) Get() {
	ipInfo, err := services.GetIPInfo(c.Ctx.Input.IP())
	if err != nil {
		log.Println("Could not get info about ip:", err)
	}
	o := orm.NewOrm()
	request := models.RequestInfo{
		IP:             c.Ctx.Input.IP(),
		Os:             services.GetOsTypeFromUserAgent(c.Ctx.Input.Header("User-Agent")),
		Country:        ipInfo.Country,
		Region:         ipInfo.Region,
		City:           ipInfo.City,
		Latitude:       ipInfo.Latitude,
		Longitude:      ipInfo.Longitude,
		UserAgent:      c.Ctx.Input.Header("User-Agent"),
		AdditionalInfo: "",
		Timestamp:      time.Now(),
	}
	id, err := o.Insert(&request)
	if err != nil {
		log.Println("Could not save into table info from address. ", err)
	}
	log.Println("new request from "+string(c.Ctx.Input.IP())+" saved to db with id ", id)
	c.Ctx.ResponseWriter.WriteHeader(204)
	c.Ctx.WriteString("")
}
