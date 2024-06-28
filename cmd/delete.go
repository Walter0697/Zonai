/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/Walter0697/zonai/model"
	"github.com/Walter0697/zonai/util"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [project | deployment] [parent name] [child name] (-a | --all)",
	Short: "Delete a project or deployment",
	Long: `Delete a project or deployment. You can delete a project or deployment by providing the parent name.
	You can also delete all under the parent by using the -a flag`,
	Args: cobra.MaximumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		allFlags, _ := cmd.Flags().GetBool("all")

		if len(args) <= 1 {
			color.Red("--> Please provide a type: [project | deployment]")
			os.Exit(1)
		}

		if args[0] != "project" && args[0] != "deployment" {
			color.Red("--> Please provide a valid type: [project | deployment]")
			os.Exit(1)
		}

		parent_name := args[1]
		child_name := ""

		if len(args) == 3 {
			child_name = args[2]
		}

		if !allFlags && child_name == "" {
			color.Red("--> Please provide a child name or use -a flag to delete all")
			os.Exit(1)
		}

		util.DrawTitle()

		var projectList model.ProjectList
		var updatedList model.ProjectList
		if args[0] == "project" {
			projectList = util.ReadProjectList()
		} else {
			projectList = util.ReadDeploymentList()
		}

		if allFlags {
			updatedList = util.RemoveWholeParentProject(projectList, parent_name)
		} else {
			updatedList = util.RemoveProject(projectList, parent_name, child_name)
		}

		if args[0] == "project" {
			util.SaveProjectList(updatedList)
		} else {
			util.SaveDeploymentList(updatedList)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.PersistentFlags().BoolP("all", "a", false, "Delete all under the parent")
}
