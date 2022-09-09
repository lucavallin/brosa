package astro

// Body represents a celestial body
type Body struct {
}

// BodyPosition is a struct that represents the position of a celestial body
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

// DayInformation is a struct that represents current sun and moon information
type DayInformation struct {
	Sunrise      string  `json:"sunrise"`
	Sunset       string  `json:"sunset"`
	DayLength    string  `json:"day_length"`
	SunAltitude  float64 `json:"sun_altitude"`
	SunAzimuth   float64 `json:"sun_azimuth"`
	Moonrise     string  `json:"moonrise"`
	Moonset      string  `json:"moonset"`
	MoonAltitude float64 `json:"moon_altitude"`
	MoonAzimuth  float64 `json:"moon_azimuth"`
}
