package tomorrowio

type Forecast struct {
	Timelines []Timeline `json:"timelines"`
}

type Data struct {
	Timelines []Timeline `json:"timelines"`
}

type Timeline struct {
	Timestep  string     `json:"timestep"`
	EndTime   string     `json:"endTime"`
	StartTime string     `json:"startTime"`
	Intervals []Interval `json:"intervals"`
}

type Interval struct {
	StartTime string `json:"startTime"`
	Values    Values `json:"values"`
}

type Values struct {
	CloudBase    interface{} `json:"cloudBase"`
	CloudCeiling interface{} `json:"cloudCeiling"`
	CloudCover   float64     `json:"cloudCover"`
	Humidity     float64     `json:"humidity"`
	Temperature  float64     `json:"temperature"`
	Visibility   int64       `json:"visibility"`
}
