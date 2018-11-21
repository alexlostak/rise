package darksky

// Forecast - Struct used to represent Forecast returned by DarkSky API
type Forecast struct {
	Currently Currently `json:"currently"`
	Daily     []Daily   `json:"daily"`
}

// Currently - Forecast for current weather
type Currently struct {
}

// Hourly - Hourly forecast for weather
type Hourly struct {
	Summary string `json:"summary"`
}

// Daily - Daily forecast for the weather
type Daily struct {
}
