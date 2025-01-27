/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"strings"
	"time"

	"github.com/Walter0697/zonai/util"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build [project name] (-a | --all) (-c | --compress) (-w | --whole) [...extra flags for your own project]",
	Short: "Build your project into a docker image.",
	Long: `Build your project into a docker image, with different instruction, such as Frontend only, Backend only or Fullstack.
	You can also select to compress all images after built.
	You can also build the whole project with all environments.`,
	Example: `
	zonai build POSSystem -a
	zonai build POSSystem f b	// f for frontend, b for backend, please use zonai list to see your project flag
	zonai build POSSystem -ac
	zonai build POSSystem -wc f
	`,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		projectList := util.ReadProjectList()
		if len(args) == 0 {
			var outputs []string
			for _, project := range projectList.List {
				if toComplete == "" || strings.Contains(project.ProjectName, toComplete) {
					outputs = append(outputs, project.ProjectName)
				}
			}
			return outputs, cobra.ShellCompDirectiveNoFileComp
		}

		if len(args) >= 1 {
			options := []string{"-a", "--all", "-c", "--compress", "-w", "--whole"}
			for _, project := range projectList.List {
				if project.ProjectName == args[0] {
					for _, child := range project.List {
						options = append(options, child.Flag)
					}
				}
			}

			filtered := []string{}
			for _, option := range options {
				exists := false
				for argIndex, arg := range args {
					if argIndex == 0 {
						continue
					}

					if arg == option {
						exists = true
						break
					}
				}

				if !exists {
					filtered = append(filtered, option)
				}
			}

			return filtered, cobra.ShellCompDirectiveNoFileComp
		}

		return nil, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		if !util.IsDockerRunning() {
			color.Red("--> Docker is not running")
			os.Exit(1)
		}

		if len(args) == 0 {
			color.Red("--> Please provide a project name")
			os.Exit(1)
		}

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

		util.DrawTitle()

		allFlag, _ := cmd.Flags().GetBool("all")
		compressFlag, _ := cmd.Flags().GetBool("compress")
		wholeFlag, _ := cmd.Flags().GetBool("whole")

		buildFlags := []string{}
		if allFlag {
			buildFlags = util.GetAllProjectFlags(currentProject)
		} else {
			buildFlags = args[1:]
		}

		configuration := util.ReadConfiguration()
		history := util.ReadBuildHistory()
		now := time.Now().Format("2006-01-02")

		if wholeFlag {
			env_list := util.GetAllEnvironments(&configuration, currentProject)
			for _, env := range env_list {
				util.BuildProjectWithImageList(currentProject, buildFlags, &configuration, &history, now, compressFlag, env)
			}
		} else {
			currentEnvironment := configuration.CurrentEnvironment
			util.BuildProjectWithImageList(currentProject, buildFlags, &configuration, &history, now, compressFlag, currentEnvironment)
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	buildCmd.PersistentFlags().BoolP("all", "a", false, "Build All Child Projects")
	buildCmd.PersistentFlags().BoolP("compress", "c", false, "Compress the image into a tar file")
	buildCmd.PersistentFlags().BoolP("whole", "w", false, "Whole Project with all environments")
}
