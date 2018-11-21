package darksky

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rise/keys"
)

func GetSummaries(longitude string, latitude string) {
	resp, err := http.Get("https://api.darksky.net/forecast/" + keys.DarkSkyAPI + "/" + longitude + "," + latitude)
	if err != nil {
		panic("fucked up")
	}
	respRaw, _ := ioutil.ReadAll(resp.Body)
	forecast := Forecast{}
	json.Unmarshal(respRaw, &forecast)
	fmt.Println("Current summary: " + forecast.Currently.Summary)
	fmt.Println("Current summary: " + forecast.Daily.Summary)
	fmt.Println("Current summary: " + forecast.Hourly.Summary)
}
