package astro

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Notice: for internal names, we use "ipg" instead of "IPGeolocation"
// ipgBaseUrl is the base URL for the ipgeolocation.io API
const ipgBaseUrl = "https://api.ipgeolocation.io"

// IPGeolocation is a client for the ipgeolocation.io API
type IPGeolocation struct {
	client *http.Client
}

// ipgTransport is a custom transport for the IPGeolocation client
type ipgTransport struct {
	apiKey string
}

// ipgResult is a struct that represents the data returned by the API
type ipgResult struct {
	Sunrise              string  `json:"sunrise"`
	Sunset               string  `json:"sunset"`
	SunStatus            string  `json:"sun_status"`
	SolarNoon            string  `json:"solar_noon"`
	DayLength            string  `json:"day_length"`
	SunAltitude          float64 `json:"sun_altitude"`
	SunDistance          float64 `json:"sun_distance"`
	SunAzimuth           float64 `json:"sun_azimuth"`
	Moonrise             string  `json:"moonrise"`
	Moonset              string  `json:"moonset"`
	MoonStatus           string  `json:"moon_status"`
	MoonAltitude         float64 `json:"moon_altitude"`
	MoonDistance         float64 `json:"moon_distance"`
	MoonAzimuth          float64 `json:"moon_azimuth"`
	MoonParallacticAngle float64 `json:"moon_parallactic_angle"`
}

// NewIPGeolocation returns a new IPGeolocation client with the given API key.
func NewIPGeolocation(apiKey string) *IPGeolocation {
	return &IPGeolocation{
		client: &http.Client{
			Transport: &ipgTransport{
				apiKey: apiKey,
			},
		},
	}
}

func (i *ipgTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Accept", "application/json")

	query := req.URL.Query()
	query.Add("apiKey", i.apiKey)
	req.URL.RawQuery = query.Encode()

	return http.DefaultTransport.RoundTrip(req)
}

// GetCoordinates returns a slice of Coordinates for the given location string
func (i *IPGeolocation) GetCoordinates(location string) (*ipgResult, error) {
	req, err := http.NewRequest("GET", ipgBaseUrl+"/search", nil)
	if err != nil {
		return nil, errors.New("ipgeolocation.com: failed to create request")
	}

	// this could be represented as a GetCoordinates struct, but I'm not sure it's worth it
	query := req.URL.Query()
	query.Add("lat", location)
	query.Add("long", location)
	req.URL.RawQuery = query.Encode()

	res, err := i.client.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		return nil, errors.New("ipgeolocation.com: failed to get response")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("ipgeolocation.com: failed to read response body")
	}

	var locationInfo ipgResult
	if err := json.Unmarshal(body, &locationInfo); err != nil {
		return nil, err
	}

	return locationInfo, nil
}
