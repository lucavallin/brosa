package tomorrowio

import (
	"github.com/go-resty/resty/v2"
)

type Tomorrowio struct {
	client *resty.Request
}

func NewClient(apiKey string) *Tomorrowio {
	return &Tomorrowio{
		client: resty.New().
			SetBaseURL("https://api.tomorrow.io/v4").
			SetHeader("Accept", "application/json").
			SetHeader("Accept-Encoding", "gzip").
			SetQueryParam("apikey", apiKey).
			R(),
	}
}

func (o *Tomorrowio) GetHourlyForecast(location string) (*resty.Response, error) {
	return o.client.
		SetQueryParam("location", location).
		SetQueryParam("fields", "temperature").
		SetQueryParam("units", "metric").
		SetQueryParam("timesteps", "1h").
		SetQueryParam("startTime", "now").
		SetQueryParam("endTime", "nowPlus6h").
		Get("timelines")
}
