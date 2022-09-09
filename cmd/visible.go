package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var objectType string

// visibleCmd represents the visible command
var visibleCmd = &cobra.Command{
	Use:   "visible <coordinates>",
	Short: "Get visible objects in the sky for a set of coordinates",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pterm.Success.Println("visible called")
	},
}

// Set flags and configuration settings.
func init() {
	// this format for the endTime is funny, we'll have to think of a way to make it more intuitive.
	visibleCmd.PersistentFlags().StringVarP(&objectType, "object-type", "t", "all", "Type of objects to search for")

	rootCmd.AddCommand(visibleCmd)
}
