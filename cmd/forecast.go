package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/lucavallin/mau/pkg/geo"
	"github.com/lucavallin/mau/pkg/weather"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// endTime is the end time for the forecast.
var endTime string

// forecastCmd represents the forecast command
var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Get the forecast for a set of coordinates",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		coordinates, err := geo.NewCoordinatesFromString(args[0])
		if err != nil {
			pterm.Error.Println(err)
			os.Exit(1)
		}

		tio := weather.NewTomorrowIoClient(tomorrowioApiKey)
		forecast, err := tio.GetForecast(coordinates, endTime)

		if err != nil {
			pterm.Error.Println(err)
			os.Exit(1)
		}

		var table = pterm.TableData{
			{"Date", "Cloud Cover (%)", "Humidity (%)", "Temperature (ºC)", "Visibility (km)"},
		}
		for _, interval := range forecast.Intervals {
			date, _ := time.Parse(time.RFC3339, interval.StartTime)
			// we'll end up using this logic elsewhere too, so it's a good candidate for a function.
			table = append(table, []string{
				date.Format("2006-01-02 15:04"),
				fmt.Sprintf("%2.f", interval.CloudCover),
				fmt.Sprintf("%2.f", interval.Humidity),
				fmt.Sprintf("%2.f", interval.Temperature),
				fmt.Sprintf("%2.f", interval.Visibility),
			})
		}

		pterm.DefaultTable.WithBoxed().WithHasHeader().WithData(table).WithRightAlignment().Render()
	},
}

// Set flags and configuration settings.
func init() {
	// this format for the endTime is funny, we'll have to think of a way to make it more intuitive.
	forecastCmd.PersistentFlags().StringVarP(&endTime, "end-time", "e", "nowPlus24h", "End time for the forecast")

	rootCmd.AddCommand(forecastCmd)
}
