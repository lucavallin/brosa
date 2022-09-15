package cmd

import (
	"os"
	"strings"

	"github.com/lucavallin/mau/pkg/geo"
	"github.com/lucavallin/mau/pkg/ui"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// locateCmd represents the locate command
var locateCmd = &cobra.Command{
	Use:   "locate <location>",
	Short: "Convert a location to coordinates",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		inputLocation := strings.Join(args, " ")

		client := geo.NewNominatim()
		coordinates, err := geo.GetCoordinates(client, inputLocation)

		if err != nil {
			pterm.Error.Println(err)
			os.Exit(1)
		}

		pterm.Success.Printf("%d coordinate(s) found!\n", len(*coordinates))
		ui.PrintCoordinates(coordinates)
	},
}

func init() {
	rootCmd.AddCommand(locateCmd)
}
