package geo

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func NewCoordinatesFromString(commaSeparatedInput string) (*Coordinates, error) {
	separatedInput := strings.Split(commaSeparatedInput, ",")
	if len(separatedInput) != 2 {
		return nil, errors.New("invalid coordinates: " + commaSeparatedInput)
	}

	// 64 is the number of bits a float64 takes up.
	latitude, err := strconv.ParseFloat(separatedInput[0], 64)
	if err != nil {
		return nil, errors.New("invalid latitude: " + separatedInput[0])
	}

	longitude, err := strconv.ParseFloat(separatedInput[1], 64)
	if err != nil {
		return nil, errors.New("invalid latitude: " + separatedInput[1])
	}

	if err := validateLatitudeLongitude(latitude, longitude); err != nil {
		return nil, err
	}

	return &Coordinates{
		Latitude:  latitude,
		Longitude: longitude,
	}, nil
}

func validateLatitudeLongitude(latitude, longitude float64) error {
	if latitude < -90 || latitude > 90 {
		return fmt.Errorf("invalid latitude: %f", latitude)
	}

	if longitude < -180 || longitude > 180 {
		return fmt.Errorf("invalid longitude: %f", longitude)
	}

	return nil
}
