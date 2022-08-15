package cmd

import (
	"fmt"
	"time"

	"github.com/lucavallin/mawu/pkg/tomorrowio"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// forecastCmd represents the forecast command
var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Get the forecast for a city",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		coordinates := args[0]

		tio := tomorrowio.NewClient(tomorrowioApiKey)
		forecast, err := tio.GetHourlyForecast(coordinates)

		if err != nil {
			panic(err)
		}

		var table = pterm.TableData{
			{"Date", "Cloud Base (%)", "Cloud Ceiling (%)", "Cloud Cover (%)", "Humidity (%)", "Temperature (ºC)", "Visibility (km)"},
		}
		for _, timeline := range forecast.Data.Timelines {
			for _, interval := range timeline.Intervals {
				date, _ := time.Parse(time.RFC3339, interval.StartTime)
				table = append(table, []string{
					date.Format("2006-01-02 15:04"),
					fmt.Sprintf("%2.f", interval.Values.CloudBase),
					fmt.Sprintf("%2.f", interval.Values.CloudCeiling),
					fmt.Sprintf("%2.f", interval.Values.CloudCover),
					fmt.Sprintf("%2.f", interval.Values.Humidity),
					fmt.Sprintf("%2.f", interval.Values.Temperature),
					fmt.Sprintf("%2.f", interval.Values.Visibility),
				})
			}
		}

		pterm.DefaultTable.WithHasHeader().WithData(table).Render()
	},
}

// Set flags and configuration settings.
func init() {
	rootCmd.AddCommand(forecastCmd)
}
