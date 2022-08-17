package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// locateCmd represents the locate command
var locateCmd = &cobra.Command{
	Use:   "locate",
	Short: "Convert a location to coordinates",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("locate called")
	},
}

func init() {
	rootCmd.AddCommand(locateCmd)
}
