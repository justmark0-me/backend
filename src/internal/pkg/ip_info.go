package pkg

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
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
	resp, err := http.Get("https://ipwhois.app/json/" + ip)
	if err != nil {
		return IPInfo{}, errors.Wrap(err, "get response from ipwhois.app")
	}

	var data IPInfo
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return IPInfo{}, errors.Wrap(err, "unmarshal ip info")
	}
	return data, nil
}
