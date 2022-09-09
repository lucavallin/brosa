package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/lucavallin/mau/pkg/geo"
	"github.com/lucavallin/mau/pkg/weather"
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

		tomorrowApiKey := viper.GetString("tomorrow.api_key")
		if tomorrowApiKey == "" {
			pterm.Error.Println("tomorrow.io API key not set. Please run 'mau init' to set it.")
			os.Exit(1)
		}

		tio := weather.NewTomorrowClient(tomorrowApiKey)
		forecast, err := tio.GetForecast(coordinates, startTime, endTime, onlyBestForecast)

		if err != nil {
			pterm.Error.Println(err)
			os.Exit(1)
		}

		var table = pterm.TableData{
			{"Date", "Cloud Cover (%)", "Humidity (%)", "Temperature (ºC)", "Visibility (km)", "Dew Point (ºC)", "Precipitation Probability (%)"},
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
				fmt.Sprintf("%2.f", interval.DewPoint),
				fmt.Sprintf("%2.f", interval.PrecipitationProbability),
			})
		}

		pterm.DefaultTable.WithBoxed().WithHasHeader().WithData(table).WithRightAlignment().Render()
	},
}

// Set flags and configuration settings.
func init() {
	forecastCmd.PersistentFlags().StringVarP(&startTime, "start-time", "s", "now", "Start time for the forecast (e.g. now, 3h, 24h, 48h, 1d, 2d, 4d, 7d)")
	forecastCmd.PersistentFlags().StringVarP(&endTime, "end-time", "e", "24h", "End time for the forecast (e.g. 3h, 24h, 48h, 1d, 2d, 4d, 7d)")
	forecastCmd.PersistentFlags().BoolVarP(&onlyBestForecast, "best", "b", true, "Retrieve only the forecast with the best weather conditions for astronomy")

	rootCmd.AddCommand(forecastCmd)
}
