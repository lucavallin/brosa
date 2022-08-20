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
		tomorrowioApiKey, err := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show("\nEnter your Tomorrow.io API key")

		if err != nil {
			pterm.Error.Println("error parsing tomorrow.io API key")
		}

		viper.Set("tomorrow_io.api_key", tomorrowioApiKey)
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
