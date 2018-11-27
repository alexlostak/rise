package forecast

import (
	"fmt"
	"strconv"
	"time"

	"github.com/rise/darksky"
)

type Forecast struct {
	Summary    string
	TempHigh   string
	TempLow    string
	Weeekday   time.Weekday
	Month      time.Month
	DayOfMonth int
	Time       string
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
	forecast.Month = t.Month()
	forecast.DayOfMonth = t.Day()
	fmt.Println(t.Weekday().String())
	return forecast
}

func (f *Forecast) GetReturnString() string {
	return ("Today is " + f.Weeekday.String() + ", " + f.Month.String() + " " + f.formatDayOfMonth() + ". The weather today will be " + f.Summary) + " with a high of " + f.TempHigh + " and a low of " + f.TempLow + "."
}

func (f *Forecast) formatDayOfMonth() string {
	remainder := f.DayOfMonth % 10
	remainderString := ""
	switch remainder {
	case 1:
		remainderString = "st"
	case 2:
		remainderString = "nd"
	case 3:
		remainderString = "rd"
	default:
		remainderString = "th"
	}
	return strconv.Itoa(f.DayOfMonth) + remainderString
}
