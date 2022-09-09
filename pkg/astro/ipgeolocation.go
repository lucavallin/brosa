package astro

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/lucavallin/mau/pkg/geo"
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

// GetDayInformation returns a pointer to a struct containing current information about the sun and the moon
func (i *IPGeolocation) GetDayInformation(coordinates *geo.Coordinates) (*DayInformation, error) {
	req, err := http.NewRequest("GET", ipgBaseUrl+"/astronomy", nil)
	if err != nil {
		return nil, errors.New("ipgeolocation.com: failed to create request")
	}

	// this could be represented as a GetDayInformation struct, but I'm not sure it's worth it
	query := req.URL.Query()
	query.Add("lat", fmt.Sprintf("%3.f", coordinates.Latitude))
	query.Add("long", fmt.Sprintf("%3.f", coordinates.Longitude))
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

	var dayInformation DayInformation
	if err := json.Unmarshal(body, &dayInformation); err != nil {
		return nil, err
	}

	return &dayInformation, nil
}
