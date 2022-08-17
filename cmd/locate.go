package cmd

import (
	"fmt"
	"os"

	"github.com/lucavallin/mau/pkg/geo"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// locateCmd represents the locate command
var locateCmd = &cobra.Command{
	Use:   "locate",
	Short: "Convert a location to coordinates",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		coordinates, err := geo.GetCoordinates(args[0])

		if err != nil {
			pterm.Error.Println(err)
			os.Exit(1)
		}

		fmt.Println(coordinates)
	},
}

func init() {
	rootCmd.AddCommand(locateCmd)
}
