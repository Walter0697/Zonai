/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Walter0697/zonai/util"
	"github.com/spf13/cobra"
)

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history [project name]",
	Short: "A deployment history of all your projects",
	Long:  `A deployment history of all your projects. It will show all the deployment history of your projects.`,
	Args:  cobra.MaximumNArgs(1),
	Example: `
	zonai history
	zonai history my_project`,
	Run: func(cmd *cobra.Command, args []string) {
		var projectName *string
		if len(args) > 0 {
			projectName = &args[0]
		}

		util.DisplayHistory(projectName)
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
}
