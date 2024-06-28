/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Check the current version of Zonai.",
	Long:  `Check the current version of Zonai. You can check the version of Zonai to see if you are using the latest version or not.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Zonai version v0.1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
