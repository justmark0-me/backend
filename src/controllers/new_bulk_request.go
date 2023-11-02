package controllers

import (
	"log"
	"main/src/models"
	"main/src/services"
	"os"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

// NewBulkAddIPAddresses struct that handles requests for bulk add ip addresses
type NewBulkAddIPAddresses struct {
	web.Controller
}

// Post method of add ip addresses endpoint
func (c *NewBulkAddIPAddresses) Post() {
	log.Println(
		"New request to add ip addresses. Password request: ",
		c.GetString("password"),
		"Actual password: ",
		os.Getenv("ADMIN_PASSWORD")[0:2]+"...")
	password := c.GetString("password")
	if password == "" {
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.Ctx.WriteString("password is required")
		return
	}
	if password != os.Getenv("ADMIN_PASSWORD") {
		c.Ctx.ResponseWriter.WriteHeader(403)
		c.Ctx.WriteString("Wrong admin password")
		return
	}
	allIPStrings := c.GetString("all_ips")
	if allIPStrings == "" {
		c.Ctx.ResponseWriter.WriteHeader(400)
		c.Ctx.WriteString("all_ips field is required")
		return
	}
	allIps := strings.Split(allIPStrings, "\n")
	for i := 0; i < len(allIps); i++ {
		ip := allIps[i]
		ipInfo, err := services.GetIPInfo(ip)
		if err != nil {
			log.Println("Could not get info about ip:", err)
		}
		o := orm.NewOrm()
		request := models.RequestInfo{
			IP:             ip,
			Os:             "",
			Country:        ipInfo.Country,
			Region:         ipInfo.Region,
			City:           ipInfo.City,
			Latitude:       ipInfo.Latitude,
			Longitude:      ipInfo.Longitude,
			UserAgent:      "",
			AdditionalInfo: "",
			Timestamp:      time.Now(),
		}
		id, err := o.Insert(&request)
		if err != nil {
			log.Println("Could not save into table info from address. ", err)
		}
		log.Println("new request from "+string(ip)+" saved to db with id ", id)
	}
	c.Ctx.ResponseWriter.WriteHeader(204)
	c.Ctx.WriteString("")
}
