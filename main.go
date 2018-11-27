package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rise/forecast"
)

func main() {
	r := gin.Default()
	r.GET("/forecast", func(c *gin.Context) {
		forecast, success := handleRequest(c)
		if success {
			c.JSON(200, gin.H{

				"forecast": forecast,
			})
		} else {
			c.JSON(400, gin.H{

				"error": "latitude or longitude not provided",
			})
		}

	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func handleRequest(c *gin.Context) (string, bool) {
	latitude := c.Query("latitude")
	longitude := c.Query("longitude")
	if len(latitude) == 0 || len(longitude) == 0 {
		return "", false
	}
	f := forecast.GetForecast(latitude, longitude)
	return f.GetReturnString(), true
}
