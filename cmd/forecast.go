package cmd

import (
	"github.com/davecgh/go-spew/spew"
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

		spew.Dump(forecast, err)
	},
}

// Set flags and configuration settings.
func init() {
	rootCmd.AddCommand(forecastCmd)
}
