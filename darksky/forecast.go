package darksky

// Forecast - Struct used to represent Forecast returned by DarkSky API
type Forecast struct {
	Currently Currently `json:"currently"`
	Daily     Daily     `json:"daily"`
	Hourly    Hourly    `json:"hourly"`
}

// Currently - Forecast for current weather
type Currently struct {
	Summary string `json:"summary"`
}

// Hourly - Hourly forecast for weather
type Hourly struct {
	Summary string `json:"summary"`
}

// Daily - Daily forecast for the weather
type Daily struct {
	Summary string `json:"summary"`
}
