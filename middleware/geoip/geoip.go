package geoip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aJuvan/pam-notify/config"
)

type MiddlewareGeoIPData struct {
	Continent string  `json:"continent_name"`
	Country   string  `json:"country_name"`
	City      string  `json:"city"`
	ISP       string  `json:"isp"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Hostname  string  `json:"hostname"`
}

type apiData struct {
	Success   bool     `json:"success"`
	Continent *string  `json:"continent_name"`
	Country   *string  `json:"country_name"`
	City      *string  `json:"city"`
	ISP       *string  `json:"isp"`
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Hostname  *string  `json:"hostname"`
}

func Run(userData *config.UserData) (*MiddlewareGeoIPData, error) {
	resp, err := http.Get("https://json.geoiplookup.io/" + userData.Rhost)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data apiData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	if !data.Success {
		return nil, fmt.Errorf("API error")
	}

	return &MiddlewareGeoIPData{
		Continent: *data.Continent,
		Country:   *data.Country,
		City:      *data.City,
		ISP:       *data.ISP,
		Latitude:  *data.Latitude,
		Longitude: *data.Longitude,
		Hostname:  *data.Hostname,
	}, nil
}
