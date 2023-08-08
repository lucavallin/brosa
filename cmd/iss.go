/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"time"

	"github.com/lucavallin/brosa/pkg/astro"
	"github.com/lucavallin/brosa/pkg/ui"
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

		ui.PrintIss(issPosition)
	},
}

func init() {
	rootCmd.AddCommand(issCmd)
}
