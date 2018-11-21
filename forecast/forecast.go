package forecast

import (
	"time"

	"github.com/rise/darksky"
)

type Forecast struct {
	Summary  string
	TempHigh string
	TempLow  string
	Weeekday time.Weekday
	Date     string
	Time     string
}

// GetForecast -
func GetForecast(latitude string, longitude string) Forecast {
	// Declare Forecast to create
	forecast := Forecast{}
	// Get Weather
	dsf := darksky.GetDarkSkyForecast(latitude, longitude)
	forecast.Summary = dsf.GetSummary()
	forecast.TempHigh = dsf.GetTodaysHigh()
	forecast.TempLow = dsf.GetTodaysLow()
	// Get Time and Date info
	t := time.Now()

	forecast.Weeekday = t.Weekday()

	return forecast
}
