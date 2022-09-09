package astro

// import (
// 	"encoding/json"
// 	"errors"
// 	"io"
// 	"net/http"
// )

// // Notice: for internal names, we use "ast" instead of "AstronomyAPI"
// // astBaseUrl is the base URL for the ipgeolocation.io API
// const astBaseUrl = "https://api.ipgeolocation.io"

// // AstronomyAPI is a client for the ipgeolocation.io API
// type AstronomyAPI struct {
// 	client *http.Client
// }

// // asaTransport is a custom transport for the AstronomyAPI client
// type astTransport struct {
// 	apiKey string
// }

// // NewAstronomyAPI returns a new AstronomyAPI client with the given API key.
// func NewAstronomyAPI(apiKey string) *IPGeolocation {
// 	return &IPGeolocation{
// 		client: &http.Client{
// 			Transport: &astTransport{
// 				apiKey: apiKey,
// 			},
// 		},
// 	}
// }

// func (a *astTransport) RoundTrip(req *http.Request) (*http.Response, error) {
// 	req.Header.Add("Accept", "application/json")

// 	query := req.URL.Query()
// 	query.Add("apiKey", a.apiKey)
// 	req.URL.RawQuery = query.Encode()

// 	return http.DefaultTransport.RoundTrip(req)
// }

// // GetBodies returns a slice of bodies visible at the given location
// func (a *IPGeolocation) GetBodies(location string) (*[]Body, error) {
// 	req, err := http.NewRequest("GET", astBaseUrl+"/search", nil)
// 	if err != nil {
// 		return nil, errors.New("astronomyapi.com: failed to create request")
// 	}

// 	// this could be represented as a GetCoordinates struct, but I'm not sure it's worth it
// 	query := req.URL.Query()
// 	query.Add("lat", location)
// 	query.Add("long", location)
// 	req.URL.RawQuery = query.Encode()

// 	res, err := a.client.Do(req)
// 	if err != nil || res.StatusCode != http.StatusOK {
// 		return nil, errors.New("astronomyapi.com: failed to get response")
// 	}
// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		return nil, errors.New("astronomyapi.com: failed to read response body")
// 	}

// 	var locationInfo ipgResult
// 	if err := json.Unmarshal(body, &locationInfo); err != nil {
// 		return nil, err
// 	}

// 	return locationInfo, nil
// }
