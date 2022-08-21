/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/lucavallin/mau/pkg/astro"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// issCmd represents the iss command
var issCmd = &cobra.Command{
	Use:   "iss",
	Short: "Get the current position of the International Space Station",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		timestamp := time.Now().Unix()
		issPosition, err := astro.GetISSPosition(timestamp)

		if err != nil {
			pterm.Error.Println(err)
			os.Exit(1)
		}

		pterm.Success.Printf("ISS Found\n")

		var table = pterm.TableData{
			{"Time", "Latitude", "Longitude", "Altitude (km)", "Velocity (km/h)", "Visibility", "Solar latitude", "Solar longitude"},
		}
		time := time.Unix(issPosition.Timestamp, 0)
		table = append(table, []string{
			time.Format("2006-01-02 15:04"),
			fmt.Sprintf("%f", issPosition.Latitude),
			fmt.Sprintf("%f", issPosition.Longitude),
			fmt.Sprintf("%2.f", issPosition.Altitude),
			fmt.Sprintf("%2.f", issPosition.Velocity),
			issPosition.Visibility,
			fmt.Sprintf("%f", issPosition.SolarLatitude),
			fmt.Sprintf("%f", issPosition.SolarLongitude),
		})

		pterm.DefaultTable.WithBoxed().WithHasHeader().WithData(table).Render()
	},
}

func init() {
	rootCmd.AddCommand(issCmd)
}
