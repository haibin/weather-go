package location

import (
	wapi "github.com/haibin/weather-go/openweather/api"
)

type WeatherClient interface {
	FindForLocation(city string, state string) (wapi.Conditions, error)
}
