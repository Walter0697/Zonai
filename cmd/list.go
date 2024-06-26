/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Walter0697/zonai/util"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List for all your current project ",
	Long:  `List for all your current project, it will show all your project and their child project.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectList := util.ReadProjectList()

		tw := table.NewWriter()
		tw.AppendHeader(table.Row{"#", "Project Name", "Child Name", "Flag", "Project Path"})
		for i, project := range projectList.List {
			if len(project.List) > 0 {
				firstProject := project.List[0]
				tw.AppendRow([]interface{}{i + 1, project.ProjectName, firstProject.ProjectName, firstProject.Flag, firstProject.ProjectPath})
			} else {
				tw.AppendRow([]interface{}{i + 1, project.ProjectName, "", "", ""})
			}

			for index, child := range project.List {
				if index == 0 {
					continue
				}
				tw.AppendRow([]interface{}{"", "", child.ProjectName, child.Flag, child.ProjectPath})
			}
		}
		fmt.Println(tw.Render())
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
