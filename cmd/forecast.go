package cmd

import (
	"fmt"
	"time"

	"github.com/lucavallin/mawu/pkg/tomorrowio"
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

		for _, timeline := range forecast.Data.Timelines {
			for _, interval := range timeline.Intervals {
				date, _ := time.Parse(time.RFC3339, interval.StartTime)
				fmt.Printf(
					"%s: CloudBase=%.2f, CloudCeiling=%.2f, CloudCover=%.2f, Humidity=%.2f, Temperature=%.2f, Visibility=%.2f\n",
					date.Format("2006-01-02 15:04"),
					interval.Values.CloudBase,
					interval.Values.CloudCeiling,
					interval.Values.CloudCover,
					interval.Values.Humidity,
					interval.Values.Temperature,
					interval.Values.Visibility,
				)
			}
		}
	},
}

// Set flags and configuration settings.
func init() {
	rootCmd.AddCommand(forecastCmd)
}
