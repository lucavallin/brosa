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
			SetQueryParam("units", "metric").
			R(),
	}
}

func (o *Tomorrowio) GetHourlyForecast(location string) (*Forecast, error) {
	res, err := o.client.
		SetQueryParam("location", location).
		SetQueryParam("fields", "temperature,humidity,visibility,cloudCover,cloudBase,cloudCeiling").
		SetQueryParam("timesteps", "1h").
		SetQueryParam("startTime", "now").
		SetQueryParam("endTime", "nowPlus24h").
		SetResult(&Forecast{}).
		Get("timelines")

	if err != nil {
		return nil, err
	}

	return res.Result().(*Forecast), nil
}
