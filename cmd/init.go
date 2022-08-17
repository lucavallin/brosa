package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the mau CLI configuration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// this command should init a config file with the API key for the Tomorrow.io weather provider.
		fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
