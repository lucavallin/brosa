package geo

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Notice: for internal names, we use "nom" instead of "Nominatim"
// nomBaseUrl is the base URL for the nominatim.org API
const nomBaseUrl = "https://nominatim.openstreetmap.org"

// GetForecast returns the forecast for the given coordinates and until the specified endTime
func GetCoordinates(location string) (*Coordinates, error) {
	req, err := http.NewRequest("GET", nomBaseUrl+"/search", nil)
	if err != nil {
		return nil, errors.New("nominatim.org: failed to create request")
	}

	// this could be represented as a GetCoordinates struct, but I'm not sure it's worth it
	query := req.URL.Query()
	query.Add("q", location)
	req.URL.RawQuery = query.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		return nil, errors.New("nominatim.org: failed to get response")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("nominatim.org: failed to read response body")
	}

	var coordinates Coordinates
	if err := json.Unmarshal(body, &coordinates); err != nil {
		return nil, errors.New("nominatim.org: failed to parse response")
	}

	return &coordinates, nil
}
