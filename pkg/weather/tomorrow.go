package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/lucavallin/mau/pkg/geo"
)

// Notice: for internal names, we use "tom" instead of "Tomorrow"
// tomBaseUrl is the base URL for the tomorrow.io API
const tomBaseUrl = "https://api.tomorrow.io/v4"

// Tomorrow is a client for the tomorrow.io API
type Tomorrow struct {
	client *http.Client
}

// tomTransport is a custom transport for the Tomorrow client
type tomTransport struct {
	apiKey string
}

// tomForecast is a struct that represents the forecast data returned by the API
type tomForecast struct {
	Data tomData `json:"data"`
}

// tomData is a struct that represents the data returned by the API
type tomData struct {
	Timelines []tomTimeline `json:"timelines"`
}

// tomTimeline is a struct that represents the timeline data returned by the API
type tomTimeline struct {
	Timestep  string        `json:"timestep"`
	EndTime   time.Time     `json:"endTime"`
	StartTime time.Time     `json:"startTime"`
	Intervals []tomInterval `json:"intervals"`
}

// tomInterval is a struct that represents the interval data returned by the API
type tomInterval struct {
	StartTime time.Time `json:"startTime"`
	Values    tomValues `json:"values"`
}

// tomValues is a struct that represents the values returned by the API
type tomValues struct {
	CloudCover               float64 `json:"cloudCover"`
	Humidity                 float64 `json:"humidity"`
	Temperature              float64 `json:"temperature"`
	Visibility               float64 `json:"visibility"`
	DewPoint                 float64 `json:"dewPoint"`
	PrecipitationProbability float64 `json:"precipitationProbability"`
}

// NewTomorrowClient returns a new Tomorrow client with the given API key.
func NewTomorrowClient(apiKey string) *Tomorrow {
	return &Tomorrow{
		client: &http.Client{
			Transport: &tomTransport{
				apiKey: apiKey,
			},
		},
	}
}

// tomTransport is a custom transport for the Tomorrow client,
// used to set common headers and provide the API key.
func (t *tomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Accept", "application/json")

	query := req.URL.Query()
	query.Add("apikey", t.apiKey)
	query.Add("units", "metric")
	req.URL.RawQuery = query.Encode()

	return http.DefaultTransport.RoundTrip(req)
}

// GetForecast returns the forecast for the given coordinates and until the specified endTime
func (t *Tomorrow) GetForecast(coordinates *geo.Coordinates, startTime time.Time, endTime time.Time, onlyBestForecast bool) (*Forecast, error) {
	req, err := http.NewRequest("GET", tomBaseUrl+"/timelines", nil)
	if err != nil {
		return nil, errors.New("tomorrow.io: failed to create request")
	}

	// this could be represented as a GetForecastRequest struct, but I'm not sure it's worth it
	query := req.URL.Query()
	query.Add("location", fmt.Sprintf("%f,%f", coordinates.Latitude, coordinates.Longitude))
	query.Add("fields", "temperature,humidity,visibility,cloudCover,dewPoint,precipitationProbability")

	// onlyBestForecast is a flag that indicates whether to retrieve only the forecast with the best weather conditions for astronomy
	// if onlyBestForecast {
	// 	query.Add("timesteps", "best")
	// } else {
	query.Add("timesteps", "1h")
	// }

	if startTime.After(endTime) {
		return nil, errors.New("tomorrow.io: startTime must be before endTime")
	}

	query.Add("startTime", startTime.Format(time.RFC3339))
	query.Add("endTime", endTime.Format(time.RFC3339))
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
// unmarshals the forecast data from the response body into a tomForecast struct,
// but returns a Forecast struct instead. This is because the API returns a lot of fields we don't need
// and we want a Forecast struct with only the bare minimum fields
func (t *Tomorrow) unmarshalForecast(forecastBody []byte) (*Forecast, error) {
	var tioForecast tomForecast

	// here we'll have to do some manual unmarshalling
	spew.Dump(forecastBody)
	if err := json.Unmarshal(forecastBody, &tioForecast); err != nil {
		return nil, errors.New("tomorrow.io: failed to unmarshal response body")
	}

	var forecast = Forecast{
		StartTime: tioForecast.Data.Timelines[0].StartTime,
		EndTime:   tioForecast.Data.Timelines[0].EndTime,
	}
	for _, interval := range tioForecast.Data.Timelines[0].Intervals {
		var forecastInterval Interval
		forecastInterval.StartTime = interval.StartTime
		forecastInterval.CloudCover = interval.Values.CloudCover
		forecastInterval.Humidity = interval.Values.Humidity
		forecastInterval.Temperature = interval.Values.Temperature
		forecastInterval.Visibility = interval.Values.Visibility
		forecastInterval.DewPoint = interval.Values.DewPoint
		forecastInterval.PrecipitationProbability = interval.Values.PrecipitationProbability
		forecast.Intervals = append(forecast.Intervals, forecastInterval)
	}

	return &forecast, nil
}
