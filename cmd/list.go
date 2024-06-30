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
	Use:   "list [project | deployment] (-a | --all) (-f | --find <key_word>)",
	Short: "List for all your current project",
	Long:  `List for all your current project, it will show all your project and their child projects.`,
	Example: `
	zonai list project
	zonai list deployment
	zonai list -a
	zonai list -f POS
	`,
	Args:      cobra.MaximumNArgs(1),
	ValidArgs: []string{"project", "deployment"},
	Run: func(cmd *cobra.Command, args []string) {
		allFlag, _ := cmd.Flags().GetBool("all")
		findFlag, _ := cmd.Flags().GetString("find")

		if allFlag {
			util.DrawTitle()
			util.ListProject(&findFlag)
			util.ListDeployment(&findFlag)
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
			util.ListProject(&findFlag)
		} else {
			util.ListDeployment(&findFlag)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().BoolP("all", "a", false, "List all components")
	listCmd.PersistentFlags().StringP("find", "f", "", "Find a project or deployment")
}
