package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/lucavallin/mau/pkg/geo"
)

// Notice: for internal names, we use "tio" instead of "TomorrowIO"
// tioBaseUrl is the base URL for the tomorrow.io API
const tioBaseUrl = "https://api.tomorrow.io/v4"

// TomorrowIo is a client for the tomorrow.io API
type TomorrowIo struct {
	client *http.Client
}

// tioTransport is a custom transport for the TomorrowIo client
type tioTransport struct {
	apiKey string
}

// tioForecast is a struct that represents the forecast data returned by the API
type tioForecast struct {
	Data tioData `json:"data"`
}

// tioData is a struct that represents the data returned by the API
type tioData struct {
	Timelines []tioTimeline `json:"timelines"`
}

// tioTimeline is a struct that represents the timeline data returned by the API
type tioTimeline struct {
	Timestep  string        `json:"timestep"`
	EndTime   string        `json:"endTime"`
	StartTime string        `json:"startTime"`
	Intervals []tioInterval `json:"intervals"`
}

// tioInterval is a struct that represents the interval data returned by the API
type tioInterval struct {
	StartTime string    `json:"startTime"`
	Values    tioValues `json:"values"`
}

// tioValues is a struct that represents the values returned by the API
type tioValues struct {
	CloudCover  float64 `json:"cloudCover"`
	Humidity    float64 `json:"humidity"`
	Temperature float64 `json:"temperature"`
	Visibility  float64 `json:"visibility"`
}

// NewTomorrowIo returns a new TomorrowIO client with the given API key.
func NewTomorrowIo(apiKey string) *TomorrowIo {
	return &TomorrowIo{
		client: &http.Client{
			Transport: &tioTransport{
				apiKey: apiKey,
			},
		},
	}
}

// tioTransport is a custom transport for the tio client,
// used to set common headers and provide the API key.
func (t *tioTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Accept", "application/json")

	query := req.URL.Query()
	query.Add("apikey", t.apiKey)
	query.Add("units", "metric")
	req.URL.RawQuery = query.Encode()

	return http.DefaultTransport.RoundTrip(req)
}

// GetForecast returns the forecast for the given coordinates and until the specified endTime
func (t *TomorrowIo) GetForecast(coordinates *geo.Coordinates, endTime string) (*Forecast, error) {
	req, err := http.NewRequest("GET", tioBaseUrl+"/timelines", nil)
	if err != nil {
		return nil, errors.New("tomorrow.io: failed to create request")
	}

	// this could be represented as a GetForecastRequest struct, but I'm not sure it's worth it
	query := req.URL.Query()
	query.Add("location", fmt.Sprintf("%f,%f", coordinates.Latitude, coordinates.Longitude))
	query.Add("fields", "temperature,humidity,visibility,cloudCover")
	query.Add("timesteps", "1h")
	query.Add("startTime", "now")
	query.Add("endTime", endTime)
	req.URL.RawQuery = query.Encode()

	res, err := t.client.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		return nil, errors.New("tomorrow.io: failed to get response")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("tomorrow.io: failed to read response body")
	}

	forecast, err := t.unmarshalForecast(body)
	if err != nil {
		return nil, err
	}

	return forecast, nil
}

// unmarshalForecast is a wrapper around the json.Unmarshal function that
// unmarshals the forecast data from the response body into a tioForecast struct,
// but returns a Forecast struct instead. This is because the API returns a lot of fields we don't need
// and we want a Forecast struct with only the bare minimum fields
func (t *TomorrowIo) unmarshalForecast(forecastBody []byte) (*Forecast, error) {
	var tioForecast tioForecast

	// here we'll have to do some manual unmarshalling
	if err := json.Unmarshal(forecastBody, &tioForecast); err != nil {
		return nil, errors.New("tomorrow.io: failed to unmarshal response body")
	}

	var forecast Forecast
	forecast.StartTime = tioForecast.Data.Timelines[0].StartTime
	forecast.EndTime = tioForecast.Data.Timelines[0].EndTime
	for _, interval := range tioForecast.Data.Timelines[0].Intervals {
		var forecastInterval Interval
		forecastInterval.StartTime = interval.StartTime
		forecastInterval.CloudCover = interval.Values.CloudCover
		forecastInterval.Humidity = interval.Values.Humidity
		forecastInterval.Temperature = interval.Values.Temperature
		forecastInterval.Visibility = interval.Values.Visibility
		forecast.Intervals = append(forecast.Intervals, forecastInterval)
	}

	return &forecast, nil
}
