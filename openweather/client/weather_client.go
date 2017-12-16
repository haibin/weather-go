package client

import (
	"fmt"
	"github.com/benschw/opin-go/rest"
	"github.com/haibin/weather-go/openweather/api"
	"log"
	"net/http"
)

var _ = log.Print

const UriString string = "http://api.openweathermap.org/data/2.5/weather?units=imperial&q=" //Austin,Texas

type WeatherClient struct {
}

func (c *WeatherClient) FindForLocation(city string, state string) (api.Conditions, error) {
	var cond api.Conditions

	url := fmt.Sprintf("%s%s,%s", UriString, city, state)
	r, err := rest.MakeRequest("GET", url, nil)
	if err != nil {
		return cond, err
	}
	err = rest.ProcessResponseEntity(r, &cond, http.StatusOK)
	return cond, err
}
