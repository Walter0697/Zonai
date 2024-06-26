/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Walter0697/zonai/model"
	"github.com/Walter0697/zonai/util"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [project name] [child name] [project path]",
	Short: "Add a project into project list",
	Long: `Add a project into project list, it does not matter if the project name exists, you can append a child project inside a single project.
	For example, [project name] can be a parent project, like POSSystem, and [child name] can be their backend and frontend.`,
	Args: cobra.ExactArgs(3),
	Example: `
	zonai add POSSystem Backend /path/to/POSSystem/Backend
	zonai add POSSystem Frontend /path/to/POSSystem/Frontend`,
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		childName := args[1]
		projectPath := args[2]

		if projectName == "" || childName == "" || projectPath == "" {
			fmt.Println("--> Please provide all arguments: [project name] [child name] [project path]")
			os.Exit(1)
		}

		projectList := util.ReadProjectList()

		parentIndex := -1
		var parentProject *model.ProjectParentModel
		for i, project := range projectList.List {
			if project.ProjectName == projectName {
				parentIndex = i
				parentProject = &project
				break
			}
		}

		if parentIndex == -1 {
			parentProject = &model.ProjectParentModel{
				ProjectName: projectName,
				List:        []model.ProjectChildModel{},
			}
		}

		childIndex := -1
		var childProject *model.ProjectChildModel
		for i, child := range parentProject.List {
			if child.ProjectName == childName {
				childIndex = i
				childProject = &child
				break
			}
		}

		flag := util.AnalysisFlag(childName, parentProject.List)
		if childIndex == -1 {
			childProject = &model.ProjectChildModel{
				ProjectName: childName,
				ProjectPath: projectPath,
				Flag:        flag,
			}
		}

		addedChild := false
		addedParent := false

		if childIndex == -1 {
			parentProject.List = append(parentProject.List, *childProject)
			addedChild = true
		} else {
			parentProject.List[childIndex] = *childProject
		}

		if parentIndex == -1 {
			projectList.List = append(projectList.List, *parentProject)
			addedParent = true
		} else {
			projectList.List[parentIndex] = *parentProject
		}

		if addedParent {
			color.Cyan("--> Added New Project " + projectName)
		} else {
			color.Cyan("--> Updated Project " + projectName)
		}
		if addedChild {
			color.Cyan("--> Added New Child Project " + childName)
			fmt.Printf("%v %v %v\n", color.CyanString("--> You can use flag"), color.YellowString(flag), color.CyanString("for Child Project "+childName))
		} else {
			color.Cyan("--> Updated Child Project " + childName)
		}
		util.SaveProjectList(projectList)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
