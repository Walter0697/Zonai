/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"strings"

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
	Example: `
	zonai delete project POSSystem Backend
	zonai delete deployment POSSystem Backend
	zonai delete project POSSystem -a
	`,
	Args: cobra.MaximumNArgs(3),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) == 0 {
			return []string{"project", "deployment"}, cobra.ShellCompDirectiveNoFileComp
		}

		var list model.ProjectList
		if len(args) >= 1 {
			if args[0] == "project" {
				list = util.ReadProjectList()
			} else {
				list = util.ReadDeploymentList()
			}
		}

		if len(args) == 1 {
			var outputs []string
			for _, project := range list.List {
				if toComplete == "" || strings.Contains(project.ProjectName, toComplete) {
					outputs = append(outputs, project.ProjectName)
				}
			}
			return outputs, cobra.ShellCompDirectiveNoFileComp
		}

		if len(args) == 2 {
			parentProjectName := args[1]
			var currentParent *model.ProjectParentModel
			for _, project := range list.List {
				if project.ProjectName == parentProjectName {
					currentParent = &project
					break
				}
			}

			if currentParent == nil {
				return nil, cobra.ShellCompDirectiveNoFileComp
			}

			var outputs []string
			for _, child := range currentParent.List {
				if toComplete == "" || strings.Contains(child.ProjectName, toComplete) {
					outputs = append(outputs, child.ProjectName)
				}
			}
			outputs = append(outputs, "-a")
			return outputs, cobra.ShellCompDirectiveNoFileComp
		}

		return nil, cobra.ShellCompDirectiveNoFileComp
	},
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
