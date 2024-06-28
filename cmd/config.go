/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Walter0697/zonai/util"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Config is a command to see the current configuration of Zonai.",
	Long:  `Config is a command to see the current configuration of Zonai. You can see the current configuration of Zonai to see if you are using the correct configuration or not.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		util.DrawTitle()

		util.ShowConfiguration()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
