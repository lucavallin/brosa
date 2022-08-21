package weather

import "github.com/lucavallin/mau/pkg/geo"

// Forecast represents the forecast for a location.
type Forecast struct {
	StartTime string     `json:"startTime"`
	EndTime   string     `json:"endTime"`
	Intervals []Interval `json:"intervals"`
}

// Interval represents the weather for a single time period.
type Interval struct {
	StartTime   string  `json:"startTime"`
	CloudCover  float64 `json:"cloudCover"`
	Humidity    float64 `json:"humidity"`
	Temperature float64 `json:"temperature"`
	Visibility  float64 `json:"visibility"`
	DewPoint    float64 `json:"dewPoint"`
}

// Provider is the interface for the weather provider.
type Provider interface {
	GetForecast(coordinates *geo.Coordinates, endTime string) (*Forecast, error)
}
