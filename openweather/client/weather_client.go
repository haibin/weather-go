package client

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/benschw/opin-go/rest"
	"github.com/haibin/weather-go/openweather/api"
)

const UriString string = "http://api.openweathermap.org/data/2.5/weather" //Austin,Texas

type WeatherClient struct {
}

func (c *WeatherClient) FindForLocation(city string, state string) (api.Conditions, error) {
	var cond api.Conditions

	wURL, _ := url.Parse(UriString)
	q := wURL.Query()
	q.Add("q", fmt.Sprintf("%s,%s", city, state))
	q.Add("units", "imperial")
	q.Add("APPID", os.Getenv("APPID"))
	wURL.RawQuery = q.Encode()

	r, err := rest.MakeRequest("GET", wURL.String(), nil)
	if err != nil {
		return cond, err
	}
	err = rest.ProcessResponseEntity(r, &cond, http.StatusOK)
	return cond, err
}
