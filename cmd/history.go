/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strings"

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
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		list := util.ReadDeploymentHistory()

		var outputMap = make(map[string]bool)
		for _, project := range list.List {
			for _, image := range project.ImageList {
				if toComplete == "" || strings.Contains(image.ProjectName, toComplete) {
					outputMap[image.ProjectName] = true
				}
			}
		}

		var outputs []string
		for key := range outputMap {
			outputs = append(outputs, key)
		}

		return outputs, cobra.ShellCompDirectiveNoFileComp
	},
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
