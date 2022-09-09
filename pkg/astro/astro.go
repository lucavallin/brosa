package astro

// BodyPosition is a struct that represents the body position data returned by the API
type BodyPosition struct {
	Name           string  `json:"name"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	Altitude       float64 `json:"altitude"`
	Velocity       float64 `json:"velocity"`
	Visibility     string  `json:"visibility"`
	Timestamp      int64   `json:"timestamp"`
	SolarLatitude  float64 `json:"solar_lat"`
	SolarLongitude float64 `json:"solar_lon"`
}
