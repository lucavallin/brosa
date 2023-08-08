package cmd

import (
	"os"
	"time"

	"github.com/lucavallin/brosa/pkg/geo"
	"github.com/lucavallin/brosa/pkg/ui"
	"github.com/lucavallin/brosa/pkg/weather"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// startTime is the start time for the forecast.
var startTime string

// endTime is the end time for the forecast.
var endTime string

// onlyBestForecast is a flag that indicates whether to retrieve only the forecast with the best weather conditions for astronomy.
var onlyBestForecast bool

// forecastCmd represents the forecast command
var forecastCmd = &cobra.Command{
	Use:   "forecast <coordinates>",
	Short: "Get the forecast for a set of coordinates",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		coordinates, err := geo.NewCoordinatesFromString(args[0])
		if err != nil {
			pterm.Error.Println(err)
			os.Exit(1)
		}

		startTime, err := time.Parse(time.RFC3339, startTime)
		if err != nil {
			pterm.Error.Println("Invalid start-time provided. Please provide a valid ISO 8601 time.")
		}

		endTime, err := time.Parse(time.RFC3339, endTime)
		if err != nil {
			pterm.Error.Println("Invalid start-time provided. Please provide a valid ISO 8601 time.")
		}

		tomorrowApiKey := viper.GetString("tomorrow.api_key")
		if tomorrowApiKey == "" {
			pterm.Error.Println("tomorrow.io API key not set. Please run 'brosa init' to set it.")
			os.Exit(1)
		}

		client := weather.NewTomorrowClient(tomorrowApiKey)
		req := &weather.ForecastRequest{
			Location:  coordinates,
			StartTime: startTime,
			EndTime:   endTime,
		}
		forecast, err := weather.GetForecast(client, req)

		if err != nil {
			pterm.Error.Println(err)
			os.Exit(1)
		}

		ui.PrintForecast(forecast)
	},
}

// Set flags and configuration settings.
func init() {
	forecastCmd.PersistentFlags().StringVarP(&startTime, "start-time", "s", time.Now().Format(time.RFC3339), "Start time for the forecast (in ISO 8601 format)")
	forecastCmd.PersistentFlags().StringVarP(&endTime, "end-time", "e", time.Now().AddDate(0, 0, 1).Format(time.RFC3339), "End time for the forecast (in ISO 8601 format)")
	forecastCmd.PersistentFlags().BoolVarP(&onlyBestForecast, "best", "b", true, "Retrieve only the forecast with the best weather conditions for astronomy")

	rootCmd.AddCommand(forecastCmd)
}
