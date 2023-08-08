package cmd

import (
	"os"

	"github.com/lucavallin/brosa/pkg/astro"
	"github.com/lucavallin/brosa/pkg/geo"
	"github.com/lucavallin/brosa/pkg/ui"
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
			pterm.Error.Println("ipgeolocation.com API key not set. Please run 'brosa init' to set it.")
			os.Exit(1)
		}

		ipg := astro.NewIPGeolocation(ipGeolocationApiKey)
		dayInformation, err := ipg.GetDayInformation(coordinates)

		if err != nil {
			pterm.Error.Println(err)
			os.Exit(1)
		}

		ui.PrintDayInformation(dayInformation)

	},
}

func init() {
	rootCmd.AddCommand(dayCmd)
}
