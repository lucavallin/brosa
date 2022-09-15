package weather

import (
	"time"

	"github.com/lucavallin/mau/pkg/geo"
)

// ForecastRequest is a struct that represents an internal forecast request
type ForecastRequest struct {
	Location  *geo.Coordinates
	StartTime time.Time
	EndTime   time.Time
}

// Forecast represents the forecast for a location.
type Forecast struct {
	StartTime time.Time  `json:"startTime"`
	EndTime   time.Time  `json:"endTime"`
	Intervals []Interval `json:"intervals"`
}

// Interval represents the weather for a single time period.
type Interval struct {
	StartTime                time.Time `json:"startTime"`
	CloudCover               float64   `json:"cloudCover"`
	Humidity                 float64   `json:"humidity"`
	Temperature              float64   `json:"temperature"`
	Visibility               float64   `json:"visibility"`
	DewPoint                 float64   `json:"dewPoint"`
	PrecipitationProbability float64   `json:"precipitationProbability"`
}

// Forecaster is the interface for the weather provider.
type Forecaster interface {
	Get(*ForecastRequest) (*Forecast, error)
}

// Forecast returns the forecast for the given request using the given forecaster
func GetForecast(forecaster Forecaster, request *ForecastRequest) (*Forecast, error) {
	return forecaster.Get(request)
}
