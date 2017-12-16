package client

import (
	"testing"

	"github.com/haibin/weather-go/openweather/api"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type IntTestSuite struct {
}

var _ = Suite(&IntTestSuite{})

// Find should return weather for a city/state
func (s *IntTestSuite) TestFind(c *C) {
	// given
	client := WeatherClient{}

	// when
	cond, err := client.FindForLocation("London", "uk")

	// then
	c.Assert(err, Equals, nil)

	c.Assert(cond.Main.Temperature > 0, Equals, true)
	c.Assert(cond.Weather[0].Description, Not(Equals), "")
}

// Client should return empty "Conditions" when a state isn't found
func (s *IntTestSuite) TestFindNotFound(c *C) {
	// given
	client := WeatherClient{}

	// when
	cond, err := client.FindForLocation("Foo", "Bar")

	// then
	// c.Assert(err, Equals, nil)
	c.Assert(err, Not(Equals), nil)

	c.Assert(cond, DeepEquals, api.Conditions{})
}
