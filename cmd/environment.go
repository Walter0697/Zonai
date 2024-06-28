/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Walter0697/zonai/util"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// environmentCmd represents the environment command
var environmentCmd = &cobra.Command{
	Use:   "environment",
	Short: "environment is a command to change the current environment of Zonai.",
	Long: `Environment is a command to change the current environment of Zonai.
	If you want see the current configuration, use zonai config.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		settingEnv := args[0]

		configuration := util.ReadConfiguration()
		configuration.CurrentEnvironment = settingEnv

		util.SaveConfiguration(configuration)

		util.DrawTitle()

		color.Cyan("Changed Environment to %s", settingEnv)
		color.Cyan("Thank you for using Zonai!")
	},
}

func init() {
	rootCmd.AddCommand(environmentCmd)
}
