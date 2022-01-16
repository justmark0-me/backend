package services

import (
	"encoding/json"
	"log"

	"github.com/beego/beego/v2/client/httplib"
)

// IPInfo information about ip address
type IPInfo struct {
	Country   string  `json:"country"`
	Region    string  `json:"region"`
	City      string  `json:"city"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// GetIPInfo returns info that were fetched from API about ip that requested something from server
func GetIPInfo(ip string) (IPInfo, error) {
	req := httplib.Get("https://ipwhois.app/json/" + ip)
	body, _ := req.Bytes()
	data := IPInfo{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Could not map structure and response of get ip info. Error:", err)
	}
	return data, nil
}
