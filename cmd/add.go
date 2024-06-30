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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [project | deployment] [project name] [child name] [project path]",
	Short: "Add a project into project list",
	Long: `Add a project into project list, it does not matter if the project name exists, you can append a child project inside a single project.
	For project, [project name] can be a parent project, like POSSystem, and [child name] can be their backend and frontend.
	For deployment, [project name] is still the project name, but [child name] is the deployment name, [project path] is the path where docker-compose located`,
	Args: cobra.ExactArgs(4),
	Example: `
	zonai add project POSSystem Backend /path/to/POSSystem/Backend
	zonai add project POSSystem Frontend /path/to/POSSystem/Frontend
	zonai add deployment POSSystem Backend /path/to/POSSystem/Backend
	`,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) == 0 {
			return []string{"project", "deployment"}, cobra.ShellCompDirectiveNoFileComp
		}

		var list model.ProjectList
		if len(args) == 1 {
			if args[0] == "project" {
				list = util.ReadProjectList()
			} else {
				list = util.ReadDeploymentList()
			}

			var outputs []string
			for _, project := range list.List {
				if toComplete == "" || strings.Contains(project.ProjectName, toComplete) {
					outputs = append(outputs, project.ProjectName)
				}
			}
			return outputs, cobra.ShellCompDirectiveNoFileComp
		}

		return nil, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] != "project" && args[0] != "deployment" {
			color.Red("--> Please provide a valid type: [project | deployment]")
			os.Exit(1)
		}

		projectName := args[1]
		childName := args[2]
		projectPath := args[3]

		if projectName == "" || childName == "" || projectPath == "" {
			color.Red("--> Please provide all arguments: [project name] [child name] [project path]")
			os.Exit(1)
		}

		util.DrawTitle()

		if args[0] == "project" {
			projectList := util.ReadProjectList()
			updatedList := util.AddProject(projectList, projectName, childName, projectPath)
			util.SaveProjectList(updatedList)
		} else {
			deploymentList := util.ReadDeploymentList()
			updatedList := util.AddProject(deploymentList, projectName, childName, projectPath)
			util.SaveDeploymentList(updatedList)
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
