/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/Walter0697/zonai/util"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [project | deployment] (-a | --all)",
	Short: "List for all your current project ",
	Long:  `List for all your current project, it will show all your project and their child project.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		allFlag, _ := cmd.Flags().GetBool("all")

		if allFlag {
			util.DrawTitle()
			util.ListProject()
			util.ListDeployment()
			return
		}

		if len(args) == 0 {
			color.Red("--> Please provide a type: [project | deployment]")
			os.Exit(1)
		}

		if args[0] != "project" && args[0] != "deployment" {
			color.Red("--> Please provide a valid type: [project | deployment]")
			os.Exit(1)
		}

		util.DrawTitle()

		if args[0] == "project" {
			util.ListProject()
		} else {
			util.ListDeployment()
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().BoolP("all", "a", false, "List all components")
}
