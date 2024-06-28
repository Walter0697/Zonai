/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Walter0697/zonai/util"
	"github.com/jedib0t/go-pretty/table"
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

		configuration := util.ReadConfiguration()
		util.Divider()
		fmt.Println("Configuration")

		tw := table.NewWriter()
		tw.AppendHeader(table.Row{"Key", "Value"})
		tw.AppendRow([]interface{}{"Docker Build Command", configuration.DockerBuildCommand})
		tw.AppendRow([]interface{}{"Output Image Path", configuration.OutputImagePath})
		tw.AppendRow([]interface{}{"Environment Path", configuration.EnviromentPath})
		tw.AppendRow([]interface{}{"Current Environment", configuration.CurrentEnvironment})

		fmt.Println(tw.Render())

		util.Divider()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
