package cmd

import (
	"fmt"
	"time"

	"github.com/lucavallin/mawu/pkg/tomorrowio"
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
		// todo: add coordinates validation
		coordinates := args[0]

		tio := tomorrowio.NewClient(tomorrowioApiKey)
		forecast, err := tio.GetForecast(coordinates, endTime)

		if err != nil {
			panic(err)
		}

		var table = pterm.TableData{
			{"Date", "Cloud Cover (%)", "Humidity (%)", "Temperature (ºC)", "Visibility (km)"},
		}
		for _, timeline := range forecast.Data.Timelines {
			for _, interval := range timeline.Intervals {
				date, _ := time.Parse(time.RFC3339, interval.StartTime)
				table = append(table, []string{
					date.Format("2006-01-02 15:04"),
					fmt.Sprintf("%2.f", interval.Values.CloudCover),
					fmt.Sprintf("%2.f", interval.Values.Humidity),
					fmt.Sprintf("%2.f", interval.Values.Temperature),
					fmt.Sprintf("%2.f", interval.Values.Visibility),
				})
			}
		}

		fmt.Println()
		pterm.DefaultTable.WithHasHeader().WithData(table).WithRightAlignment().Render()
	},
}

// Set flags and configuration settings.
func init() {
	forecastCmd.PersistentFlags().StringVarP(&endTime, "end-time", "e", "nowPlus24h", "End time for the forecast")

	rootCmd.AddCommand(forecastCmd)
}
