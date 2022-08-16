package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var tomorrowioApiKey string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mau",
	Short: "The Mighty Astronomical Weather Utility.",
	Long:  `Mighty Astronomical Weather Utility. A CLI tool written in Golang that uses OpenWeatherMap to check when the weather is good for stargazing.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Set flags and configuration settings.
func init() {
	rootCmd.PersistentFlags().StringVarP(&tomorrowioApiKey, "tomorrowio-key", "k", "", "Tomorrow.io API Key (required)")
	rootCmd.MarkPersistentFlagRequired("tomorrowio-key")
}
