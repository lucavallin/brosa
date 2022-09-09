package cmd

import (
	"fmt"
	"os"

	"github.com/lucavallin/mau/pkg/astro"
	"github.com/lucavallin/mau/pkg/geo"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// dayCmd represents the day command
var dayCmd = &cobra.Command{
	Use:   "day",
	Short: "Get current information about the Sun and the Moon for a given location",

	Run: func(cmd *cobra.Command, args []string) {
		coordinates, err := geo.NewCoordinatesFromString(args[0])
		if err != nil {
			pterm.Error.Println(err)
			os.Exit(1)
		}

		ipGeolocationApiKey := viper.GetString("ipgeolocation.api_key")
		if ipGeolocationApiKey == "" {
			pterm.Error.Println("ipgeolocation.com API key not set. Please run 'mau init' to set it.")
			os.Exit(1)
		}

		ipg := astro.NewIPGeolocation(ipGeolocationApiKey)
		dayInformation, err := ipg.GetDayInformation(coordinates)

		if err != nil {
			pterm.Error.Println(err)
			os.Exit(1)
		}

		var table = pterm.TableData{
			{"Sunrise", "Sunset", "Day length (h)", "Sun altitude (°)", "Sun azimuth (°)", "Moonrise", "Moonset", "Moon altitude (°)", "Moon azimuth (°)"},
		}

		// we'll end up using this logic elsewhere too, so it's a good candidate for a function.
		table = append(table, []string{
			dayInformation.Sunrise,
			dayInformation.Sunset,
			dayInformation.DayLength,
			fmt.Sprintf("%2.f", dayInformation.SunAltitude),
			fmt.Sprintf("%2.f", dayInformation.SunAzimuth),
			dayInformation.Moonrise,
			dayInformation.Moonset,
			fmt.Sprintf("%2.f", dayInformation.MoonAltitude),
			fmt.Sprintf("%2.f", dayInformation.MoonAzimuth),
		})

		pterm.DefaultTable.WithBoxed().WithHasHeader().WithData(table).WithRightAlignment().Render()

	},
}

func init() {
	rootCmd.AddCommand(dayCmd)
}
