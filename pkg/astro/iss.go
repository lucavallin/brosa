package astro

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

// GetISSBodyPosition returns the position of the ISS at a given UNIX timestamp
func GetISSPosition(timestamp int64) (*BodyPosition, error) {
	req, err := http.NewRequest("GET", "https://api.wheretheiss.at/v1/satellites/25544", nil)
	if err != nil {
		return nil, errors.New("wheretheiss.at: failed to create request")
	}

	// this could be represented as a GetCoordinates struct, but I'm not sure it's worth it
	query := req.URL.Query()
	query.Add("timestamp", strconv.FormatInt(timestamp, 10))
	req.URL.RawQuery = query.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		return nil, errors.New("wheretheiss.at: failed to get response")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("wheretheiss.at: failed to read response body")
	}

	var bodyPosition *BodyPosition
	if err := json.Unmarshal(body, &bodyPosition); err != nil {
		return nil, err
	}

	return bodyPosition, nil
}
