package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/lucavallin/mau/pkg/geo"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// locateCmd represents the locate command
var locateCmd = &cobra.Command{
	Use:   "locate <location>",
	Short: "Convert a location to coordinates",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		nominatim := geo.NewNominatim()
		inputLocation := strings.Join(args, "")
		coordinates, err := nominatim.GetCoordinates(inputLocation)

		if err != nil {
			pterm.Error.Println(err)
			os.Exit(1)
		}

		pterm.Success.Printf("%d coordinate(s) found!\n", len(*coordinates))
		var table = pterm.TableData{
			{"Name", "Latitude", "Longitude", "Latitude, Longitude"},
		}
		for _, coordinate := range *coordinates {
			// we'll end up using this logic elsewhere too, so it's a good candidate for a function.
			table = append(table, []string{
				coordinate.Name,
				fmt.Sprintf("%f", coordinate.Latitude),
				fmt.Sprintf("%f", coordinate.Longitude),
				fmt.Sprintf("%f,%f", coordinate.Latitude, coordinate.Longitude),
			})
		}

		pterm.DefaultTable.WithBoxed().WithHasHeader().WithData(table).Render()
	},
}

func init() {
	rootCmd.AddCommand(locateCmd)
}
