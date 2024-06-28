/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/Walter0697/zonai/util"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy [.gz file path]",
	Short: "Deploy a project using a compressed .gz file",
	Long: `Deploy a project using a compressed .gz file. When you built a project by zonai build -c, there will be a .gz file in the output
	Put the output directly into the server and deploy using this command`,
	Example: `
	zonai deploy /path/to/file.gz`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !util.IsDockerRunning() {
			color.Red("--> Docker is not running")
			os.Exit(1)
		}

		filename := args[0]
		pwd, err := os.Getwd()
		currentFilename := path.Join(pwd, filename)
		if err != nil {
			panic(err)
		}

		if path.Ext(currentFilename) != ".gz" {
			color.Red("File must be a .gz file")
			return
		}

		util.DrawTitle()

		// check if the file exists
		if _, err := os.Stat(currentFilename); os.IsNotExist(err) {
			color.Red("File does not exist")
			return
		}

		pwd, _ = os.Getwd()
		loadedImageList := util.LoadAllImagesFromGz(currentFilename, pwd)
		pathList := []string{}
		for _, imageTag := range loadedImageList {
			destination := util.FindComposeAndEdit(imageTag)
			if destination != "" {
				pathList = append(pathList, destination)
			}
		}

		deployInstruction := color.YellowString("docker-compose up -d")
		fmt.Println("--> To deploy, run " + deployInstruction + " in the following paths")
		for _, path := range pathList {
			color.Green("cd " + path)
		}
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
