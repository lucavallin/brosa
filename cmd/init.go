package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the mau CLI configuration",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		pterm.Info.Println("initializing mau configuration")

		// Set the Tomorrow.io API key
		tomorrowApiKey, err := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("\nEnter your Tomorrow.io API key")
		if err != nil {
			pterm.Error.Println("error parsing tomorrow.io API key")
		}
		viper.Set("tomorrow.api_key", tomorrowApiKey)

		// Set the IPGeolocation.io API key
		ipGeolocationApiKey, err := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("\nEnter your IPGeolocation.io API key")
		if err != nil {
			pterm.Error.Println("error parsing ipgeolocation.io API key")
		}
		viper.Set("ipgeolocation.api_key", ipGeolocationApiKey)

		// Set the Astronomyapi.com Application Id
		astronomyApiApplicationId, err := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("\nEnter your Astronomyapi.com Application Id")
		if err != nil {
			pterm.Error.Println("error parsing Astronomyapi.com Application Id")
		}
		viper.Set("astronomyapi.application_id", astronomyApiApplicationId)

		// Set the Astronomyapi.com Application Secret
		astronomyApiApplicationSecret, err := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("\nEnter your Astronomyapi.com Application Secret")
		if err != nil {
			pterm.Error.Println("error parsing Astronomyapi.com Application Secret")
		}
		viper.Set("astronomyapi.application_secret", astronomyApiApplicationSecret)

		// Save the configuration to mau.yml
		viper.WriteConfigAs(viper.ConfigFileUsed())
		if err != nil {
			pterm.Error.Println("error writing mau configuration")
		}

		pterm.Println("\n")
		pterm.Success.Println("mau configuration initialized")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
