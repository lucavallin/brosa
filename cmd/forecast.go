/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

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

		fmt.Println(forecast, err)
	},
}

func init() {
	rootCmd.AddCommand(forecastCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// forecastCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// forecastCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
