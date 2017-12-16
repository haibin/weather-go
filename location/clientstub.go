package location

import (
	"testing"

	wapi "github.com/haibin/weather-go/openweather/api"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type WeatherClientStub struct {
}

func (c *WeatherClientStub) FindForLocation(city string, state string) (wapi.Conditions, error) {
	if city == "Austin" && state == "Texas" {
		return wapi.Conditions{
			Main: wapi.Main{
				Temperature: 75,
			},
			Weather: []wapi.Weather{
				wapi.Weather{
					Description: "sunny",
				},
			},
		}, nil
	} else {
		return wapi.Conditions{}, nil
	}
}
