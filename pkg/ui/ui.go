package ui

import (
	"fmt"
	"time"

	"github.com/lucavallin/brosa/pkg/astro"
	"github.com/lucavallin/brosa/pkg/geo"
	"github.com/lucavallin/brosa/pkg/weather"
	"github.com/pterm/pterm"
)

// PrintForecast prints a forecast to the terminal.
func PrintForecast(f *weather.Forecast) error {
	var table = pterm.TableData{
		{"Date", "Cloud Cover (%)", "Humidity (%)", "Temperature (ºC)", "Visibility (km)", "Dew Point (ºC)", "Precipitation Probability (%)"},
	}
	for _, interval := range f.Intervals {
		table = append(table, []string{
			interval.StartTime.Format("2006-01-02 15:04"),
			fmt.Sprintf("%2.f", interval.CloudCover),
			fmt.Sprintf("%2.f", interval.Humidity),
			fmt.Sprintf("%2.f", interval.Temperature),
			fmt.Sprintf("%2.f", interval.Visibility),
			fmt.Sprintf("%2.f", interval.DewPoint),
			fmt.Sprintf("%2.f", interval.PrecipitationProbability),
		})
	}

	return pterm.DefaultTable.WithBoxed().WithHasHeader().WithData(table).WithRightAlignment().Render()
}

// PrintIss prints the current position of the ISS to the terminal.
func PrintIss(i *astro.BodyPosition) error {
	var table = pterm.TableData{
		{"Time", "Latitude", "Longitude", "Altitude (km)", "Velocity (km/h)", "Visibility", "Solar latitude", "Solar longitude"},
	}
	time := time.Unix(i.Timestamp, 0)
	table = append(table, []string{
		time.Format("2006-01-02 15:04"),
		fmt.Sprintf("%f", i.Latitude),
		fmt.Sprintf("%f", i.Longitude),
		fmt.Sprintf("%2.f", i.Altitude),
		fmt.Sprintf("%2.f", i.Velocity),
		i.Visibility,
		fmt.Sprintf("%f", i.SolarLatitude),
		fmt.Sprintf("%f", i.SolarLongitude),
	})

	return pterm.DefaultTable.WithBoxed().WithHasHeader().WithData(table).Render()
}

// PrintDayInformation prints information about a day to the terminal.
func PrintDayInformation(d *astro.DayInformation) error {
	var table = pterm.TableData{
		{"Sunrise", "Sunset", "Day length (h)", "Sun altitude (°)", "Sun azimuth (°)", "Moonrise", "Moonset", "Moon altitude (°)", "Moon azimuth (°)"},
	}

	table = append(table, []string{
		d.Sunrise,
		d.Sunset,
		d.DayLength,
		fmt.Sprintf("%2.f", d.SunAltitude),
		fmt.Sprintf("%2.f", d.SunAzimuth),
		d.Moonrise,
		d.Moonset,
		fmt.Sprintf("%2.f", d.MoonAltitude),
		fmt.Sprintf("%2.f", d.MoonAzimuth),
	})

	return pterm.DefaultTable.WithBoxed().WithHasHeader().WithData(table).WithRightAlignment().Render()
}

// PrintCoordinates prints location information to the terminal.
func PrintCoordinates(c *[]geo.Coordinates) error {
	var table = pterm.TableData{
		{"Name", "Latitude", "Longitude", "Latitude, Longitude"},
	}
	for _, coordinate := range *c {
		table = append(table, []string{
			coordinate.Name,
			fmt.Sprintf("%f", coordinate.Latitude),
			fmt.Sprintf("%f", coordinate.Longitude),
			fmt.Sprintf("%f,%f", coordinate.Latitude, coordinate.Longitude),
		})
	}

	return pterm.DefaultTable.WithBoxed().WithHasHeader().WithData(table).Render()
}
