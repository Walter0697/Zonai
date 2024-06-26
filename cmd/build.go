/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/Walter0697/zonai/model"
	"github.com/Walter0697/zonai/util"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build [project name] (-a | --all) [...extra flags for your own project]",
	Short: "Build your project into a docker image.",
	Long:  `Build your project into a docker image, with different instruction, such as Frontend only, Backend only or Fullstack.`,
	Example: `
	zonai build POSSystem -a
	zonai build POSSystem f b	// f for frontend, b for backend, please use zonai list to see your project flag
	`,
	Run: func(cmd *cobra.Command, args []string) {
		util.CheckSetup()
		projectName := args[0]
		if projectName == "" {
			color.Red("--> Please provide a project name")
			os.Exit(1)
		}

		currentProject := util.FindProject(projectName)
		if currentProject == nil {
			color.Red("--> Project not found")
			os.Exit(1)
		}

		allFlag, _ := cmd.Flags().GetBool("all")
		buildFlags := []string{}
		if allFlag {
			buildFlags = util.GetAllProjectFlags(currentProject)
		} else {
			buildFlags = args[1:]
		}

		configuration := util.ReadConfiguration()
		history := util.ReadBuildHistory()
		now := time.Now().Format("2006-01-02")
		for _, projects := range currentProject.List {
			for _, flag := range buildFlags {
				if projects.Flag == flag {
					imageName := util.GetImageName(currentProject, &projects)
					version := 1
					found := false
					for hindex, h := range history.List {
						if h.ImageName == imageName {
							if h.BuildDate == now {
								version = h.BuildVersion + 1
								h.BuildVersion = version
							} else {
								version = 1
								h.BuildVersion = version
								h.BuildDate = now
							}
							found = true
							history.List[hindex] = h
							break
						}
					}
					if !found {
						history.List = append(history.List, model.BuildItem{
							ImageName:    imageName,
							BuildDate:    now,
							BuildVersion: 1,
						})
					}
					imageTag := now
					if version != 1 {
						imageTag = imageTag + "-" + fmt.Sprintf("%d", version)
					}
					util.BuildProject(currentProject, &projects, &configuration, imageTag)
					break
				}
			}
		}
		util.SaveBuildHistory(history)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	rootCmd.PersistentFlags().BoolP("all", "a", false, "Build All Child Projects")
}
