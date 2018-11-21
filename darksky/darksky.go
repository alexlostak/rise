package darksky

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/rise/keys"
)

// DarkSkyForecast - Struct used to represent a forecast returned by DarkSky API
type DarkSkyForecast struct {
	Currently Currently `json:"currently"`
	Daily     Daily     `json:"daily"`
	Hourly    Hourly    `json:"hourly"`
	Timezone  string    `json:"timezone"`
}

// Currently - Forecast for current weather
type Currently struct {
	Temperature float64 `json:"temperature"`
}

// Hourly - Hourly forecast for weather
type Hourly struct {
	Summary string `json:"summary"`
}

// Daily - Daily forecast for the weather
type Daily struct {
	Data []DailyData `json:"data"`
}

// DailyData - Data for each individual day inside of a Data struct
type DailyData struct {
	TemperatureHigh float64 `json:"temperatureHigh"`
	TemperatureLow  float64 `json:"temperatureLow"`
}

// GetDarkSkyForecast -
func GetDarkSkyForecast(latitude string, longitude string) DarkSkyForecast {
	resp, err := http.Get("https://api.darksky.net/forecast/" + keys.DarkSkyAPI + "/" + latitude + "," + longitude)
	if err != nil {
		panic("fucked up")
	}
	respRaw, _ := ioutil.ReadAll(resp.Body)
	forecast := DarkSkyForecast{}
	json.Unmarshal(respRaw, &forecast)
	fmt.Println(forecast.Timezone)
	return forecast
}

// GetSummary - Gets the hourly weather summary
func (d *DarkSkyForecast) GetSummary() string {
	return d.Hourly.Summary
}

// GetTodaysHigh - Gets the temperature high from the first daily data and returns it as a string
func (d *DarkSkyForecast) GetTodaysHigh() string {
	return strconv.FormatFloat(d.Daily.Data[0].TemperatureHigh, 'f', 2, 64)
}

// GetTodaysLow - Gets the temperature low from the first daily data and returns it as a string
func (d *DarkSkyForecast) GetTodaysLow() string {
	return strconv.FormatFloat(d.Daily.Data[0].TemperatureLow, 'f', 2, 64)
}
