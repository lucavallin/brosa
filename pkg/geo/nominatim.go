package geo

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

// Notice: for internal names, we use "nom" instead of "Nominatim"
// nomBaseUrl is the base URL for the nominatim.org API
const nomBaseUrl = "https://nominatim.openstreetmap.org"

// Nominatim is a client for the nominatim.org API
type Nominatim struct {
	client *http.Client
}

// nomTransport is a custom transport for the Nominatim client
type nomTransport struct{}

// nomLocation is a struct that represents the location data returned by the API
type nomLocation struct {
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	DisplayName string `json:"display_name"`
}

// NewNominatim returns a new Nominatim client with the given API key.
func NewNominatim() *Nominatim {
	return &Nominatim{
		client: &http.Client{
			Transport: &nomTransport{},
		},
	}
}

func (n *nomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Accept", "application/json")

	query := req.URL.Query()
	query.Add("format", "jsonv2")
	req.URL.RawQuery = query.Encode()

	return http.DefaultTransport.RoundTrip(req)
}

// GetCoordinates returns a slice of Coordinates for the given location string
func (n *Nominatim) GetCoordinates(location string) (*[]Coordinates, error) {
	req, err := http.NewRequest("GET", nomBaseUrl+"/search", nil)
	if err != nil {
		return nil, errors.New("nominatim.org: failed to create request")
	}

	query := req.URL.Query()
	query.Add("q", location)
	req.URL.RawQuery = query.Encode()

	res, err := n.client.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		return nil, errors.New("nominatim.org: failed to get response")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("nominatim.org: failed to read response body")
	}

	coordinatesList, err := n.unmarshalLocationList(body)
	if err != nil {
		return nil, err
	}

	return coordinatesList, nil
}

// unmarshalLocationList is a wrapper around the json.Unmarshal function that
// unmarshals the location list data from the response body into a nomLocationList struct,
// but returns a slice of Coordinates instead. This is because the API returns a lot of fields we don't need
// and we want a slice of Coordinates with only the bare minimum fields
func (n *Nominatim) unmarshalLocationList(locationListBody []byte) (*[]Coordinates, error) {
	var nomLocationList []nomLocation

	if err := json.Unmarshal(locationListBody, &nomLocationList); err != nil {
		return nil, errors.New("nominatim.org: failed to unmarshal response body")
	}

	var coordinates []Coordinates
	for _, nomLocation := range nomLocationList {
		latitude, err := strconv.ParseFloat(nomLocation.Lat, 64)
		if err != nil {
			return nil, errors.New("nominatim.org: failed to parse latitude")
		}

		longitude, err := strconv.ParseFloat(nomLocation.Lon, 64)
		if err != nil {
			return nil, errors.New("nominatim.org: failed to parse longitude")
		}

		coordinates = append(coordinates, Coordinates{
			Latitude:  latitude,
			Longitude: longitude,
			Name:      nomLocation.DisplayName,
		})
	}

	return &coordinates, nil
}
